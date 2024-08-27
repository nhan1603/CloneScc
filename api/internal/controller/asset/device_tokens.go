package asset

import (
	"context"

	"github.com/nhan1603/CloneScc/api/internal/repository/asset"
)

// UpdateDeviceTokenInput is the input type for the UpdateDeviceToken controller
type UpdateDeviceTokenInput struct {
	UserID                int64
	DeviceToken, Platform string
}

// UpdateDeviceToken updates token of devices with provided prameters
func (i impl) UpdateDeviceToken(ctx context.Context, input UpdateDeviceTokenInput) error {
	return i.repo.Asset().UpsertDeviceToken(ctx, asset.UpsertDeviceTokenInput{
		UserID:      input.UserID,
		DeviceToken: input.DeviceToken,
		Platform:    input.Platform,
	})
}
