package request

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/websocket"
	requestCtrl "github.com/nhan1603/CloneScc/api/internal/controller/requests"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHandler_BroadCastResponse(t *testing.T) {
	type mockResponseCtrl struct {
		expCall bool
	}
	tcs := map[string]struct {
		mockAssetCtrl mockResponseCtrl
		expCode       int
		expRes        string
	}{
		"success": {
			mockAssetCtrl: mockResponseCtrl{
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

				mockCtrl := new(requestCtrl.MockController)
				if tc.mockAssetCtrl.expCall {
					mockCtrl.ExpectedCalls = []*mock.Call{
						mockCtrl.On("BroadCastResponse", context.Background(), conn),
					}
				}

				h := Handler{requestCtrl: mockCtrl}
				handler := http.HandlerFunc(h.BroadCastResponse())
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
