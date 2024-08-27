package request

import (
	"context"
	"database/sql"
	"errors"

	"github.com/nhan1603/CloneScc/api/internal/model"
	"github.com/nhan1603/CloneScc/api/internal/repository/dbmodel"
	pkgerrors "github.com/pkg/errors"
)

// GetByID gets a request by its ID
func (i impl) GetByID(ctx context.Context, id int64) (model.VerificationRequest, error) {
	ormModel, err := dbmodel.VerificationRequests(
		dbmodel.VerificationRequestWhere.ID.EQ(id),
	).One(ctx, i.dbConn)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.VerificationRequest{}, ErrNotFound
		}
		return model.VerificationRequest{}, pkgerrors.WithStack(err)
	}

	return toRequestModel(ormModel), nil
}
