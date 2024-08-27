package request

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

// GetRequest returns the request detail
func (i impl) GetRequest(ctx context.Context, requestID int64) (model.RequestDetail, error) {
	o, err := dbmodel.VerificationRequests(
		dbmodel.VerificationRequestWhere.ID.EQ(requestID),
		qm.Load(dbmodel.VerificationRequestRels.VerificationRequestResponses),
		qm.Load(dbmodel.VerificationRequestRels.Alert),
		qm.Load(dbmodel.VerificationRequestRels.RequestByUser),
		qm.Load(dbmodel.VerificationRequestRels.AssignedUser),
		qm.Load(fmt.Sprintf("%v.%v", dbmodel.VerificationRequestRels.Alert, dbmodel.AlertRels.CCTVDevice)),
		qm.Load(fmt.Sprintf("%v.%v.%v", dbmodel.VerificationRequestRels.Alert, dbmodel.AlertRels.CCTVDevice, dbmodel.CCTVDeviceRels.Premise)),
	).One(ctx, i.dbConn)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.RequestDetail{}, ErrNotFound
		}
		return model.RequestDetail{}, pkgerrors.WithStack(err)
	}

	if err := RequestOrm(*o).validate(); err != nil {
		return model.RequestDetail{}, pkgerrors.WithStack(err)
	}

	return toRequestDetail(o), nil
}

type RequestOrm dbmodel.VerificationRequest

func (r RequestOrm) validate() error {
	if r.R == nil ||
		r.R.Alert == nil ||
		r.R.RequestByUser == nil ||
		r.R.AssignedUser == nil ||
		r.R.Alert.R.CCTVDevice == nil ||
		r.R.Alert.R.CCTVDevice.R.Premise == nil {
		return ErrCanNotLoadRelatedTablesData
	}

	return nil
}
