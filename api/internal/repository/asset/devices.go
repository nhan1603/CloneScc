package asset

import (
	"context"
	"fmt"

	"github.com/nhan1603/CloneScc/api/internal/model"
	"github.com/nhan1603/CloneScc/api/internal/repository/dbmodel"
	pkgerrors "github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// GetDevicesInput represents for each item of an input
type GetDevicesInput struct {
	Name        string
	PremiseID   int64
	Limit, Page int
}

// GetDevices returns list of premises with provided parameters
func (i impl) GetDevices(ctx context.Context, input GetDevicesInput) ([]model.Devices, int64, error) {
	qms := []qm.QueryMod{
		qm.Load(dbmodel.CCTVDeviceRels.Premise),
	}

	if input.Name != "" {
		qms = append(qms, qm.Where(dbmodel.CCTVDeviceColumns.DeviceName+" ILIKE ?", fmt.Sprintf("%%%s%%", input.Name)))
	}

	if input.PremiseID > 0 {
		qms = append(qms, dbmodel.CCTVDeviceWhere.PremiseID.EQ(input.PremiseID))
	}

	totalCount, err := dbmodel.CCTVDevices(qms...).Count(ctx, i.dbConn)
	if err != nil {
		return nil, 0, err
	}

	if input.Limit > 0 && input.Page > 0 {
		qms = append(qms,
			qm.Offset(input.Limit*(input.Page-1)),
			qm.Limit(input.Limit),
		)
	}

	qms = append(qms, qm.OrderBy(dbmodel.CCTVDeviceColumns.DeviceName))

	devices, err := dbmodel.CCTVDevices(qms...).All(ctx, i.dbConn)
	if err != nil {
		return nil, 0, pkgerrors.WithStack(err)
	}

	if len(devices) == 0 {
		return nil, totalCount, nil
	}

	// Translate the ORM object to internal model object
	result := make([]model.Devices, len(devices))
	for idx, record := range devices {
		if record.R == nil || record.R.Premise == nil {
			return nil, 0, ErrCanNotLoadPremisesTable
		}

		result[idx] = model.Devices{
			ID:              record.ID,
			PremiseID:       record.PremiseID,
			PremiseName:     record.R.Premise.Name,
			PremiseLocation: record.R.Premise.Location,
			DeviceName:      record.DeviceName,
			DeviceCode:      record.DeviceCode,
			IsActive:        record.IsActive,
			FloorNumber:     record.FloorNumber.Int,
		}
	}

	return result, totalCount, nil
}
