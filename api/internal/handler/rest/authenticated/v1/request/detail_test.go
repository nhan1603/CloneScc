package request

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/nhan1603/CloneScc/api/internal/controller/requests"
	"github.com/nhan1603/CloneScc/api/internal/model"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHandler_GetRequestDetail(t *testing.T) {
	tcs := map[string]struct {
		requestID        string
		mockRequestsCtrl struct {
			input  int64
			output model.RequestDetail
			err    error
		}
		expStatusCode int
		expRespond    string
	}{
		"success": {
			requestID: "40",
			mockRequestsCtrl: struct {
				input  int64
				output model.RequestDetail
				err    error
			}{
				input: 40,
				output: model.RequestDetail{
					ID:        40,
					Title:     "Camera 1-30",
					Author:    "John Doe",
					Assignee:  "John Henry",
					Message:   "Property Damage Detected! Please check the monitored area and take screenshots for verification. Report any damage or suspicious activities to the authorities.",
					StartTime: time.Date(2023, 07, 25, 0, 0, 0, 0, time.UTC),
					AlertDetail: model.Alert{
						ID:              30,
						Type:            "Suspicious Activities",
						Description:     "TEST",
						PremiseName:     "Sunrise Tower",
						PremiseLocation: "307/12 Nguyen Van Troi St, W1, Tan Binh",
						CCTVDevice:      "Camera 1",
						CCTVDeviceFloor: 1,
						IsAcknowledged:  true,
						IncidentAt:      time.Date(2023, 07, 25, 0, 0, 0, 0, time.UTC),
					},
				},
			},
			expStatusCode: http.StatusOK,
			expRespond: `{
				"id": "40",
				"title": "Camera 1-30",
				"author": "John Doe",
				"assignee": "John Henry",
				"message": "Property Damage Detected! Please check the monitored area and take screenshots for verification. Report any damage or suspicious activities to the authorities.",
				"startTime": "2023-07-25T00:00:00Z",
				"alertDetail": {
					"id": "30",
					"type": "Suspicious Activities",
					"premiseName": "Sunrise Tower",
					"premiseLocation": "307/12 Nguyen Van Troi St, W1, Tan Binh",
					"cctvDevice": "Camera 1",
					"cctvDeviceFloor": 1,
					"isAcknowledged": true,
					"incidentAt": "2023-07-25T00:00:00Z"
				},
				"respond": null
			}`,
		},
		"error|failure to get request id param": {
			requestID:     "40abc",
			expStatusCode: http.StatusBadRequest,
			expRespond: `{
				"error": "invalid_request_body",
				"error_description": "unable to get request id parameter"
			}`,
		},
		"error|failure request not found": {
			requestID: "40",
			mockRequestsCtrl: struct {
				input  int64
				output model.RequestDetail
				err    error
			}{
				input: 40,
				err:   requests.ErrNotFound,
			},
			expStatusCode: http.StatusBadRequest,
			expRespond: `{
				"error": "request_not_found",
				"error_description": "the request not found"
			}`,
		},
		"error|failure get request got some error": {
			requestID: "40",
			mockRequestsCtrl: struct {
				input  int64
				output model.RequestDetail
				err    error
			}{
				input: 40,
				err:   errors.New("something's error"),
			},
			expStatusCode: http.StatusInternalServerError,
			expRespond: `{
				"error": "internal_error",
				"error_description": "Something went wrong"
			}`,
		},
	}
	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			req.Header.Set("Content-Type", "application/json")

			chiCtx := chi.NewRouteContext()
			chiCtx.URLParams.Add("requestID", tc.requestID)
			ctx := context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx)
			req = req.WithContext(ctx)
			res := httptest.NewRecorder()

			requestsCtrl := new(requests.MockController)
			requestsCtrl.ExpectedCalls = []*mock.Call{
				requestsCtrl.On("Detail", ctx, tc.mockRequestsCtrl.input).Return(tc.mockRequestsCtrl.output, tc.mockRequestsCtrl.err),
			}

			// When
			instance := New(requestsCtrl)
			handler := instance.GetRequestDetail()
			handler.ServeHTTP(res, req)

			// Then
			require.Equal(t, tc.expStatusCode, res.Code)
			require.JSONEq(t, tc.expRespond, res.Body.String())
		})
	}
}
