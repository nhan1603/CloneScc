package request

import (
	"context"
	"errors"
	"image"
	"image/color"
	"image/png"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/nhan1603/CloneScc/api/internal/controller/requests"
	requestCtrl "github.com/nhan1603/CloneScc/api/internal/controller/requests"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHandler_CreateResponse_testData(t *testing.T) {
	tcs := map[string]struct {
		givenInput      string
		callCntroller   bool
		controllerInput requests.CreateResponseInput
		wsInput         int64
		mockControler   error
		expCode         int
		expRes          string
		expErr          error
	}{
		"success": {
			givenInput: `{"requestID":"10000","message": "Test"}`,
			controllerInput: requests.CreateResponseInput{
				RequestID: 10000,
				Message:   "Test",
			},
			wsInput:       10000,
			callCntroller: true,
			expCode:       http.StatusOK,
			expRes:        `{"success":true}`,
		},
		"error requestID negative": {
			givenInput:    `{"requestID":"-1","message": "Test"}`,
			callCntroller: false,
			expRes:        `{"error":"validation_failed","error_description":"invalid requestID"}`,
			expCode:       http.StatusBadRequest,
			expErr:        webErrInvalidRequestID,
		},
		"error empty message": {
			givenInput:    `{"requestID":"100","message": ""}`,
			callCntroller: false,
			expRes:        `{"error":"validation_failed","error_description":"message is required"}`,
			expCode:       http.StatusBadRequest,
			expErr:        webErrInvalidRequestID,
		},
		"err not found": {
			givenInput: `{"requestID":"10000","message": "Test"}`,
			controllerInput: requests.CreateResponseInput{
				RequestID: 10000,
				Message:   "Test",
			},
			callCntroller: true,
			mockControler: requests.ErrNotFound,
			expCode:       http.StatusBadRequest,
			expRes:        `{"error":"request_not_found","error_description":"the request not found"}`,
		},
		"already resolved": {
			givenInput: `{"requestID":"10000","message": "Test"}`,
			controllerInput: requests.CreateResponseInput{
				RequestID: 10000,
				Message:   "Test",
			},
			callCntroller: true,
			mockControler: requests.ErrRequestResolved,
			expCode:       http.StatusBadRequest,
			expRes:        `{"error":"validation_failed","error_description":"this request is already resolved"}`,
		},
		"internal": {
			givenInput: `{"requestID":"10000","message": "Test"}`,
			controllerInput: requests.CreateResponseInput{
				RequestID: 10000,
				Message:   "Test",
			},
			callCntroller: true,
			mockControler: errors.New("test"),
			expCode:       http.StatusInternalServerError,
			expRes:        `{"error":"internal_error","error_description":"Something went wrong"}`,
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			// Setup
			pr, pw := io.Pipe()

			writer := multipart.NewWriter(pw)

			go func() {
				defer writer.Close()

				part, err := writer.CreateFormField("data")
				if err != nil {
					t.Error(err)
				}

				part.Write([]byte(tc.givenInput))
			}()

			req := httptest.NewRequest(http.MethodPost, "/api/authenticated/v1/requests/response", pr)
			req.Header.Set("Content-Type", writer.FormDataContentType())

			ctx := context.WithValue(req.Context(), chi.RouteCtxKey, chi.NewRouteContext())
			req = req.WithContext(ctx)

			res := httptest.NewRecorder()

			// Given

			mockCtrl := new(requestCtrl.MockController)
			if tc.callCntroller {
				mockCtrl.ExpectedCalls = []*mock.Call{
					mockCtrl.On("CreateResponse", ctx, tc.controllerInput).Return(tc.mockControler),
				}
				if tc.mockControler == nil {
					mockCtrl.ExpectedCalls = append(mockCtrl.ExpectedCalls, []*mock.Call{mockCtrl.On("Push", ctx, tc.wsInput).Return(nil)}...)
				}
			}

			// When
			h := Handler{requestCtrl: mockCtrl}
			handler := http.HandlerFunc(h.CreateResponse())
			handler.ServeHTTP(res, req)

			// Then
			mockCtrl.AssertExpectations(t)
			require.Equal(t, tc.expCode, res.Code)
			require.Equal(t, tc.expRes, res.Body.String())
		})
	}
}

func TestHandler_CreateResponse_testImage(t *testing.T) {
	tcs := map[string]struct {
		givenInput    string
		callCntroller bool
		mockControler error
		wsInput       int64
		expCode       int
		expRes        string
		expErr        error
		imageNumber   int
	}{
		"success": {
			givenInput:    `{"requestID":"10000","message": "Test"}`,
			wsInput:       10000,
			callCntroller: true,
			expCode:       http.StatusOK,
			expRes:        `{"success":true}`,
			imageNumber:   1,
		},
		"fail too many image": {
			givenInput:    `{"requestID":"10000","message": "Test"}`,
			wsInput:       10000,
			callCntroller: false,
			expCode:       http.StatusBadRequest,
			expRes:        `{"error":"invalid_request_body","error_description":"max media allowed is 20"}`,
			imageNumber:   22,
		},
		"fail too large": {
			givenInput:    `{"requestID":"10000","message": "Test"}`,
			wsInput:       10000,
			callCntroller: false,
			expCode:       http.StatusBadRequest,
			expRes:        `{"error":"invalid_request_body","error_description":"max memory allowed is 100MB"}`,
			imageNumber:   10000,
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			// Setup
			pr, pw := io.Pipe()

			writer := multipart.NewWriter(pw)

			go func() {
				defer writer.Close()

				for i := 0; i < tc.imageNumber; i++ {
					part, err := writer.CreateFormFile("file", strconv.Itoa(i)+"someimg.png")
					if err != nil {
						t.Error(err)
					}

					img := createImage()

					err = png.Encode(part, img)
					if err != nil {
						t.Error(err)
					}
				}

				data, err := writer.CreateFormField("data")
				if err != nil {
					t.Error(err)
				}

				data.Write([]byte(tc.givenInput))
			}()

			req := httptest.NewRequest(http.MethodPost, "/api/authenticated/v1/requests/response", pr)
			req.Header.Set("Content-Type", writer.FormDataContentType())

			ctx := context.WithValue(req.Context(), chi.RouteCtxKey, chi.NewRouteContext())
			req = req.WithContext(ctx)

			res := httptest.NewRecorder()

			// Given
			mockCtrl := new(requestCtrl.MockController)
			if tc.callCntroller {
				mockCtrl.ExpectedCalls = []*mock.Call{
					mockCtrl.On("CreateResponse", ctx, mock.Anything).Return(tc.mockControler),
				}
				if tc.mockControler == nil {
					mockCtrl.ExpectedCalls = append(mockCtrl.ExpectedCalls, []*mock.Call{mockCtrl.On("Push", ctx, tc.wsInput).Return(nil)}...)
				}
			}

			// When
			h := Handler{requestCtrl: mockCtrl}
			handler := http.HandlerFunc(h.CreateResponse())
			handler.ServeHTTP(res, req)

			// Then
			mockCtrl.AssertExpectations(t)
			require.Equal(t, tc.expCode, res.Code)
			require.Equal(t, tc.expRes, res.Body.String())
		})
	}
}

func createImage() *image.RGBA {
	width := 200
	height := 100

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Colors are defined by Red, Green, Blue, Alpha uint8 values.
	cyan := color.RGBA{100, 200, 200, 0xff}

	// Set color for each pixel.
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			switch {
			case x < width/2 && y < height/2: // upper left quadrant
				img.Set(x, y, cyan)
			case x >= width/2 && y >= height/2: // lower right quadrant
				img.Set(x, y, color.White)
			default:
				// Use zero value.
			}
		}
	}

	// Encode as PNG.
	f, _ := os.Create("image.png")
	png.Encode(f, img)
	return img
}
