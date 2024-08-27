package alert

import (
	"context"

	"github.com/nhan1603/CloneScc/api/internal/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// Repository provides the specification of the functionality provided by this pkg
type Repository interface {
	// GetAlerts returns the alerts list
	GetAlerts(context.Context, model.GetAlertsInput) ([]model.Alert, int64, error)
	// GetAlert returns the alert detail
	GetAlert(context.Context, int64) (model.Alert, error)
	// CreateAlert insert a new alert record in the database
	CreateAlert(ctx context.Context, dataModel model.AlertMessage, cctvDeviceID int64) (int64, error)
}

// New returns an implementation instance satisfying Repository
func New(dbConn boil.ContextExecutor) Repository {
	return impl{dbConn: dbConn}
}

type impl struct {
	dbConn boil.ContextExecutor
}
