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

func TestController_impl_Detail(t *testing.T) {
	type mockAlertRepo struct {
		expCall bool
		input   int64
		output  model.Alert
		err     error
	}
	tcs := map[string]struct {
		alertID       int64
		mockAlertRepo mockAlertRepo
		expRes        model.Alert
		expErr        error
	}{
		"success": {
			alertID: 30,
			mockAlertRepo: mockAlertRepo{
				expCall: true,
				input:   30,
				output: model.Alert{
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
			},
			expRes: model.Alert{
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
		},
		"alert not found": {
			alertID: 40,
			mockAlertRepo: mockAlertRepo{
				expCall: true,
				input:   40,
				err:     alert.ErrNotFound,
			},
			expErr: ErrNotFound,
		},
		"repo return something error": {
			alertID: 40,
			mockAlertRepo: mockAlertRepo{
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
			mockAlertRep := new(alert.MockRepository)
			mockAlertRep.ExpectedCalls = []*mock.Call{}
			if tc.mockAlertRepo.expCall {
				mockAlertRep.On("GetAlert", ctx, tc.mockAlertRepo.input).Return(tc.mockAlertRepo.output, tc.mockAlertRepo.err)
			}

			mockReg := new(repository.MockRegistry)
			mockReg.ExpectedCalls = []*mock.Call{
				mockReg.On("Alert").Return(mockAlertRep),
			}

			got, err := New(mockReg, nil, nil).Detail(ctx, tc.alertID)
			if tc.expErr != nil {
				require.EqualError(t, tc.expErr, err.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expRes, got)
			}
		})
	}
}
