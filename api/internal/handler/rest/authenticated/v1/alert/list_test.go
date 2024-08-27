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

func TestHandler_GetAlerts(t *testing.T) {
	tcs := map[string]struct {
		givenInput     string
		mockAlertsCtrl struct {
			input  model.GetAlertsInput
			output []model.Alert
			total  int64
			err    error
		}
		expStatusCode int
		expRespond    string
	}{
		"success": {
			givenInput: "premiseID=5&limit=10&page=1",
			mockAlertsCtrl: struct {
				input  model.GetAlertsInput
				output []model.Alert
				total  int64
				err    error
			}{
				input: model.GetAlertsInput{
					PremiseID: 5,
					Limit:     10,
					Page:      1,
				},
				output: []model.Alert{
					{
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
					{
						ID:              31,
						CCTVDeviceID:    20,
						Type:            "Suspicious Activities",
						Description:     "TEST",
						PremiseName:     "Sunrise Tower",
						PremiseLocation: "307/12 Nguyen Van Troi St, W1, Tan Binh",
						CCTVDevice:      "CCTV 2",
						CCTVDeviceFloor: 1,
						IsAcknowledged:  true,
						IncidentAt:      time.Date(2023, 7, 25, 0, 0, 0, 0, time.UTC),
					},
				},
				total: 2,
			},
			expStatusCode: http.StatusOK,
			expRespond:    `{"items":[{"id":"30","type":"Suspicious Activities","premiseName":"Sunrise Tower","premiseLocation":"307/12 Nguyen Van Troi St, W1, Tan Binh","cctvDevice":"CCTV 1","cctvDeviceFloor":1,"isAcknowledged":true,"incidentAt":"2023-07-25T00:00:00Z"},{"id":"31","type":"Suspicious Activities","premiseName":"Sunrise Tower","premiseLocation":"307/12 Nguyen Van Troi St, W1, Tan Binh","cctvDevice":"CCTV 2","cctvDeviceFloor":1,"isAcknowledged":true,"incidentAt":"2023-07-25T00:00:00Z"}],"pagination":{"totalCount":2,"currentPage":1,"limit":10}}`,
		},
		"error|failure invalid premise id": {
			givenInput:    "premiseID=aaa&limit=10&page=1",
			expStatusCode: http.StatusBadRequest,
			expRespond: `{
				"error": "validation_failed",
				"error_description": "invalid premise id"
			}`,
		},
		"error|failure limit greater than zero ": {
			givenInput:    "premiseID=5&limit=-10&page=1",
			expStatusCode: http.StatusBadRequest,
			expRespond: `{
				"error": "validation_failed",
				"error_description": "limit must be greater than 0"
			}`,
		},
		"error|failure invalid limit": {
			givenInput:    "premiseID=5&limit=aaa&page=1",
			expStatusCode: http.StatusBadRequest,
			expRespond: `{
				"error": "validation_failed",
				"error_description": "invalid limit"
			}`,
		},
		"error|failure over max limit": {
			givenInput:    "premiseID=5&limit=10000&page=1",
			expStatusCode: http.StatusBadRequest,
			expRespond: `{
				"error": "validation_failed",
				"error_description": "limit must be less than 1000"
			}`,
		},
		"error|failure invalid page": {
			givenInput:    "premiseID=5&limit=10&page=aaa",
			expStatusCode: http.StatusBadRequest,
			expRespond: `{
				"error": "validation_failed",
				"error_description": "invalid page"
			}`,
		},
		"error|failure page greater than zero ": {
			givenInput:    "premiseID=5&limit=10&page=-1",
			expStatusCode: http.StatusBadRequest,
			expRespond: `{
				"error": "validation_failed",
				"error_description": "page must be greater than 0"
			}`,
		},
		"error|get list got error": {
			givenInput: "premiseID=5&limit=10&page=1",
			mockAlertsCtrl: struct {
				input  model.GetAlertsInput
				output []model.Alert
				total  int64
				err    error
			}{
				input: model.GetAlertsInput{
					PremiseID: 5,
					Limit:     10,
					Page:      1,
				},
				err: errors.New("something's error"),
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
			//req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(tc.givenInput))
			req := httptest.NewRequest(http.MethodGet, "/api/authenticated/v1/alerts?"+tc.givenInput, nil)
			req.Header.Set("Content-Type", "application/json")

			ctx := context.WithValue(req.Context(), chi.RouteCtxKey, chi.NewRouteContext())
			req = req.WithContext(ctx)

			res := httptest.NewRecorder()

			alertsCtrl := new(alerts.MockController)
			alertsCtrl.ExpectedCalls = []*mock.Call{
				alertsCtrl.On("List", ctx, tc.mockAlertsCtrl.input).Return(tc.mockAlertsCtrl.output, tc.mockAlertsCtrl.total, tc.mockAlertsCtrl.err),
			}

			// When
			instance := New(alertsCtrl)
			handler := instance.GetAlerts()
			handler.ServeHTTP(res, req)

			// Then
			require.Equal(t, tc.expStatusCode, res.Code)
			require.JSONEq(t, tc.expRespond, res.Body.String())
		})
	}
}
