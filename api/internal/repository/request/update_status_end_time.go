package request

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/nhan1603/CloneScc/api/internal/model"
	"github.com/nhan1603/CloneScc/api/internal/repository/dbmodel"
	pkgerrors "github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// UpdateStatusAndEndTime updates the status and end time of a request
func (i impl) UpdateStatusAndEndTime(ctx context.Context, reqID int64, status model.VerificationRequestStatus, endTime time.Time) (model.VerificationRequest, error) {
	ormModel, err := dbmodel.VerificationRequests(
		dbmodel.VerificationRequestWhere.ID.EQ(reqID),
		qm.For("UPDATE"),
	).One(ctx, i.dbConn)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.VerificationRequest{}, ErrNotFound
		}
		return model.VerificationRequest{}, pkgerrors.WithStack(err)
	}

	ormModel.Status = status.ToString()
	ormModel.EndTime = null.TimeFrom(endTime)

	if _, err := ormModel.Update(ctx, i.dbConn, boil.Whitelist(
		dbmodel.VerificationRequestColumns.Status,
		dbmodel.VerificationRequestColumns.EndTime,
		dbmodel.VerificationRequestColumns.UpdatedAt,
	)); err != nil {
		return model.VerificationRequest{}, pkgerrors.WithStack(err)
	}

	return toRequestModel(ormModel), nil
}
