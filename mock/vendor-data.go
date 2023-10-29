package mock

import (
	"go-mini-project/model"

	mock "github.com/stretchr/testify/mock"
)

type VendorData struct {
	mock.Mock
}

func (_m *VendorData) Get(name string) ([]model.Vendor, error) {
	ret := _m.Called(name)

	var r0 []model.Vendor
	if rf, ok := ret.Get(0).(func(string) []model.Vendor); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).([]model.Vendor)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *VendorData) GetById(id int) (model.Vendor, error) {
	ret := _m.Called(id)

	var r0 model.Vendor
	if rf, ok := ret.Get(0).(func(int) model.Vendor); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(model.Vendor)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *VendorData) Create(data model.Vendor) (model.Vendor, error) {
	ret := _m.Called(data)

	var r0 model.Vendor
	if rf, ok := ret.Get(0).(func(model.Vendor) model.Vendor); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(model.Vendor)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.Vendor) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *VendorData) Update(data model.Vendor, ID int) error {
	ret := _m.Called(data, ID)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.Vendor, int) error); ok {
		r0 = rf(data, ID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *VendorData) Delete(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}