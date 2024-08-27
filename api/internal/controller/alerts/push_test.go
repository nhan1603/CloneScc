package alerts

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/nhan1603/CloneScc/api/internal/model"
	kafkaPkg "github.com/nhan1603/CloneScc/api/internal/pkg/kafka"
	"github.com/nhan1603/CloneScc/api/internal/repository"
	"github.com/nhan1603/CloneScc/api/internal/repository/alert"
	"github.com/nhan1603/CloneScc/api/internal/repository/asset"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestController_Push(t *testing.T) {
	tcs := map[string]struct {
		givenInput model.AlertMessage
		expErr     error
	}{
		"success": {
			givenInput: model.AlertMessage{
				CCTVName: "test",
			},
		},
	}

	for scenario, tc := range tcs {
		t.Run(scenario, func(t *testing.T) {
			// Given
			ctx := context.Background()

			// When
			instance := New(nil, nil, nil)

			// Then
			require.NotPanics(t, func() { instance.Push(ctx, tc.givenInput) })
		})
	}
}

func TestController_impl_HandleMessage(t *testing.T) {
	type mockGetCCTVRepo struct {
		expCall bool
		input   string
		output  int64
		err     error
	}
	type mockCreateAlertRepo struct {
		expCall      bool
		cctvDeviceID int64
		dataMessage  model.AlertMessage
		output       int64
		err          error
	}

	type mockRepo struct {
		mockGetCCTVRepo     mockGetCCTVRepo
		mockCreateAlertRepo mockCreateAlertRepo
	}

	tcs := map[string]struct {
		input    kafkaPkg.ConsumerMessage
		mockRepo mockRepo
		expErr   error
	}{
		"success": {
			input: kafkaPkg.ConsumerMessage{
				ID: kafkaPkg.ConsumerMessageID{
					Topic:     "TEST",
					Partition: 10,
					Offset:    1,
					Key:       "TEST",
				},
				Value: []byte(`{"id": "30", "cctvName": "CCTV 1","floorNumber": "1", "type": "Property Damage","description": "Test","incidentAt": "2023-08-04T00:00:00Z"}`),
			},
			mockRepo: mockRepo{
				mockGetCCTVRepo: mockGetCCTVRepo{
					expCall: true,
					input:   "CCTV 1",
					output:  100,
				},
				mockCreateAlertRepo: mockCreateAlertRepo{
					expCall:      true,
					cctvDeviceID: 100,
					dataMessage: model.AlertMessage{
						ID:          "30",
						CCTVName:    "CCTV 1",
						FloorNumber: "1",
						Type:        model.AlertTypePropertyDamage.ToString(),
						Description: "Test",
						IncidentAt:  time.Date(2023, 8, 4, 0, 0, 0, 0, time.UTC),
					},
					output: 100,
				},
			},
		},
		"error|create alert|not found cctv device": {
			input: kafkaPkg.ConsumerMessage{
				ID: kafkaPkg.ConsumerMessageID{
					Topic:     "TEST",
					Partition: 10,
					Offset:    1,
					Key:       "TEST",
				},
				Value: []byte(`{"id": "30", "cctvName": "CCTV 1","floorNumber": "1", "type": "Property Damage","description": "Test","incidentAt": "2023-08-04T00:00:00Z"}`),
			},
			mockRepo: mockRepo{
				mockGetCCTVRepo: mockGetCCTVRepo{
					expCall: true,
					input:   "CCTV 1",
					err:     ErrCCTVNotFound,
				},
			},
			expErr: ErrCCTVNotFound,
		},
		"error|create alert|find cctv device got error": {
			input: kafkaPkg.ConsumerMessage{
				ID: kafkaPkg.ConsumerMessageID{
					Topic:     "TEST",
					Partition: 10,
					Offset:    1,
					Key:       "TEST",
				},
				Value: []byte(`{"id": "30", "cctvName": "CCTV 1","floorNumber": "1", "type": "Property Damage","description": "Test","incidentAt": "2023-08-04T00:00:00Z"}`),
			},
			mockRepo: mockRepo{
				mockGetCCTVRepo: mockGetCCTVRepo{
					expCall: true,
					input:   "CCTV 1",
					err:     errors.New("some errors"),
				},
			},
			expErr: errors.New("some errors"),
		},
		"error|create alert got error": {
			input: kafkaPkg.ConsumerMessage{
				ID: kafkaPkg.ConsumerMessageID{
					Topic:     "TEST",
					Partition: 10,
					Offset:    1,
					Key:       "TEST",
				},
				Value: []byte(`{"id": "30", "cctvName": "CCTV 1","floorNumber": "1", "type": "Property Damage","description": "Test","incidentAt": "2023-08-04T00:00:00Z"}`),
			},
			mockRepo: mockRepo{
				mockGetCCTVRepo: mockGetCCTVRepo{
					expCall: true,
					input:   "CCTV 1",
					output:  100,
				},
				mockCreateAlertRepo: mockCreateAlertRepo{
					expCall:      true,
					cctvDeviceID: 100,
					dataMessage: model.AlertMessage{
						ID:          "30",
						CCTVName:    "CCTV 1",
						FloorNumber: "1",
						Type:        model.AlertTypePropertyDamage.ToString(),
						Description: "Test",
						IncidentAt:  time.Date(2023, 8, 4, 0, 0, 0, 0, time.UTC),
					},
					err: errors.New("some errors"),
				},
			},
			expErr: ErrCreateAlert,
		},
	}
	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			ctx := context.Background()

			mockAssetRep := new(asset.MockRepository)
			mockAssetRep.ExpectedCalls = []*mock.Call{}

			if tc.mockRepo.mockGetCCTVRepo.expCall {
				mockAssetRep.On("GetCCTVKeyByName", ctx, tc.mockRepo.mockGetCCTVRepo.input).Return(tc.mockRepo.mockGetCCTVRepo.output, tc.mockRepo.mockGetCCTVRepo.err)
			}

			mockAlertRep := new(alert.MockRepository)
			mockAlertRep.ExpectedCalls = []*mock.Call{}
			if tc.mockRepo.mockCreateAlertRepo.expCall {
				mockAlertRep.On("CreateAlert", ctx, tc.mockRepo.mockCreateAlertRepo.dataMessage, tc.mockRepo.mockCreateAlertRepo.cctvDeviceID).Return(tc.mockRepo.mockCreateAlertRepo.output, tc.mockRepo.mockCreateAlertRepo.err)
			}

			mockReg := new(repository.MockRegistry)
			mockReg.ExpectedCalls = []*mock.Call{
				mockReg.On("Asset").Return(mockAssetRep),
				mockReg.On("Alert").Return(mockAlertRep),
			}

			err := New(mockReg, nil, nil).HandleMessage(ctx, tc.input)
			if tc.expErr != nil {
				require.EqualError(t, tc.expErr, err.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}
