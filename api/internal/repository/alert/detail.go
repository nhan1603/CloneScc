package alert

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/nhan1603/CloneScc/api/internal/model"
	"github.com/nhan1603/CloneScc/api/internal/repository/dbmodel"
	pkgerrors "github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// GetAlert returns the alert detail
func (i impl) GetAlert(ctx context.Context, alertID int64) (model.Alert, error) {
	a, err := dbmodel.Alerts(
		dbmodel.AlertWhere.ID.EQ(alertID),
		qm.Load(dbmodel.AlertRels.CCTVDevice),
		qm.Load(fmt.Sprintf("%v.%v", dbmodel.AlertRels.CCTVDevice, dbmodel.CCTVDeviceRels.Premise)),
	).One(ctx, i.dbConn)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Alert{}, ErrNotFound
		}
		return model.Alert{}, pkgerrors.WithStack(err)
	}

	return toAlert(a), nil
}
