package asset

import (
	"context"

	"github.com/nhan1603/CloneScc/api/internal/repository/dbmodel"
	"github.com/nhan1603/CloneScc/api/internal/repository/generator"
	pkgerrors "github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// GetDeviceToken returns a single token of device with provided parameters
func (i impl) GetDeviceToken(ctx context.Context, userID int64) (string, error) {
	res, err := dbmodel.DeviceTokens(dbmodel.DeviceTokenWhere.UserID.EQ(userID)).One(ctx, i.dbConn)
	if err != nil {
		return "", pkgerrors.WithStack(err)
	}

	return res.DeviceToken, nil
}

// UpsertDeviceTokenInput represents an input struct for UpsertDeviceToken
type UpsertDeviceTokenInput struct {
	UserID                int64
	DeviceToken, Platform string
}

// UpsertDeviceToken creates or updates deviceToken
func (i impl) UpsertDeviceToken(ctx context.Context, input UpsertDeviceTokenInput) error {
	user, err := dbmodel.FindUser(ctx, i.dbConn, input.UserID)
	if err != nil {
		return pkgerrors.WithStack(err)
	}

	id, err := generator.DeviceTokenIDSNF.Generate()
	if err != nil {
		return pkgerrors.WithStack(err)
	}

	dtObj := dbmodel.DeviceToken{
		ID:          id,
		UserID:      user.ID,
		DeviceToken: input.DeviceToken,
		Platform:    input.Platform,
	}

	if err := dtObj.Upsert(
		ctx,
		i.dbConn,
		true,
		[]string{dbmodel.DeviceTokenColumns.UserID},
		boil.Whitelist(
			dbmodel.DeviceTokenColumns.DeviceToken,
			dbmodel.DeviceTokenColumns.Platform,
			dbmodel.DeviceTokenColumns.UpdatedAt,
		),
		boil.Infer(),
	); err != nil {
		return pkgerrors.WithStack(err)
	}

	return nil
}
