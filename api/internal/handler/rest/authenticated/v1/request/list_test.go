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
	"github.com/volatiletech/null/v8"
)

func TestHandler_GetRequests(t *testing.T) {
	tcs := map[string]struct {
		givenInput       string
		mockRequestsCtrl struct {
			input  model.GetRequestsInput
			output []model.RequestSummary
			total  int64
			err    error
		}
		expStatusCode int
		expRespond    string
	}{
		"success|get all requests": {
			givenInput: "premiseID=5&limit=10&page=1",
			mockRequestsCtrl: struct {
				input  model.GetRequestsInput
				output []model.RequestSummary
				total  int64
				err    error
			}{
				input: model.GetRequestsInput{
					PremiseID: 5,
					Limit:     10,
					Page:      1,
				},
				output: []model.RequestSummary{
					{
						ID:              40,
						AlertID:         30,
						Alert:           "CCTV 1-30",
						AlertType:       "Suspicious Activities",
						PremiseName:     "Sunrise Tower",
						PremiseLocation: "307/12 Nguyen Van Troi St, W1, Tan Binh",
						Author:          "John",
						Assignee:        "Thomas",
						Message:         "Property Damage Detected! Please check the monitored area and take screenshots for verification. Report any damage or suspicious activities to the authorities.",
						Status:          "NEW",
						StartTime:       time.Date(2023, 07, 25, 0, 0, 0, 0, time.UTC),
					},
					{
						ID:              41,
						AlertID:         30,
						Alert:           "CCTV 1-30",
						AlertType:       "Suspicious Activities",
						PremiseName:     "Sunrise Tower",
						PremiseLocation: "307/12 Nguyen Van Troi St, W1, Tan Binh",
						Author:          "John",
						Assignee:        "Thomas",
						Message:         "Property Damage Detected! Please check the monitored area and take screenshots for verification. Report any damage or suspicious activities to the authorities.",
						Status:          "RESOLVED",
						StartTime:       time.Date(2023, 07, 26, 0, 0, 0, 0, time.UTC),
						VerifiedAt:      null.TimeFrom(time.Date(2023, 07, 27, 0, 0, 0, 0, time.UTC)),
					},
				},
				total: 2,
			},
			expStatusCode: http.StatusOK,
			expRespond:    `{"items":[{"id":"40","alertID":"30","alert":"CCTV 1-30","alertType":"Suspicious Activities","premiseName":"Sunrise Tower","premiseLocation":"307/12 Nguyen Van Troi St, W1, Tan Binh","author":"John","assignee":"Thomas","status":"NEW","startTime":"2023-07-25T00:00:00Z","verifiedAt":null},{"id":"41","alertID":"30","alert":"CCTV 1-30","alertType":"Suspicious Activities","premiseName":"Sunrise Tower","premiseLocation":"307/12 Nguyen Van Troi St, W1, Tan Binh","author":"John","assignee":"Thomas","status":"RESOLVED","startTime":"2023-07-26T00:00:00Z","verifiedAt":"2023-07-27T00:00:00Z"}],"pagination":{"totalCount":2,"currentPage":1,"limit":10}}`,
		},
		"success|get all requests for assignee": {
			givenInput: "premiseID=5&assigneeID=101&limit=10&page=1",
			mockRequestsCtrl: struct {
				input  model.GetRequestsInput
				output []model.RequestSummary
				total  int64
				err    error
			}{
				input: model.GetRequestsInput{
					PremiseID:  5,
					AssigneeID: 101,
					Limit:      10,
					Page:       1,
				},
				output: []model.RequestSummary{
					{
						ID:              40,
						AlertID:         30,
						Alert:           "CCTV 1-30",
						AlertType:       "Suspicious Activities",
						PremiseName:     "Sunrise Tower",
						PremiseLocation: "307/12 Nguyen Van Troi St, W1, Tan Binh",
						Author:          "John",
						Assignee:        "Thomas",
						Message:         "Property Damage Detected! Please check the monitored area and take screenshots for verification. Report any damage or suspicious activities to the authorities.",
						Status:          "NEW",
						StartTime:       time.Date(2023, 07, 25, 0, 0, 0, 0, time.UTC),
					},
					{
						ID:              41,
						AlertID:         30,
						Alert:           "CCTV 1-30",
						AlertType:       "Suspicious Activities",
						PremiseName:     "Sunrise Tower",
						PremiseLocation: "307/12 Nguyen Van Troi St, W1, Tan Binh",
						Author:          "John",
						Assignee:        "Thomas",
						Message:         "Property Damage Detected! Please check the monitored area and take screenshots for verification. Report any damage or suspicious activities to the authorities.",
						Status:          "RESOLVED",
						StartTime:       time.Date(2023, 07, 26, 0, 0, 0, 0, time.UTC),
						VerifiedAt:      null.TimeFrom(time.Date(2023, 07, 27, 0, 0, 0, 0, time.UTC)),
					},
				},
				total: 2,
			},
			expStatusCode: http.StatusOK,
			expRespond:    `{"items":[{"id":"40","alertID":"30","alert":"CCTV 1-30","alertType":"Suspicious Activities","premiseName":"Sunrise Tower","premiseLocation":"307/12 Nguyen Van Troi St, W1, Tan Binh","author":"John","assignee":"Thomas","status":"NEW","startTime":"2023-07-25T00:00:00Z","verifiedAt":null},{"id":"41","alertID":"30","alert":"CCTV 1-30","alertType":"Suspicious Activities","premiseName":"Sunrise Tower","premiseLocation":"307/12 Nguyen Van Troi St, W1, Tan Binh","author":"John","assignee":"Thomas","status":"RESOLVED","startTime":"2023-07-26T00:00:00Z","verifiedAt":"2023-07-27T00:00:00Z"}],"pagination":{"totalCount":2,"currentPage":1,"limit":10}}`,
		},
		"error|failure invalid premise id": {
			givenInput:    "premiseID=aaa&limit=10&page=1",
			expStatusCode: http.StatusBadRequest,
			expRespond: `{
				"error": "validation_failed",
				"error_description": "invalid premise id"
			}`,
		},
		"error|failure invalid assignee id": {
			givenInput:    "premiseID=5&assigneeID=aaa&limit=10&page=1",
			expStatusCode: http.StatusBadRequest,
			expRespond: `{
				"error": "validation_failed",
				"error_description": "invalid assignee id"
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
			mockRequestsCtrl: struct {
				input  model.GetRequestsInput
				output []model.RequestSummary
				total  int64
				err    error
			}{
				input: model.GetRequestsInput{
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
			req := httptest.NewRequest(http.MethodGet, "/api/authenticated/v1/requests?"+tc.givenInput, nil)
			req.Header.Set("Content-Type", "application/json")

			ctx := context.WithValue(req.Context(), chi.RouteCtxKey, chi.NewRouteContext())
			req = req.WithContext(ctx)

			res := httptest.NewRecorder()

			requestsCtrl := new(requests.MockController)
			requestsCtrl.ExpectedCalls = []*mock.Call{
				requestsCtrl.On("List", ctx, tc.mockRequestsCtrl.input).Return(tc.mockRequestsCtrl.output, tc.mockRequestsCtrl.total, tc.mockRequestsCtrl.err),
			}

			// When
			instance := New(requestsCtrl)
			handler := instance.GetRequests()
			handler.ServeHTTP(res, req)

			// Then
			require.Equal(t, tc.expStatusCode, res.Code)
			require.JSONEq(t, tc.expRespond, res.Body.String())
		})
	}
}
