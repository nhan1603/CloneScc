package alert

import (
	"github.com/nhan1603/CloneScc/api/internal/model"
	"github.com/nhan1603/CloneScc/api/internal/repository/dbmodel"
)

func toAlert(a *dbmodel.Alert) model.Alert {
	alert := model.Alert{
		ID:             a.ID,
		CCTVDeviceID:   a.CCTVDeviceID,
		Type:           a.Type,
		Description:    a.Description,
		IsAcknowledged: a.IsAcknowledged,
		IncidentAt:     a.IncidentAt,
		CreatedAt:      a.CreatedAt,
		UpdatedAt:      a.UpdatedAt,
	}

	if a.R != nil && a.R.CCTVDevice != nil && a.R.CCTVDevice.R.Premise != nil {
		alert.CCTVDevice = a.R.CCTVDevice.DeviceName
		alert.CCTVDeviceFloor = a.R.CCTVDevice.FloorNumber.Int
		alert.PremiseName = a.R.CCTVDevice.R.Premise.Name
		alert.PremiseLocation = a.R.CCTVDevice.R.Premise.Location
	}

	return alert
}
