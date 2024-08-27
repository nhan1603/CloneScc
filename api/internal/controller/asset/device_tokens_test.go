package asset

import (
	"context"
	"errors"
	"testing"

	"github.com/nhan1603/CloneScc/api/internal/repository"
	assetRepo "github.com/nhan1603/CloneScc/api/internal/repository/asset"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestController_UpdateDeviceToken(t *testing.T) {
	type mockAssetRepo struct {
		expCall bool
		in      assetRepo.UpsertDeviceTokenInput
		err     error
	}
	tcs := map[string]struct {
		input         UpdateDeviceTokenInput
		mockAssetRepo mockAssetRepo
		expErr        error
	}{
		"success": {
			input: UpdateDeviceTokenInput{
				UserID:      50,
				DeviceToken: "ABCD",
				Platform:    "ios",
			},
			mockAssetRepo: mockAssetRepo{
				expCall: true,
				in: assetRepo.UpsertDeviceTokenInput{
					UserID:      50,
					DeviceToken: "ABCD",
					Platform:    "ios",
				},
			},
		},
		"error": {
			input: UpdateDeviceTokenInput{
				UserID:      51,
				DeviceToken: "ABCDEF",
				Platform:    "ios",
			},
			mockAssetRepo: mockAssetRepo{
				expCall: true,
				in: assetRepo.UpsertDeviceTokenInput{
					UserID:      51,
					DeviceToken: "ABCDEF",
					Platform:    "ios",
				},
				err: errors.New("some errors"),
			},
			expErr: errors.New("some errors"),
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			// Given
			ctx := context.Background()
			mockAssetRepo := &assetRepo.MockRepository{}
			if tc.mockAssetRepo.expCall {
				mockAssetRepo.ExpectedCalls = []*mock.Call{
					mockAssetRepo.On("UpsertDeviceToken", ctx, tc.mockAssetRepo.in).Return(tc.mockAssetRepo.err),
				}
			}

			repo := &repository.MockRegistry{}
			repo.ExpectedCalls = []*mock.Call{
				repo.On("Asset").Return(mockAssetRepo),
			}

			// When
			ctrl := New(repo)
			err := ctrl.UpdateDeviceToken(ctx, tc.input)

			// Then
			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}
