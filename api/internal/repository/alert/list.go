package alert

import (
	"context"
	"fmt"
	"log"

	"github.com/nhan1603/CloneScc/api/internal/model"
	"github.com/nhan1603/CloneScc/api/internal/repository/dbmodel"
	pkgerrors "github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// GetAlerts returns the alerts list
func (i impl) GetAlerts(ctx context.Context, inp model.GetAlertsInput) ([]model.Alert, int64, error) {
	qms := []qm.QueryMod{
		qm.Load(dbmodel.AlertRels.CCTVDevice),
		qm.Load(fmt.Sprintf("%v.%v", dbmodel.AlertRels.CCTVDevice, dbmodel.CCTVDeviceRels.Premise)),
	}

	if inp.PremiseID > 0 {
		qms = append(qms,
			qm.InnerJoin(fmt.Sprintf("%s ON %s.%s = %s.%s",
				dbmodel.TableNames.CCTVDevices,
				dbmodel.TableNames.CCTVDevices,
				dbmodel.CCTVDeviceColumns.ID,
				dbmodel.TableNames.Alerts,
				dbmodel.AlertColumns.CCTVDeviceID)),
			qm.InnerJoin(fmt.Sprintf("%s ON %s.%s = %s.%s",
				dbmodel.TableNames.Premises,
				dbmodel.TableNames.Premises,
				dbmodel.PremiseColumns.ID,
				dbmodel.TableNames.CCTVDevices,
				dbmodel.CCTVDeviceColumns.PremiseID)),
			dbmodel.PremiseWhere.ID.EQ(inp.PremiseID),
		)
	}

	totalCount, err := dbmodel.Alerts(qms...).Count(ctx, i.dbConn)
	if err != nil {
		return nil, 0, err
	}

	if inp.Limit > 0 && inp.Page > 0 {
		qms = append(qms,
			qm.Offset((inp.Page-1)*inp.Limit),
			qm.Limit(inp.Limit),
		)
	}

	qms = append(qms, qm.OrderBy(dbmodel.AlertColumns.CreatedAt+" DESC"))

	alerts, err := dbmodel.Alerts(qms...).All(ctx, i.dbConn)
	if err != nil {
		log.Printf("[GetAlerts] failed: %+v\n", err.Error())
		return nil, 0, pkgerrors.WithStack(err)
	}

	var result []model.Alert
	for _, a := range alerts {
		result = append(result, toAlert(a))
	}

	return result, totalCount, nil
}
