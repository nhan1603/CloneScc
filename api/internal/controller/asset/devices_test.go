package asset

import (
	"context"
	"testing"

	"github.com/nhan1603/CloneScc/api/internal/model"
	"github.com/nhan1603/CloneScc/api/internal/repository"
	"github.com/nhan1603/CloneScc/api/internal/repository/asset"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestController_GetDevices(t *testing.T) {
	type mockAssetRepo struct {
		expCall bool
		input   asset.GetDevicesInput
		output  []model.Devices
		total   int64
		err     error
	}

	tcs := map[string]struct {
		input         GetDevicesInput
		mockAssetRepo mockAssetRepo
		expRes        []model.Devices
		expTotal      int64
		expErr        error
	}{
		"success": {
			input: GetDevicesInput{
				Name:      "Camera",
				PremiseID: 50,
			},
			mockAssetRepo: mockAssetRepo{
				expCall: true,
				input: asset.GetDevicesInput{
					Name:      "Camera",
					PremiseID: 50,
				},
				output: []model.Devices{
					{
						ID:          100,
						PremiseID:   50,
						DeviceName:  "Camera 50",
						DeviceCode:  "cctv_cam50",
						IsActive:    true,
						FloorNumber: 1,
					},
					{
						ID:          101,
						PremiseID:   50,
						DeviceName:  "Camera 51",
						DeviceCode:  "cctv_cam51",
						IsActive:    true,
						FloorNumber: 2,
					},
				},
				total: 2,
			},
			expRes: []model.Devices{
				{
					ID:          100,
					PremiseID:   50,
					DeviceName:  "Camera 50",
					DeviceCode:  "cctv_cam50",
					IsActive:    true,
					FloorNumber: 1,
				},
				{
					ID:          101,
					PremiseID:   50,
					DeviceName:  "Camera 51",
					DeviceCode:  "cctv_cam51",
					IsActive:    true,
					FloorNumber: 2,
				},
			},
			expTotal: 2,
		},
	}
	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			// Given:
			ctx := context.Background()
			mockAssetRepo := &asset.MockRepository{}
			if tc.mockAssetRepo.expCall {
				mockAssetRepo.ExpectedCalls = []*mock.Call{
					mockAssetRepo.On("GetDevices",
						ctx,
						tc.mockAssetRepo.input,
					).Return(
						tc.mockAssetRepo.output,
						tc.mockAssetRepo.total,
						tc.mockAssetRepo.err,
					),
				}
			}
			repo := repository.MockRegistry{}
			repo.ExpectedCalls = []*mock.Call{
				repo.On("Asset").Return(mockAssetRepo),
			}
			s := New(&repo)
			// When:
			actRes, actTotal, err := s.GetDevices(ctx, tc.input)
			// Then:
			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, len(tc.expRes), len(actRes))
				require.Equal(t, tc.expRes, actRes)
				require.Equal(t, tc.expTotal, actTotal)
			}
		})
	}
}
