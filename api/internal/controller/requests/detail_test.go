package requests

import (
	"context"
	"errors"
	"testing"

	"github.com/nhan1603/CloneScc/api/internal/model"
	"github.com/nhan1603/CloneScc/api/internal/repository"
	"github.com/nhan1603/CloneScc/api/internal/repository/request"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestController_Detail(t *testing.T) {
	type mockRequestRepo struct {
		expCall bool
		input   int64
		output  model.RequestDetail
		err     error
	}
	tcs := map[string]struct {
		requestID       int64
		mockRequestRepo mockRequestRepo
		expRes          model.RequestDetail
		expErr          error
	}{
		"success": {
			requestID: 40,
			mockRequestRepo: mockRequestRepo{
				expCall: true,
				input:   40,
				output: model.RequestDetail{
					ID:       40,
					Title:    "Camera 1-30",
					Author:   "John Doe",
					Assignee: "John Henry",
					Message:  "Property Damage Detected! Please check the monitored area and take screenshots for verification. Report any damage or suspicious activities to the authorities.",
					AlertDetail: model.Alert{
						ID:              30,
						Type:            "Suspicious Activities",
						Description:     "TEST",
						PremiseName:     "Sunrise Tower",
						PremiseLocation: "307/12 Nguyen Van Troi St, W1, Tan Binh",
						CCTVDevice:      "Camera 1",
						CCTVDeviceFloor: 1,
						IsAcknowledged:  true,
					},
					Respond: &model.RequestRespond{
						ID:      50,
						User:    "John Henry",
						Message: "Ive checked the monitored area, and it's a false alarm. No property damage or suspicious activities were found.",
					},
				},
			},
			expRes: model.RequestDetail{
				ID:       40,
				Title:    "Camera 1-30",
				Author:   "John Doe",
				Assignee: "John Henry",
				Message:  "Property Damage Detected! Please check the monitored area and take screenshots for verification. Report any damage or suspicious activities to the authorities.",
				AlertDetail: model.Alert{
					ID:              30,
					Type:            "Suspicious Activities",
					Description:     "TEST",
					PremiseName:     "Sunrise Tower",
					PremiseLocation: "307/12 Nguyen Van Troi St, W1, Tan Binh",
					CCTVDevice:      "Camera 1",
					CCTVDeviceFloor: 1,
					IsAcknowledged:  true,
				},
				Respond: &model.RequestRespond{
					ID:      50,
					User:    "John Henry",
					Message: "Ive checked the monitored area, and it's a false alarm. No property damage or suspicious activities were found.",
				},
			},
		},
		"request not found": {
			requestID: 40,
			mockRequestRepo: mockRequestRepo{
				expCall: true,
				input:   40,
				err:     request.ErrNotFound,
			},
			expErr: ErrNotFound,
		},
		"repo return something error": {
			requestID: 40,
			mockRequestRepo: mockRequestRepo{
				expCall: true,
				input:   40,
				err:     errors.New("something is error"),
			},
			expErr: errors.New("something is error"),
		},
	}
	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			ctx := context.Background()
			mockRequestRep := new(request.MockRepository)
			mockRequestRep.ExpectedCalls = []*mock.Call{}
			if tc.mockRequestRepo.expCall {
				mockRequestRep.On("GetRequest", ctx, tc.mockRequestRepo.input).Return(tc.mockRequestRepo.output, tc.mockRequestRepo.err)
			}

			mockReg := new(repository.MockRegistry)
			mockReg.ExpectedCalls = []*mock.Call{
				mockReg.On("Request").Return(mockRequestRep),
			}

			got, err := New(mockReg, nil, nil, nil).Detail(ctx, tc.requestID)
			if tc.expErr != nil {
				require.EqualError(t, tc.expErr, err.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expRes, got)
			}
		})
	}
}
