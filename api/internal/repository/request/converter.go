package request

import (
	"encoding/json"

	pkgerrors "github.com/pkg/errors"
	"github.com/volatiletech/null/v8"

	"github.com/nhan1603/CloneScc/api/internal/model"
	"github.com/nhan1603/CloneScc/api/internal/repository/dbmodel"
)

func toRequestORM(m model.VerificationRequest) dbmodel.VerificationRequest {
	return dbmodel.VerificationRequest{
		ID:             m.ID,
		AlertID:        m.AlertID,
		RequestBy:      m.RequestBy,
		AssignedUserID: m.AssignedUserID,
		Message:        null.StringFrom(m.Message),
		Status:         m.Status.ToString(),
		StartTime:      m.StartTime,
		EndTime:        null.NewTime(m.EndTime, !m.EndTime.IsZero()),
	}
}

func toRequestResponsesORM(m model.VerificationRequestResponses) (dbmodel.VerificationRequestResponse, error) {
	media := null.JSONFrom(nil)
	if m.MediaData != nil {
		mediaBytes, err := json.Marshal(m.MediaData)
		if err != nil {
			return dbmodel.VerificationRequestResponse{}, pkgerrors.WithStack(err)
		}
		media = null.JSONFrom(mediaBytes)
	}
	return dbmodel.VerificationRequestResponse{
		ID:                    m.ID,
		VerificationRequestID: m.VerificationRequestID,
		Message:               m.Message,
		MediaData:             media,
		VerifiedAt:            m.VerifiedAt,
	}, nil
}

func toRequestModel(o *dbmodel.VerificationRequest) model.VerificationRequest {
	return model.VerificationRequest{
		ID:             o.ID,
		AlertID:        o.AlertID,
		RequestBy:      o.RequestBy,
		AssignedUserID: o.AssignedUserID,
		Message:        o.Message.String,
		Status:         model.VerificationRequestStatus(o.Status),
		StartTime:      o.StartTime,
		EndTime:        o.EndTime.Time,
		CreatedAt:      o.CreatedAt,
		UpdatedAt:      o.UpdatedAt,
	}
}
