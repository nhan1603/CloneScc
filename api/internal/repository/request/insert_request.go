package request

import (
	"context"

	"github.com/nhan1603/CloneScc/api/internal/repository/generator"
	pkgerrors "github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/nhan1603/CloneScc/api/internal/model"
)

var generateRequestIDFunc = generateRequestID

func generateRequestID() (int64, error) {
	return generator.RequestIDSNF.Generate()
}

// Insert a new request to database
func (i impl) Insert(ctx context.Context, input model.VerificationRequest) (model.VerificationRequest, error) {
	id, err := generateRequestIDFunc()
	if err != nil {
		return model.VerificationRequest{}, pkgerrors.WithStack(err)
	}
	input.ID = id

	ormModel := toRequestORM(input)
	if err := ormModel.Insert(ctx, i.dbConn, boil.Infer()); err != nil {
		return model.VerificationRequest{}, pkgerrors.WithStack(err)
	}

	input.CreatedAt = ormModel.CreatedAt
	input.UpdatedAt = ormModel.UpdatedAt

	return input, nil
}
