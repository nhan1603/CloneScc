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
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHandler_GetDevices(t *testing.T) {
	type mockAssetCtrl struct {
		expCall bool
		input   assetCtrl.GetDevicesInput
		total   int64
		output  []model.Devices
		err     error
	}

	tcs := map[string]struct {
		url           string
		mockAssetCtrl mockAssetCtrl
		expCode       int
		expRes        string
		expErr        error
	}{
		"success": {
			url: "/api/authenticated/devices?name=Camera&premiseID=50&limit=2&page=1",
			mockAssetCtrl: mockAssetCtrl{
				expCall: true,
				input: assetCtrl.GetDevicesInput{
					Name:      "Camera",
					PremiseID: 50,
					Limit:     2,
					Page:      1,
				},
				total: 2,
				output: []model.Devices{
					{
						ID:              100,
						PremiseID:       50,
						PremiseName:     "Sunrise Tower",
						PremiseLocation: "307/12 Nguyen Van Troi St, W1, Tan Binh",
						DeviceName:      "Camera 50",
						DeviceCode:      "cctv_cam50",
						IsActive:        true,
						FloorNumber:     1,
					},
					{
						ID:              101,
						PremiseID:       50,
						PremiseName:     "Sunrise Tower",
						PremiseLocation: "307/12 Nguyen Van Troi St, W1, Tan Binh",
						DeviceName:      "Camera 51",
						DeviceCode:      "cctv_cam51",
						IsActive:        true,
						FloorNumber:     2,
					},
				},
			},
			expCode: http.StatusOK,
			expRes: `{
						"items": [
							{
								"premiseName": "Sunrise Tower",
								"premiseLocation": "307/12 Nguyen Van Troi St, W1, Tan Binh",
								"deviceCode": "cctv_cam50",
								"deviceName": "Camera 50",
								"deviceURL": "",
								"floorNumber": 1,
								"id": 100,
								"isActive": true,
								"premiseID": 50
							},
							{
								"premiseName": "Sunrise Tower",
								"premiseLocation": "307/12 Nguyen Van Troi St, W1, Tan Binh",
								"deviceCode": "cctv_cam51",
								"deviceName": "Camera 51",
								"deviceURL": "",
								"floorNumber": 2,
								"id": 101,
								"isActive": true,
								"premiseID": 50
							}
						],
						"pagination": {
							"currentPage": 1,
							"limit": 2,
							"totalCount": 2
						}
					}`,
		},
		"success with missing limit number": {
			url: "/api/authenticated/devices?name=Camera&premiseID=50&page=1",
			mockAssetCtrl: mockAssetCtrl{
				expCall: true,
				input: assetCtrl.GetDevicesInput{
					Name:      "Camera",
					PremiseID: 50,
					Limit:     20,
					Page:      1,
				},
				total: 2,
				output: []model.Devices{
					{
						ID:              100,
						PremiseID:       50,
						PremiseName:     "Sunrise Tower",
						PremiseLocation: "307/12 Nguyen Van Troi St, W1, Tan Binh",
						DeviceName:      "Camera 50",
						DeviceCode:      "cctv_cam50",
						IsActive:        true,
						FloorNumber:     1,
					},
					{
						ID:              101,
						PremiseID:       50,
						PremiseName:     "Sunrise Tower",
						PremiseLocation: "307/12 Nguyen Van Troi St, W1, Tan Binh",
						DeviceName:      "Camera 51",
						DeviceCode:      "cctv_cam51",
						IsActive:        true,
						FloorNumber:     2,
					},
				},
			},
			expCode: http.StatusOK,
			expRes: `{
						"items": [
							{
								"premiseName": "Sunrise Tower",
								"premiseLocation": "307/12 Nguyen Van Troi St, W1, Tan Binh",
								"deviceCode": "cctv_cam50",
								"deviceName": "Camera 50",
								"deviceURL": "",
								"floorNumber": 1,
								"id": 100,
								"isActive": true,
								"premiseID": 50
							},
							{
								"premiseName": "Sunrise Tower",
								"premiseLocation": "307/12 Nguyen Van Troi St, W1, Tan Binh",
								"deviceCode": "cctv_cam51",
								"deviceName": "Camera 51",
								"deviceURL": "",
								"floorNumber": 2,
								"id": 101,
								"isActive": true,
								"premiseID": 50
							}
						],
						"pagination": {
							"currentPage": 1,
							"limit": 20,
							"totalCount": 2
						}
					}`,
		},
		"success with missing page": {
			url: "/api/authenticated/devices?name=Camera&premiseID=50&limit=2",
			mockAssetCtrl: mockAssetCtrl{
				expCall: true,
				input: assetCtrl.GetDevicesInput{
					Name:      "Camera",
					PremiseID: 50,
					Limit:     2,
					Page:      1,
				},
				total: 2,
				output: []model.Devices{
					{
						ID:              100,
						PremiseID:       50,
						PremiseName:     "Sunrise Tower",
						PremiseLocation: "307/12 Nguyen Van Troi St, W1, Tan Binh",
						DeviceName:      "Camera 50",
						DeviceCode:      "cctv_cam50",
						IsActive:        true,
						FloorNumber:     1,
					},
					{
						ID:              101,
						PremiseID:       50,
						PremiseName:     "Sunrise Tower",
						PremiseLocation: "307/12 Nguyen Van Troi St, W1, Tan Binh",
						DeviceName:      "Camera 51",
						DeviceCode:      "cctv_cam51",
						IsActive:        true,
						FloorNumber:     2,
					},
				},
			},
			expCode: http.StatusOK,
			expRes: `{
						"items": [
							{
								"premiseName": "Sunrise Tower",
								"premiseLocation": "307/12 Nguyen Van Troi St, W1, Tan Binh",
								"deviceCode": "cctv_cam50",
								"deviceName": "Camera 50",
								"deviceURL": "",
								"floorNumber": 1,
								"id": 100,
								"isActive": true,
								"premiseID": 50
							},
							{
								"premiseName": "Sunrise Tower",
								"premiseLocation": "307/12 Nguyen Van Troi St, W1, Tan Binh",
								"deviceCode": "cctv_cam51",
								"deviceName": "Camera 51",
								"deviceURL": "",
								"floorNumber": 2,
								"id": 101,
								"isActive": true,
								"premiseID": 50
							}
						],
						"pagination": {
							"currentPage": 1,
							"limit": 2,
							"totalCount": 2
						}
					}`,
		},
		"failed with controller error": {
			url: "/api/authenticated/devices?name=Camera&premiseID=50&limit=2&page=1",
			mockAssetCtrl: mockAssetCtrl{
				expCall: true,
				input: assetCtrl.GetDevicesInput{
					Name:      "Camera",
					PremiseID: 50,
					Limit:     2,
					Page:      1,
				},
				err: errors.New("internal error"),
			},
			expCode: http.StatusInternalServerError,
			expRes:  `{"error":"internal_error", "error_description":"Something went wrong"}`,
		},
		"failed with invalid premise id format": {
			url:     "/api/authenticated/devices?name=Camera&premiseID=unknown&limit=2&page=1",
			expCode: http.StatusBadRequest,
			expRes:  `{"error":"validation_failed", "error_description":"invalid id"}`,
		},
		"failed with invalid limit number format": {
			url:     "/api/authenticated/devices?name=Camera&premiseID=50&limit=unknown&page=1",
			expCode: http.StatusBadRequest,
			expRes:  `{"error":"validation_failed", "error_description":"invalid number"}`,
		},
		"failed with limit number < 0": {
			url:     "/api/authenticated/devices?name=Camera&premiseID=50&limit=-99&page=1",
			expCode: http.StatusBadRequest,
			expRes:  `{"error":"validation_failed", "error_description":"limit must be greater than 0"}`,
		},
		"failed with limit number greater than page size": {
			url:     "/api/authenticated/devices?name=Camera&premiseID=50&limit=1001&page=1",
			expCode: http.StatusBadRequest,
			expRes:  `{"error":"validation_failed", "error_description":"limit must be less than 1000"}`,
		},
		"failed with invalid page format": {
			url:     "/api/authenticated/devices?name=Camera&premiseID=50&limit=2&page=unknown",
			expCode: http.StatusBadRequest,
			expRes:  `{"error":"validation_failed", "error_description":"invalid number"}`,
		},
		"failed with page < 0": {
			url:     "/api/authenticated/devices?name=Camera&premiseID=50&limit=2&page=-99",
			expCode: http.StatusBadRequest,
			expRes:  `{"error":"validation_failed", "error_description":"page must be greater than 0"}`,
		},
	}
	for scenario, tc := range tcs {
		t.Run(scenario, func(t *testing.T) {
			// Given
			req := httptest.NewRequest(http.MethodGet, tc.url, nil)
			req.Header.Set("Content-Type", "application/json")

			ctx := context.WithValue(req.Context(), chi.RouteCtxKey, chi.NewRouteContext())
			req = req.WithContext(ctx)

			res := httptest.NewRecorder()

			mockAssetCtrl := new(asset.MockController)
			if tc.mockAssetCtrl.expCall {
				mockAssetCtrl.On("GetDevices", ctx, mock.Anything).Return(
					tc.mockAssetCtrl.output,
					tc.mockAssetCtrl.total,
					tc.mockAssetCtrl.err,
				)
			}

			// When
			instance := New(mockAssetCtrl)
			handler := instance.GetDevices()
			handler.ServeHTTP(res, req)

			// Then
			require.Equal(t, tc.expCode, res.Code)
			require.JSONEq(t, tc.expRes, res.Body.String())
		})
	}
}
