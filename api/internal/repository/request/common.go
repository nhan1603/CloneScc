package request

import (
	"strconv"

	"github.com/nhan1603/CloneScc/api/internal/model"
	"github.com/nhan1603/CloneScc/api/internal/repository/dbmodel"
	"github.com/volatiletech/null/v8"
)

func toRequestDetail(r *dbmodel.VerificationRequest) model.RequestDetail {
	requestD := model.RequestDetail{
		ID:        r.ID,
		Message:   r.Message.String,
		StartTime: r.StartTime,
	}

	requestD.Title = r.R.Alert.R.CCTVDevice.DeviceName + "-" + strconv.FormatInt(r.R.Alert.ID, 10)
	requestD.Author = r.R.RequestByUser.DisplayName
	requestD.Assignee = r.R.AssignedUser.DisplayName

	// Get alert detail
	if r.R != nil && r.R.Alert != nil {
		a := r.R.Alert
		requestD.AlertDetail = model.Alert{
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
			requestD.AlertDetail.CCTVDevice = a.R.CCTVDevice.DeviceName
			requestD.AlertDetail.CCTVDeviceFloor = a.R.CCTVDevice.FloorNumber.Int
			requestD.AlertDetail.PremiseName = a.R.CCTVDevice.R.Premise.Name
			requestD.AlertDetail.PremiseLocation = a.R.CCTVDevice.R.Premise.Location
		}
	}

	// Get request respond
	if r.R != nil && r.R.VerificationRequestResponses != nil {
		respond := r.R.VerificationRequestResponses[0]
		requestD.Respond = &model.RequestRespond{
			ID:         respond.ID,
			User:       r.R.AssignedUser.DisplayName, //User in respond is assignee user
			Message:    respond.Message,
			MediaData:  respond.MediaData.JSON,
			VerifiedAt: respond.VerifiedAt,
		}
	}

	return requestD
}

func toRequestListing(r *dbmodel.VerificationRequest) model.RequestSummary {
	request := model.RequestSummary{
		ID:        r.ID,
		AlertID:   r.AlertID,
		Message:   r.Message.String,
		StartTime: r.StartTime,
		EndTime:   r.EndTime.Time,
	}

	request.Alert = r.R.Alert.R.CCTVDevice.DeviceName + "-" + strconv.FormatInt(r.R.Alert.ID, 10)
	request.AlertType = r.R.Alert.Type
	request.PremiseName = r.R.Alert.R.CCTVDevice.R.Premise.Name
	request.PremiseLocation = r.R.Alert.R.CCTVDevice.R.Premise.Location
	request.Author = r.R.RequestByUser.DisplayName
	request.Assignee = r.R.AssignedUser.DisplayName
	request.Status = r.Status

	if r.R.VerificationRequestResponses != nil {
		request.VerifiedAt = null.TimeFrom(r.R.VerificationRequestResponses[0].VerifiedAt)
	}

	return request
}
