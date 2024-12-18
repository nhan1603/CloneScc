// Code generated by mockery v2.20.0. DO NOT EDIT.

package asset

import (
	context "context"

	model "github.com/nhan1603/CloneScc/api/internal/model"
	mock "github.com/stretchr/testify/mock"
)

// MockController is an autogenerated mock type for the Controller type
type MockController struct {
	mock.Mock
}

// GetDevices provides a mock function with given fields: ctx, input
func (_m *MockController) GetDevices(ctx context.Context, input GetDevicesInput) ([]model.Devices, int64, error) {
	ret := _m.Called(ctx, input)

	var r0 []model.Devices
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, GetDevicesInput) ([]model.Devices, int64, error)); ok {
		return rf(ctx, input)
	}
	if rf, ok := ret.Get(0).(func(context.Context, GetDevicesInput) []model.Devices); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Devices)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, GetDevicesInput) int64); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(context.Context, GetDevicesInput) error); ok {
		r2 = rf(ctx, input)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetPremises provides a mock function with given fields: ctx, input
func (_m *MockController) GetPremises(ctx context.Context, input GetPremisesInput) ([]model.Premises, error) {
	ret := _m.Called(ctx, input)

	var r0 []model.Premises
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, GetPremisesInput) ([]model.Premises, error)); ok {
		return rf(ctx, input)
	}
	if rf, ok := ret.Get(0).(func(context.Context, GetPremisesInput) []model.Premises); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Premises)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, GetPremisesInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateDeviceToken provides a mock function with given fields: ctx, input
func (_m *MockController) UpdateDeviceToken(ctx context.Context, input UpdateDeviceTokenInput) error {
	ret := _m.Called(ctx, input)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, UpdateDeviceTokenInput) error); ok {
		r0 = rf(ctx, input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewMockController interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockController creates a new instance of MockController. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockController(t mockConstructorTestingTNewMockController) *MockController {
	mock := &MockController{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
