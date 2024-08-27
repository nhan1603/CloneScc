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

func TestController_impl_List(t *testing.T) {
	type mockRequestRepo struct {
		expCall bool
		input   model.GetRequestsInput
		output  []model.RequestSummary
		total   int64
		err     error
	}
	tcs := map[string]struct {
		input           model.GetRequestsInput
		mockRequestRepo mockRequestRepo
		expRes          []model.RequestSummary
		expTotal        int64
		expErr          error
	}{
		"success|get all requests": {
			input: model.GetRequestsInput{
				PremiseID:  100,
				AssigneeID: 0,
				Limit:      10,
				Page:       1,
			},
			mockRequestRepo: mockRequestRepo{
				expCall: true,
				input: model.GetRequestsInput{
					PremiseID:  100,
					AssigneeID: 0,
					Limit:      10,
					Page:       1,
				},
				output: []model.RequestSummary{
					{
						ID:              400,
						AlertID:         300,
						Alert:           "CCTV 1-300",
						AlertType:       "Suspicious Activities",
						PremiseName:     "Sunrise Tower",
						PremiseLocation: "307/12 Nguyen Van Troi St, W1, Tan Binh",
						Author:          "John",
						Assignee:        "Thomas",
						Message:         "Property Damage Detected! Please check the monitored area and take screenshots for verification. Report any damage or suspicious activities to the authorities.",
						Status:          "NEW",
					},
					{
						ID:              401,
						AlertID:         300,
						Alert:           "CCTV 1-300",
						AlertType:       "Suspicious Activities",
						PremiseName:     "Sunrise Tower",
						PremiseLocation: "307/12 Nguyen Van Troi St, W1, Tan Binh",
						Author:          "John",
						Assignee:        "Thomas",
						Message:         "Property Damage Detected! Please check the monitored area and take screenshots for verification. Report any damage or suspicious activities to the authorities.",
						Status:          "NEW",
					},
				},
				total: 2,
			},
			expTotal: 2,
			expRes: []model.RequestSummary{
				{
					ID:              400,
					AlertID:         300,
					Alert:           "CCTV 1-300",
					AlertType:       "Suspicious Activities",
					PremiseName:     "Sunrise Tower",
					PremiseLocation: "307/12 Nguyen Van Troi St, W1, Tan Binh",
					Author:          "John",
					Assignee:        "Thomas",
					Message:         "Property Damage Detected! Please check the monitored area and take screenshots for verification. Report any damage or suspicious activities to the authorities.",
					Status:          "NEW",
				},
				{
					ID:              401,
					AlertID:         300,
					Alert:           "CCTV 1-300",
					AlertType:       "Suspicious Activities",
					PremiseName:     "Sunrise Tower",
					PremiseLocation: "307/12 Nguyen Van Troi St, W1, Tan Binh",
					Author:          "John",
					Assignee:        "Thomas",
					Message:         "Property Damage Detected! Please check the monitored area and take screenshots for verification. Report any damage or suspicious activities to the authorities.",
					Status:          "NEW",
				},
			},
		},
		"success|get all requests for assignee": {
			input: model.GetRequestsInput{
				PremiseID:  100,
				AssigneeID: 101,
				Limit:      10,
				Page:       1,
			},
			mockRequestRepo: mockRequestRepo{
				expCall: true,
				input: model.GetRequestsInput{
					PremiseID:  100,
					AssigneeID: 101,
					Limit:      10,
					Page:       1,
				},
				output: []model.RequestSummary{
					{
						ID:              400,
						AlertID:         300,
						Alert:           "CCTV 1-300",
						AlertType:       "Suspicious Activities",
						PremiseName:     "Sunrise Tower",
						PremiseLocation: "307/12 Nguyen Van Troi St, W1, Tan Binh",
						Author:          "John",
						Assignee:        "Thomas",
						Message:         "Property Damage Detected! Please check the monitored area and take screenshots for verification. Report any damage or suspicious activities to the authorities.",
						Status:          "NEW",
					},
				},
				total: 1,
			},
			expTotal: 1,
			expRes: []model.RequestSummary{
				{
					ID:              400,
					AlertID:         300,
					Alert:           "CCTV 1-300",
					AlertType:       "Suspicious Activities",
					PremiseName:     "Sunrise Tower",
					PremiseLocation: "307/12 Nguyen Van Troi St, W1, Tan Binh",
					Author:          "John",
					Assignee:        "Thomas",
					Message:         "Property Damage Detected! Please check the monitored area and take screenshots for verification. Report any damage or suspicious activities to the authorities.",
					Status:          "NEW",
				},
			},
		},
		"success|without premises": {
			input: model.GetRequestsInput{
				Limit: 10,
				Page:  1,
			},
			mockRequestRepo: mockRequestRepo{
				expCall: true,
				input: model.GetRequestsInput{
					Limit: 10,
					Page:  1,
				},
				output: []model.RequestSummary{
					{
						ID:              400,
						AlertID:         300,
						Alert:           "CCTV 1-300",
						AlertType:       "Suspicious Activities",
						PremiseName:     "Sunrise Tower",
						PremiseLocation: "307/12 Nguyen Van Troi St, W1, Tan Binh",
						Author:          "John",
						Assignee:        "Thomas",
						Message:         "Property Damage Detected! Please check the monitored area and take screenshots for verification. Report any damage or suspicious activities to the authorities.",
						Status:          "NEW",
					},
					{
						ID:              401,
						AlertID:         300,
						Alert:           "CCTV 1-300",
						AlertType:       "Suspicious Activities",
						PremiseName:     "Sunrise Tower",
						PremiseLocation: "307/12 Nguyen Van Troi St, W1, Tan Binh",
						Author:          "John",
						Assignee:        "Thomas",
						Message:         "Property Damage Detected! Please check the monitored area and take screenshots for verification. Report any damage or suspicious activities to the authorities.",
						Status:          "NEW",
					},
				},
				total: 2,
			},
			expTotal: 2,
			expRes: []model.RequestSummary{
				{
					ID:              400,
					AlertID:         300,
					Alert:           "CCTV 1-300",
					AlertType:       "Suspicious Activities",
					PremiseName:     "Sunrise Tower",
					PremiseLocation: "307/12 Nguyen Van Troi St, W1, Tan Binh",
					Author:          "John",
					Assignee:        "Thomas",
					Message:         "Property Damage Detected! Please check the monitored area and take screenshots for verification. Report any damage or suspicious activities to the authorities.",
					Status:          "NEW",
				},
				{
					ID:              401,
					AlertID:         300,
					Alert:           "CCTV 1-300",
					AlertType:       "Suspicious Activities",
					PremiseName:     "Sunrise Tower",
					PremiseLocation: "307/12 Nguyen Van Troi St, W1, Tan Binh",
					Author:          "John",
					Assignee:        "Thomas",
					Message:         "Property Damage Detected! Please check the monitored area and take screenshots for verification. Report any damage or suspicious activities to the authorities.",
					Status:          "NEW",
				},
			},
		},
		"empty": {
			input: model.GetRequestsInput{
				PremiseID:  100,
				AssigneeID: 0,
				Limit:      10,
				Page:       1,
			},
			mockRequestRepo: mockRequestRepo{
				expCall: true,
				input: model.GetRequestsInput{
					PremiseID:  100,
					AssigneeID: 0,
					Limit:      10,
					Page:       1,
				},
				total: 0,
			},
			expTotal: 0,
		},
		"error": {
			input: model.GetRequestsInput{
				PremiseID:  100,
				AssigneeID: 0,
				Limit:      10,
				Page:       1,
			},
			mockRequestRepo: mockRequestRepo{
				expCall: true,
				input: model.GetRequestsInput{
					PremiseID:  100,
					AssigneeID: 0,
					Limit:      10,
					Page:       1,
				},
				err:   errors.New("something error"),
				total: 0,
			},
			expTotal: 0,
			expErr:   errors.New("something error"),
		},
	}
	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			ctx := context.Background()
			mockRequestRep := new(request.MockRepository)
			mockRequestRep.ExpectedCalls = []*mock.Call{}
			if tc.mockRequestRepo.expCall {
				mockRequestRep.On("GetRequests", ctx, tc.mockRequestRepo.input).Return(tc.mockRequestRepo.output, tc.mockRequestRepo.total, tc.mockRequestRepo.err)
			}

			mockReg := new(repository.MockRegistry)
			mockReg.ExpectedCalls = []*mock.Call{
				mockReg.On("Request").Return(mockRequestRep),
			}

			got, totalCount, err := New(mockReg, nil, nil, nil).List(ctx, tc.input)
			if tc.expErr != nil {
				require.EqualError(t, tc.expErr, err.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expTotal, totalCount)
				require.Equal(t, tc.expRes, got)
			}
		})
	}
}
