package alert

import (
	"context"
	"time"

	"github.com/nhan1603/CloneScc/api/internal/model"
	"github.com/nhan1603/CloneScc/api/internal/repository/dbmodel"
	"github.com/nhan1603/CloneScc/api/internal/repository/generator"
	pkgerrors "github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/types"
)

// CreateAlert hande create alert
func (i impl) CreateAlert(ctx context.Context, dataModel model.AlertMessage, cctvDeviceID int64) (int64, error) {
	id, err := generator.AlertIDSNF.Generate()
	if err != nil {
		return 0, pkgerrors.WithStack(err)
	}
	alertRecord := dbmoderAlertConverter(dataModel, id, cctvDeviceID)

	if err := alertRecord.Insert(ctx, i.dbConn, boil.Infer()); err != nil {
		return 0, pkgerrors.WithStack(err)
	}

	return id, nil
}

func dbmoderAlertConverter(dataModel model.AlertMessage, id, cctvDeviceID int64) dbmodel.Alert {
	return dbmodel.Alert{
		ID:             id,
		CCTVDeviceID:   cctvDeviceID,
		Type:           dataModel.Type,
		Description:    dataModel.Description,
		MediaData:      types.JSON("[]"),
		IsAcknowledged: false,
		IncidentAt:     dataModel.IncidentAt,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
}
