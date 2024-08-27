package request

import (
	"context"

	"github.com/nhan1603/CloneScc/api/internal/model"
	"github.com/nhan1603/CloneScc/api/internal/repository/generator"
	pkgerrors "github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var generateRequestResponseIDFunc = generateRequestResponseID

func generateRequestResponseID() (int64, error) {
	return generator.RequestIDSNF.Generate()
}

func (i impl) InsertRequestResponses(ctx context.Context, input model.VerificationRequestResponses) (model.VerificationRequestResponses, error) {
	id, err := generateRequestResponseIDFunc()
	if err != nil {
		return model.VerificationRequestResponses{}, pkgerrors.WithStack(err)
	}
	input.ID = id

	ormModel, err := toRequestResponsesORM(input)
	if err != nil {
		return model.VerificationRequestResponses{}, err
	}
	if err := ormModel.Insert(ctx, i.dbConn, boil.Infer()); err != nil {
		return model.VerificationRequestResponses{}, pkgerrors.WithStack(err)
	}

	input.CreatedAt = ormModel.CreatedAt
	input.UpdatedAt = ormModel.UpdatedAt

	return input, nil
}
