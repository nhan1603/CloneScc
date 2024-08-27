package alert

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/nhan1603/CloneScc/api/internal/controller/alerts"
	"github.com/nhan1603/CloneScc/api/internal/model"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHandler_GetAlertDetail(t *testing.T) {
	tcs := map[string]struct {
		alertID        string
		mockAlertsCtrl struct {
			input  int64
			output model.Alert
			err    error
		}
		expStatusCode int
		expRespond    string
	}{
		"success": {
			alertID: "30",
			mockAlertsCtrl: struct {
				input  int64
				output model.Alert
				err    error
			}{
				input: 30,
				output: model.Alert{
					ID:              30,
					CCTVDeviceID:    20,
					Type:            "Suspicious Activities",
					Description:     "TEST",
					PremiseName:     "Sunrise Tower",
					PremiseLocation: "307/12 Nguyen Van Troi St, W1, Tan Binh",
					CCTVDevice:      "CCTV 1",
					CCTVDeviceFloor: 1,
					IsAcknowledged:  true,
					IncidentAt:      time.Date(2023, 7, 25, 0, 0, 0, 0, time.UTC),
				},
			},
			expStatusCode: http.StatusOK,
			expRespond: `{
				"id": "30",
				"type": "Suspicious Activities",
				"premiseName": "Sunrise Tower",
				"premiseLocation": "307/12 Nguyen Van Troi St, W1, Tan Binh",
				"cctvDevice": "CCTV 1",
				"cctvDeviceFloor": 1,
				"isAcknowledged": true,
				"incidentAt": "2023-07-25T00:00:00Z"
			}`,
		},
		"error|failure to get alert id param": {
			alertID:       "40abc",
			expStatusCode: http.StatusBadRequest,
			expRespond: `{
				"error": "invalid_request_body",
				"error_description": "unable to get alert id parameter"
			}`,
		},
		"error|failure request not found": {
			alertID: "40",
			mockAlertsCtrl: struct {
				input  int64
				output model.Alert
				err    error
			}{
				input: 40,
				err:   alerts.ErrNotFound,
			},
			expStatusCode: http.StatusBadRequest,
			expRespond: `{
				"error": "alert_not_found",
				"error_description": "the alert not found"
			}`,
		},
		"error|failure get request got some error": {
			alertID: "40",
			mockAlertsCtrl: struct {
				input  int64
				output model.Alert
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
			chiCtx.URLParams.Add("alertID", tc.alertID)
			ctx := context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx)
			req = req.WithContext(ctx)
			res := httptest.NewRecorder()

			alertsCtrl := new(alerts.MockController)
			alertsCtrl.ExpectedCalls = []*mock.Call{
				alertsCtrl.On("Detail", ctx, tc.mockAlertsCtrl.input).Return(tc.mockAlertsCtrl.output, tc.mockAlertsCtrl.err),
			}

			// When
			instance := New(alertsCtrl)
			handler := instance.GetAlertDetail()
			handler.ServeHTTP(res, req)

			// Then
			require.Equal(t, tc.expStatusCode, res.Code)
			require.JSONEq(t, tc.expRespond, res.Body.String())
		})
	}
}
