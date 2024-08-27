package request

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi"
	"github.com/nhan1603/CloneScc/api/internal/controller/requests"
	requestCtrl "github.com/nhan1603/CloneScc/api/internal/controller/requests"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHandler_CreateNewRequest(t *testing.T) {
	tcs := map[string]struct {
		callCntroller   bool
		controllerInput requests.CreateRequestInput
		mockControler   error
		givenInput      string
		expCode         int
		expRes          string
		expErr          error
	}{
		"success": {
			givenInput: `{"alertId":"50","requestBy":"10","assignedUserId":"12","content":"null"}`,
			controllerInput: requests.CreateRequestInput{
				AlertID:        50,
				RequestBy:      10,
				AssignedUserID: 12,
				Content:        "null",
			},
			callCntroller: true,
			expCode:       http.StatusOK,
			expRes:        `{"success":true}`,
		},
		"error alert id": {
			givenInput:    `{"alertId":"-10","requestBy":"10","assignedUserId":"12","content":"null"}`,
			callCntroller: false,
			expCode:       http.StatusBadRequest,
			expRes:        `{"error":"validation_failed","error_description":"invalid alertID"}`,
		},
		"error request id": {
			givenInput:    `{"alertId":"10","requestBy":"-10","assignedUserId":"12","content":"null"}`,
			callCntroller: false,
			expCode:       http.StatusBadRequest,
			expRes:        `{"error":"validation_failed","error_description":"invalid requestBy"}`,
		},
		"error assign id": {
			givenInput:    `{"alertId":"10","requestBy":"10","assignedUserId":"-12","content":"null"}`,
			callCntroller: false,
			expCode:       http.StatusBadRequest,
			expRes:        `{"error":"validation_failed","error_description":"invalid assignedUserId"}`,
		},
		"error same id": {
			givenInput:    `{"alertId":"10","requestBy":"10","assignedUserId":"10","content":"null"}`,
			callCntroller: false,
			expCode:       http.StatusBadRequest,
			expRes:        `{"error":"validation_failed","error_description":"requester and assigned user should be different"}`,
		},
		"err from controller": {
			givenInput: `{"alertId":"50","requestBy":"10","assignedUserId":"12","content":"null"}`,
			controllerInput: requests.CreateRequestInput{
				AlertID:        50,
				RequestBy:      10,
				AssignedUserID: 12,
				Content:        "null",
			},
			callCntroller: true,
			mockControler: errors.New("error from controlller"),
			expCode:       http.StatusInternalServerError,
			expRes:        `{"error":"internal_error","error_description":"Something went wrong"}`,
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			// Setup
			req := httptest.NewRequest(http.MethodPost, "/api/authenticated/v1/requests", strings.NewReader(tc.givenInput))
			req.Header.Set("Content-Type", "application/json")

			ctx := context.WithValue(req.Context(), chi.RouteCtxKey, chi.NewRouteContext())
			req = req.WithContext(ctx)

			res := httptest.NewRecorder()

			// Given
			mockCtrl := new(requestCtrl.MockController)
			if tc.callCntroller {
				mockCtrl.ExpectedCalls = []*mock.Call{
					mockCtrl.On("CreateRequest", ctx, tc.controllerInput).Return(tc.mockControler),
				}
			}

			// When
			h := Handler{requestCtrl: mockCtrl}
			handler := http.HandlerFunc(h.CreateNewRequest())
			handler.ServeHTTP(res, req)

			// Then
			mockCtrl.AssertExpectations(t)
			require.Equal(t, tc.expCode, res.Code)
			require.Equal(t, tc.expRes, res.Body.String())
		})
	}
}
