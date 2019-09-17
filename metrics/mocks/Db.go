// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import models "github.com/nymtech/directory-server/models"

// Db is an autogenerated mock type for the Db type
type Db struct {
	mock.Mock
}

// Add provides a mock function with given fields: _a0
func (_m *Db) Add(_a0 models.MixMetric) {
	_m.Called(_a0)
}

// List provides a mock function with given fields:
func (_m *Db) List() []models.MixMetric {
	ret := _m.Called()

	var r0 []models.MixMetric
	if rf, ok := ret.Get(0).(func() []models.MixMetric); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.MixMetric)
		}
	}

	return r0
}
