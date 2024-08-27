package alerts

import (
	"context"
	"errors"
	"testing"

	"github.com/nhan1603/CloneScc/api/internal/model"
	"github.com/nhan1603/CloneScc/api/internal/repository"
	"github.com/nhan1603/CloneScc/api/internal/repository/alert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestController_impl_List(t *testing.T) {
	type mockAlertRepo struct {
		expCall bool
		input   model.GetAlertsInput
		output  []model.Alert
		total   int64
		err     error
	}
	tcs := map[string]struct {
		input         model.GetAlertsInput
		mockAlertRepo mockAlertRepo
		expRes        []model.Alert
		expTotal      int64
		expErr        error
	}{
		"success": {
			input: model.GetAlertsInput{
				PremiseID: 100,
				Limit:     10,
				Page:      1,
			},
			mockAlertRepo: mockAlertRepo{
				expCall: true,
				input: model.GetAlertsInput{
					PremiseID: 100,
					Limit:     10,
					Page:      1,
				},
				output: []model.Alert{
					{
						ID:              30,
						CCTVDeviceID:    20,
						Type:            "Suspicious Activities",
						Description:     "TEST",
						PremiseName:     "Sunrise Tower",
						PremiseLocation: "307/12 Nguyen Van Troi St, W1, Tan Binh",
						CCTVDevice:      "CCTV 1",
						CCTVDeviceFloor: 1,
						IsAcknowledged:  true,
					},
					{
						ID:              31,
						CCTVDeviceID:    20,
						Type:            "Suspicious Activities",
						Description:     "TEST",
						PremiseName:     "Sunrise Tower",
						PremiseLocation: "307/12 Nguyen Van Troi St, W1, Tan Binh",
						CCTVDevice:      "CCTV 1",
						CCTVDeviceFloor: 1,
						IsAcknowledged:  true,
					},
				},
				total: 2,
			},
			expTotal: 2,
			expRes: []model.Alert{
				{
					ID:              30,
					CCTVDeviceID:    20,
					Type:            "Suspicious Activities",
					Description:     "TEST",
					PremiseName:     "Sunrise Tower",
					PremiseLocation: "307/12 Nguyen Van Troi St, W1, Tan Binh",
					CCTVDevice:      "CCTV 1",
					CCTVDeviceFloor: 1,
					IsAcknowledged:  true,
				},
				{
					ID:              31,
					CCTVDeviceID:    20,
					Type:            "Suspicious Activities",
					Description:     "TEST",
					PremiseName:     "Sunrise Tower",
					PremiseLocation: "307/12 Nguyen Van Troi St, W1, Tan Binh",
					CCTVDevice:      "CCTV 1",
					CCTVDeviceFloor: 1,
					IsAcknowledged:  true,
				},
			},
		},
		"success|without premises": {
			input: model.GetAlertsInput{
				Limit: 10,
				Page:  1,
			},
			mockAlertRepo: mockAlertRepo{
				expCall: true,
				input: model.GetAlertsInput{
					Limit: 10,
					Page:  1,
				},
				output: []model.Alert{
					{
						ID:              30,
						CCTVDeviceID:    20,
						Type:            "Suspicious Activities",
						Description:     "TEST",
						PremiseName:     "Sunrise Tower",
						PremiseLocation: "307/12 Nguyen Van Troi St, W1, Tan Binh",
						CCTVDevice:      "CCTV 1",
						CCTVDeviceFloor: 1,
						IsAcknowledged:  true,
					},
					{
						ID:              31,
						CCTVDeviceID:    20,
						Type:            "Suspicious Activities",
						Description:     "TEST",
						PremiseName:     "Sunrise Tower",
						PremiseLocation: "307/12 Nguyen Van Troi St, W1, Tan Binh",
						CCTVDevice:      "CCTV 1",
						CCTVDeviceFloor: 1,
						IsAcknowledged:  true,
					},
					{
						ID:              32,
						CCTVDeviceID:    21,
						Type:            "Suspicious Activities",
						Description:     "TEST",
						PremiseName:     "Sunrise Tower",
						PremiseLocation: "307/12 Nguyen Van Troi St, W1, Tan Binh",
						CCTVDevice:      "CCTV 2",
						CCTVDeviceFloor: 1,
						IsAcknowledged:  true,
					},
				},
				total: 3,
			},
			expTotal: 3,
			expRes: []model.Alert{
				{
					ID:              30,
					CCTVDeviceID:    20,
					Type:            "Suspicious Activities",
					Description:     "TEST",
					PremiseName:     "Sunrise Tower",
					PremiseLocation: "307/12 Nguyen Van Troi St, W1, Tan Binh",
					CCTVDevice:      "CCTV 1",
					CCTVDeviceFloor: 1,
					IsAcknowledged:  true,
				},
				{
					ID:              31,
					CCTVDeviceID:    20,
					Type:            "Suspicious Activities",
					Description:     "TEST",
					PremiseName:     "Sunrise Tower",
					PremiseLocation: "307/12 Nguyen Van Troi St, W1, Tan Binh",
					CCTVDevice:      "CCTV 1",
					CCTVDeviceFloor: 1,
					IsAcknowledged:  true,
				},
				{
					ID:              32,
					CCTVDeviceID:    21,
					Type:            "Suspicious Activities",
					Description:     "TEST",
					PremiseName:     "Sunrise Tower",
					PremiseLocation: "307/12 Nguyen Van Troi St, W1, Tan Binh",
					CCTVDevice:      "CCTV 2",
					CCTVDeviceFloor: 1,
					IsAcknowledged:  true,
				},
			},
		},
		"empty": {
			input: model.GetAlertsInput{
				PremiseID: 100,
				Limit:     10,
				Page:      1,
			},
			mockAlertRepo: mockAlertRepo{
				expCall: true,
				input: model.GetAlertsInput{
					PremiseID: 100,
					Limit:     10,
					Page:      1,
				},
				total: 0,
			},
			expTotal: 0,
		},
		"error": {
			input: model.GetAlertsInput{
				PremiseID: 100,
				Limit:     10,
				Page:      1,
			},
			mockAlertRepo: mockAlertRepo{
				expCall: true,
				input: model.GetAlertsInput{
					PremiseID: 100,
					Limit:     10,
					Page:      1,
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
			mockAlertRep := new(alert.MockRepository)
			mockAlertRep.ExpectedCalls = []*mock.Call{}
			if tc.mockAlertRepo.expCall {
				mockAlertRep.On("GetAlerts", ctx, tc.mockAlertRepo.input).Return(tc.mockAlertRepo.output, tc.mockAlertRepo.total, tc.mockAlertRepo.err)
			}

			mockReg := new(repository.MockRegistry)
			mockReg.ExpectedCalls = []*mock.Call{
				mockReg.On("Alert").Return(mockAlertRep),
			}

			got, totalCount, err := New(mockReg, nil, nil).List(ctx, tc.input)
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
