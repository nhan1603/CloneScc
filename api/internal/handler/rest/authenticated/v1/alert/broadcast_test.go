package alert

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
	alertCtrl "github.com/nhan1603/CloneScc/api/internal/controller/alerts"
	"github.com/nhan1603/CloneScc/api/internal/model"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHandler_BroadCastAlert(t *testing.T) {
	type mockAlertCtrl struct {
		expCall bool
	}
	tcs := map[string]struct {
		mockAssetCtrl mockAlertCtrl
		expCode       int
		expRes        string
	}{
		"success": {
			mockAssetCtrl: mockAlertCtrl{
				expCall: true,
			},
			expCode: http.StatusOK,
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			// Given
			serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				upgrader := websocket.Upgrader{}
				conn, err := upgrader.Upgrade(w, r, nil)
				if err != nil {
					t.Fatalf("Failed to upgrade WebSocket connection: %v", err)
				}
				defer conn.Close()

				mockCtrl := new(alertCtrl.MockController)
				if tc.mockAssetCtrl.expCall {
					mockCtrl.ExpectedCalls = []*mock.Call{
						mockCtrl.On("BroadCast", context.Background(), conn),
					}
				}

				h := Handler{alertCtrl: mockCtrl}
				handler := http.HandlerFunc(h.BroadCastAlert())
				handler.ServeHTTP(w, r)
				res := httptest.NewRecorder()

				// Then
				mockCtrl.AssertExpectations(t)
				require.Equal(t, res.Code, tc.expCode)
			}))
			defer serv.Close()
		})
	}
}

func TestHandler_PushAlert(t *testing.T) {
	type mockPushAlertCtrl struct {
		expCall bool
		input   model.AlertMessage
	}
	tcs := map[string]struct {
		mockPushAlertCtrl mockPushAlertCtrl
		givenInput        string
		expStatusCD       int
		expRes            string
	}{
		"success": {
			mockPushAlertCtrl: mockPushAlertCtrl{
				expCall: true,
			},
			givenInput:  `{}`,
			expStatusCD: http.StatusOK,
		},
		"error when json Decode": {
			givenInput:  `{`,
			expRes:      `{"error":"internal_error","error_description":"Something went wrong"}`,
			expStatusCD: http.StatusInternalServerError,
		},
	}

	for scenario, tc := range tcs {
		t.Run(scenario, func(t *testing.T) {
			// Setup
			req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(tc.givenInput))
			req.Header.Set("Content-Type", "application/json")

			ctx := context.WithValue(req.Context(), chi.RouteCtxKey, chi.NewRouteContext())
			req = req.WithContext(ctx)

			res := httptest.NewRecorder()

			mockAlertCtrl := new(alertCtrl.MockController)
			if tc.mockPushAlertCtrl.expCall {
				mockAlertCtrl.ExpectedCalls = append(mockAlertCtrl.ExpectedCalls, []*mock.Call{
					mockAlertCtrl.On("Push", ctx, tc.mockPushAlertCtrl.input),
				}...)
			}

			// When
			instance := New(mockAlertCtrl)
			handler := instance.PushAlert()
			handler.ServeHTTP(res, req)
			// Then
			require.Equal(t, tc.expStatusCD, res.Code)
			require.Equal(t, tc.expRes, res.Body.String())
		})
	}
}
