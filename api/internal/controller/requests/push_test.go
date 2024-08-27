package requests

import (
	"context"
	"testing"

	"github.com/nhan1603/CloneScc/api/internal/model"
	"github.com/nhan1603/CloneScc/api/internal/repository"
	"github.com/nhan1603/CloneScc/api/internal/repository/alert"
	"github.com/nhan1603/CloneScc/api/internal/repository/request"
	mock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestController_Push(t *testing.T) {
	tcs := map[string]struct {
		givenInput  int64
		returnData  model.VerificationRequest
		returnAlert model.Alert
		expErr      error
	}{
		"success": {
			givenInput: 12,
			returnData: model.VerificationRequest{
				AlertID: 1,
			},
			returnAlert: model.Alert{
				ID: 1,
			},
		},
	}

	for scenario, tc := range tcs {
		t.Run(scenario, func(t *testing.T) {
			// Given
			ctx := context.Background()

			mockReg := new(repository.MockRegistry)
			mockRequestRep := new(request.MockRepository)

			mockRequestRep.ExpectedCalls = []*mock.Call{}
			mockRequestRep.On("GetByID", ctx, tc.givenInput).Return(tc.returnData, nil)

			mockAlertRep := new(alert.MockRepository)
			mockAlertRep.ExpectedCalls = []*mock.Call{}
			mockAlertRep.On("GetAlert", ctx, tc.returnData.AlertID).Return(tc.returnAlert, nil)

			mockReg.ExpectedCalls = []*mock.Call{
				mockReg.On("Request").Return(mockRequestRep),
				mockReg.On("Alert").Return(mockAlertRep),
			}

			// When
			instance := New(mockReg, nil, nil, nil)

			// Then
			require.NotPanics(t, func() { instance.Push(ctx, tc.givenInput) })
		})
	}
}
