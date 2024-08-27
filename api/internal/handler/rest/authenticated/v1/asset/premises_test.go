package asset

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/nhan1603/CloneScc/api/internal/controller/asset"
	assetCtrl "github.com/nhan1603/CloneScc/api/internal/controller/asset"
	"github.com/nhan1603/CloneScc/api/internal/model"
	"github.com/stretchr/testify/require"
)

func TestHandler_GetPremises(t *testing.T) {
	type mockAssetCtrl struct {
		expCall bool
		input   assetCtrl.GetPremisesInput
		output  []model.Premises
		err     error
	}

	tcs := map[string]struct {
		mockAssetCtrl mockAssetCtrl
		expCode       int
		expRes        string
		expErr        error
	}{
		"success": {
			mockAssetCtrl: mockAssetCtrl{
				expCall: true,
				input: assetCtrl.GetPremisesInput{
					Name: "MyTower",
				},
				output: []model.Premises{
					{
						ID:           50,
						Name:         "Sunrise Tower",
						Location:     "307/12 Nguyen Van Troi St, W1, Tan Binh",
						PremisesCode: "P001",
						Description:  "Sunrise Tower",
						CCTVCount:    4,
					},
					{
						ID:           51,
						Name:         "Bitexco Financial Tower",
						Location:     "2 Hai Ba Trung St, Ben Nghe, District 1",
						PremisesCode: "P002",
						Description:  "Bitexco Financial Tower",
						CCTVCount:    4,
					},
				},
			},
			expCode: http.StatusOK,
			expRes: `{
				"items": [
					{
						"cctv_count": 4,
						"description": "Sunrise Tower",
						"id": 50,
						"location": "307/12 Nguyen Van Troi St, W1, Tan Binh",
						"name": "Sunrise Tower",
						"premises_code": "P001"
					},
					{
						"cctv_count": 4,
						"description": "Bitexco Financial Tower",
						"id": 51,
						"location": "2 Hai Ba Trung St, Ben Nghe, District 1",
						"name": "Bitexco Financial Tower",
						"premises_code": "P002"
					}
				]
			}`,
		},
		"Premise not found": {
			mockAssetCtrl: mockAssetCtrl{
				expCall: true,
				input: assetCtrl.GetPremisesInput{
					Name: "MyTower",
				},
			},
			expCode: http.StatusOK,
			expRes:  `{"items":[]}`,
		},
		"failed with controller error": {
			mockAssetCtrl: mockAssetCtrl{
				expCall: true,
				input: assetCtrl.GetPremisesInput{
					Name: "MyTower",
				},
				err: errors.New("internal error"),
			},
			expCode: http.StatusInternalServerError,
			expRes:  `{"error":"internal_error", "error_description":"Something went wrong"}`,
		},
	}

	for scenario, tc := range tcs {
		t.Run(scenario, func(t *testing.T) {
			// Given
			req := httptest.NewRequest(http.MethodGet, "/api/authenticated/premises?name=MyTower", nil)
			req.Header.Set("Content-Type", "application/json")

			ctx := context.WithValue(req.Context(), chi.RouteCtxKey, chi.NewRouteContext())
			req = req.WithContext(ctx)

			res := httptest.NewRecorder()

			mockAssetCtrl := new(asset.MockController)
			if tc.mockAssetCtrl.expCall {
				mockAssetCtrl.On("GetPremises", ctx, tc.mockAssetCtrl.input).Return(
					tc.mockAssetCtrl.output,
					tc.mockAssetCtrl.err,
				)
			}

			// When
			instance := New(mockAssetCtrl)
			handler := instance.GetPremises()
			handler.ServeHTTP(res, req)

			// Then
			require.Equal(t, tc.expCode, res.Code)
			require.JSONEq(t, tc.expRes, res.Body.String())
		})
	}
}
