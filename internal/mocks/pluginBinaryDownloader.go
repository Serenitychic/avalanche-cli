// Code generated by mockery v2.18.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// PluginBinaryDownloader is an autogenerated mock type for the PluginBinaryDownloader type
type PluginBinaryDownloader struct {
	mock.Mock
}

// InstallVM provides a mock function with given fields: vmID, vmBin
func (_m *PluginBinaryDownloader) InstallVM(vmID string, vmBin string) error {
	ret := _m.Called(vmID, vmBin)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(vmID, vmBin)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpgradeVM provides a mock function with given fields: vmID, vmBin
func (_m *PluginBinaryDownloader) UpgradeVM(vmID string, vmBin string) error {
	ret := _m.Called(vmID, vmBin)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(vmID, vmBin)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewPluginBinaryDownloader interface {
	mock.TestingT
	Cleanup(func())
}

// NewPluginBinaryDownloader creates a new instance of PluginBinaryDownloader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPluginBinaryDownloader(t mockConstructorTestingTNewPluginBinaryDownloader) *PluginBinaryDownloader {
	mock := &PluginBinaryDownloader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
