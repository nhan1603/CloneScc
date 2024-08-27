package asset

import (
	"context"

	"github.com/nhan1603/CloneScc/api/internal/model"
	"github.com/nhan1603/CloneScc/api/internal/repository"
)

type Controller interface {
	GetPremises(ctx context.Context, input GetPremisesInput) ([]model.Premises, error)
	GetDevices(ctx context.Context, input GetDevicesInput) ([]model.Devices, int64, error)
	UpdateDeviceToken(ctx context.Context, input UpdateDeviceTokenInput) error
}

// New initializes a new Controller instance and returns it
func New(repo repository.Registry) Controller {
	return impl{
		repo: repo,
	}
}

type impl struct {
	repo repository.Registry
}
