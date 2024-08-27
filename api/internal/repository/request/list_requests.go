package request

import (
	"context"
	"fmt"
	"log"

	"github.com/nhan1603/CloneScc/api/internal/model"
	"github.com/nhan1603/CloneScc/api/internal/repository/dbmodel"
	pkgerrors "github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// GetRequests returns the requests list
func (i impl) GetRequests(ctx context.Context, inp model.GetRequestsInput) ([]model.RequestSummary, int64, error) {
	qms := []qm.QueryMod{
		qm.Load(dbmodel.VerificationRequestRels.VerificationRequestResponses),
		qm.Load(dbmodel.VerificationRequestRels.Alert),
		qm.Load(dbmodel.VerificationRequestRels.RequestByUser),
		qm.Load(dbmodel.VerificationRequestRels.AssignedUser),
		qm.Load(fmt.Sprintf("%v.%v", dbmodel.VerificationRequestRels.Alert, dbmodel.AlertRels.CCTVDevice)),
		qm.Load(fmt.Sprintf("%v.%v.%v", dbmodel.VerificationRequestRels.Alert, dbmodel.AlertRels.CCTVDevice, dbmodel.CCTVDeviceRels.Premise)),
	}

	if inp.PremiseID > 0 {
		qms = append(qms,
			qm.InnerJoin(fmt.Sprintf("%s ON %s.%s = %s.%s",
				dbmodel.TableNames.Alerts,
				dbmodel.TableNames.Alerts,
				dbmodel.AlertColumns.ID,
				dbmodel.TableNames.VerificationRequests,
				dbmodel.VerificationRequestColumns.AlertID)),
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

	if inp.AssigneeID > 0 {
		qms = append(qms, dbmodel.VerificationRequestWhere.AssignedUserID.EQ(inp.AssigneeID))
	}

	totalCount, err := dbmodel.VerificationRequests(qms...).Count(ctx, i.dbConn)
	if err != nil {
		return nil, 0, err
	}

	if inp.Limit > 0 && inp.Page > 0 {
		qms = append(qms,
			qm.Offset((inp.Page-1)*inp.Limit),
			qm.Limit(inp.Limit),
		)
	}

	qms = append(qms, qm.OrderBy(dbmodel.VerificationRequestColumns.UpdatedAt+" DESC"))

	requests, err := dbmodel.VerificationRequests(qms...).All(ctx, i.dbConn)
	if err != nil {
		log.Printf("[GetRequests] failed: %+v\n", err.Error())
		return nil, 0, pkgerrors.WithStack(err)
	}

	var result []model.RequestSummary
	for _, r := range requests {
		// Translate the ORM object to internal model object
		if err = RequestOrm(*r).validate(); err != nil {
			return nil, 0, err
		}

		result = append(result, toRequestListing(r))
	}

	return result, totalCount, nil
}
