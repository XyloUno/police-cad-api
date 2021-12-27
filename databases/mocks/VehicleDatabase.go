// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	models "github.com/linesmerrill/police-cad-api/models"
)

// VehicleDatabase is an autogenerated mock type for the VehicleDatabase type
type VehicleDatabase struct {
	mock.Mock
}

// Find provides a mock function with given fields: ctx, filter
func (_m *VehicleDatabase) Find(ctx context.Context, filter interface{}) ([]models.Vehicle, error) {
	ret := _m.Called(ctx, filter)

	var r0 []models.Vehicle
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) []models.Vehicle); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Vehicle)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindOne provides a mock function with given fields: ctx, filter
func (_m *VehicleDatabase) FindOne(ctx context.Context, filter interface{}) (*models.Vehicle, error) {
	ret := _m.Called(ctx, filter)

	var r0 *models.Vehicle
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) *models.Vehicle); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Vehicle)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
