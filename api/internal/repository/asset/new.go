// Package asset represents for an asset pkg on repository
package asset

import (
	"context"

	"github.com/nhan1603/CloneScc/api/internal/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// Repository provides the specification of the functionality provided by this pkg
type Repository interface {
	GetPremises(ctx context.Context, input GetPremisesInput) ([]model.Premises, error)
	GetDevices(ctx context.Context, input GetDevicesInput) ([]model.Devices, int64, error)
	GetDeviceToken(ctx context.Context, userID int64) (string, error)
	UpsertDeviceToken(ctx context.Context, input UpsertDeviceTokenInput) error
	GetCCTVKeyByName(ctx context.Context, cctvName string) (int64, error)
	GetAllCCTV(ctx context.Context) ([]model.CctvData, error)
}

// New returns an implementation instance satisfying Repository
func New(dbConn boil.ContextExecutor) Repository {
	return impl{
		dbConn: dbConn,
	}

}

type impl struct {
	dbConn boil.ContextExecutor
}
