package user

import (
	"context"

	"github.com/nhan1603/CloneScc/api/internal/model"
	"github.com/nhan1603/CloneScc/api/internal/repository/dbmodel"
	pkgerrors "github.com/pkg/errors"
)

// GetAll retrieves all users
func (i impl) GetAll(context.Context) ([]model.User, error) {
	ormModel, err := dbmodel.Users().All(context.Background(), i.dbConn)
	if err != nil {
		return nil, pkgerrors.WithStack(err)
	}
	return toModelUsers(ormModel), nil
}
