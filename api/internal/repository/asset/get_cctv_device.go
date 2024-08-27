package asset

import (
	"context"
	"database/sql"
	"errors"
	"strconv"

	"github.com/nhan1603/CloneScc/api/internal/model"
	"github.com/nhan1603/CloneScc/api/internal/repository/dbmodel"
	pkgerrors "github.com/pkg/errors"
)

// GetCCTVKey returns the primary key of the cctv
func (i impl) GetCCTVKeyByName(ctx context.Context, cctvName string) (int64, error) {
	res, err := dbmodel.CCTVDevices(dbmodel.CCTVDeviceWhere.DeviceName.EQ(cctvName)).One(ctx, i.dbConn)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, ErrCctvNotFound
		}
		return 0, pkgerrors.WithStack(err)
	}

	return res.ID, nil
}

// GetCCTVKey returns the primary key of the cctv
func (i impl) GetAllCCTV(ctx context.Context) ([]model.CctvData, error) {
	result := []model.CctvData{}
	res, err := dbmodel.CCTVDevices().All(ctx, i.dbConn)

	if err != nil {
		return result, pkgerrors.WithStack(err)
	}

	for _, cctv := range res {
		result = append(result, model.CctvData{
			CctvName:    cctv.DeviceName,
			FloorNumber: strconv.Itoa(cctv.FloorNumber.Int),
		})
	}

	return result, nil
}
