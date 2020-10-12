// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	models "github.com/nymtech/nym-directory/models"
	mock "github.com/stretchr/testify/mock"
)

// IDb is an autogenerated mock type for the IDb type
type IDb struct {
	mock.Mock
}

// Add provides a mock function with given fields: _a0
func (_m *IDb) Add(_a0 models.PersistedMixStatus) {
	_m.Called(_a0)
}

// BatchAdd provides a mock function with given fields: status
func (_m *IDb) BatchAdd(status []models.PersistedMixStatus) {
	_m.Called(status)
}

// BatchLoadReports provides a mock function with given fields: pubkeys
func (_m *IDb) BatchLoadReports(pubkeys []string) models.BatchMixStatusReport {
	ret := _m.Called(pubkeys)

	var r0 models.BatchMixStatusReport
	if rf, ok := ret.Get(0).(func([]string) models.BatchMixStatusReport); ok {
		r0 = rf(pubkeys)
	} else {
		r0 = ret.Get(0).(models.BatchMixStatusReport)
	}

	return r0
}

// List provides a mock function with given fields: pubkey, limit
func (_m *IDb) List(pubkey string, limit int) []models.PersistedMixStatus {
	ret := _m.Called(pubkey, limit)

	var r0 []models.PersistedMixStatus
	if rf, ok := ret.Get(0).(func(string, int) []models.PersistedMixStatus); ok {
		r0 = rf(pubkey, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.PersistedMixStatus)
		}
	}

	return r0
}

// ListDateRange provides a mock function with given fields: pubkey, ipVersion, start, end
func (_m *IDb) ListDateRange(pubkey string, ipVersion string, start int64, end int64) []models.PersistedMixStatus {
	ret := _m.Called(pubkey, ipVersion, start, end)

	var r0 []models.PersistedMixStatus
	if rf, ok := ret.Get(0).(func(string, string, int64, int64) []models.PersistedMixStatus); ok {
		r0 = rf(pubkey, ipVersion, start, end)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.PersistedMixStatus)
		}
	}

	return r0
}

// LoadNonStaleReports provides a mock function with given fields:
func (_m *IDb) LoadNonStaleReports() models.BatchMixStatusReport {
	ret := _m.Called()

	var r0 models.BatchMixStatusReport
	if rf, ok := ret.Get(0).(func() models.BatchMixStatusReport); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(models.BatchMixStatusReport)
	}

	return r0
}

// LoadReport provides a mock function with given fields: pubkey
func (_m *IDb) LoadReport(pubkey string) models.MixStatusReport {
	ret := _m.Called(pubkey)

	var r0 models.MixStatusReport
	if rf, ok := ret.Get(0).(func(string) models.MixStatusReport); ok {
		r0 = rf(pubkey)
	} else {
		r0 = ret.Get(0).(models.MixStatusReport)
	}

	return r0
}

// SaveBatchMixStatusReport provides a mock function with given fields: _a0
func (_m *IDb) SaveBatchMixStatusReport(_a0 models.BatchMixStatusReport) {
	_m.Called(_a0)
}

// SaveMixStatusReport provides a mock function with given fields: _a0
func (_m *IDb) SaveMixStatusReport(_a0 models.MixStatusReport) {
	_m.Called(_a0)
}
