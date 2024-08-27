package asset

import (
	"context"

	"github.com/nhan1603/CloneScc/api/internal/model"
	"github.com/nhan1603/CloneScc/api/internal/repository/asset"
)

// GetDevicesInput is the search input for GetDevices
type GetDevicesInput struct {
	Name        string
	PremiseID   int64
	Limit, Page int
}

// GetDevices get list of premises with provided prameters
func (i impl) GetDevices(ctx context.Context, input GetDevicesInput) ([]model.Devices, int64, error) {
	return i.repo.Asset().GetDevices(ctx, asset.GetDevicesInput{
		Name:      input.Name,
		PremiseID: input.PremiseID,
		Limit:     input.Limit,
		Page:      input.Page,
	})
}
