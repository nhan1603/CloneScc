package asset

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi"
	assetCtrl "github.com/nhan1603/CloneScc/api/internal/controller/asset"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHandler_UpdateDeviceToken(t *testing.T) {
	type mockAssetCtrl struct {
		expCall bool
		in      assetCtrl.UpdateDeviceTokenInput
		err     error
	}
	tcs := map[string]struct {
		mockAssetCtrl mockAssetCtrl
		givenInput    string
		expCode       int
		expRes        string
		expErr        error
	}{
		"success": {
			givenInput: `{"userID":50,"deviceToken": "ABCD", "platform": "ios"}`,
			mockAssetCtrl: mockAssetCtrl{
				expCall: true,
				in: assetCtrl.UpdateDeviceTokenInput{
					UserID:      50,
					DeviceToken: "ABCD",
					Platform:    "ios",
				},
			},
			expCode: http.StatusOK,
			expRes:  `{"message":true}`,
		},
		"error": {
			givenInput: `{"userID":51,"deviceToken": "ABCDEF", "platform": "ios"}`,
			mockAssetCtrl: mockAssetCtrl{
				expCall: true,
				in: assetCtrl.UpdateDeviceTokenInput{
					UserID:      51,
					DeviceToken: "ABCDEF",
					Platform:    "ios",
				},
				err: errors.New("something went wrong"),
			},
			expRes:  `{"error":"internal_error","error_description":"Something went wrong"}`,
			expCode: http.StatusInternalServerError,
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			// Setup
			req := httptest.NewRequest(http.MethodPost, "/api/authenticated/v1/device-token", strings.NewReader(tc.givenInput))
			req.Header.Set("Content-Type", "application/json")

			ctx := context.WithValue(req.Context(), chi.RouteCtxKey, chi.NewRouteContext())
			req = req.WithContext(ctx)

			res := httptest.NewRecorder()

			// Given
			mockCtrl := new(assetCtrl.MockController)
			if tc.mockAssetCtrl.expCall {
				mockCtrl.ExpectedCalls = []*mock.Call{
					mockCtrl.On("UpdateDeviceToken", ctx, tc.mockAssetCtrl.in).Return(tc.mockAssetCtrl.err),
				}
			}

			// When
			h := Handler{assetCtrl: mockCtrl}
			handler := http.HandlerFunc(h.UpdateDeviceToken())
			handler.ServeHTTP(res, req)

			// Then
			mockCtrl.AssertExpectations(t)
			if tc.expErr != nil {
				require.Equal(t, res.Code, tc.expCode)
				require.Equal(t, tc.expRes, res.Body.String())
			} else {
				require.Equal(t, tc.expCode, res.Code)
				require.Equal(t, tc.expRes, res.Body.String())
			}
		})
	}
}
