package mock

import (
	"go-mini-project/model"

	mock "github.com/stretchr/testify/mock"
)

type CategoryData struct {
	mock.Mock
}

func (_m *CategoryData) Get(name string) ([]model.Category, error) {
	ret := _m.Called(name)

	var r0 []model.Category
	if rf, ok := ret.Get(0).(func(string) []model.Category); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).([]model.Category)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *CategoryData) GetById(id int) (model.Category, error) {
	ret := _m.Called(id)

	var r0 model.Category
	if rf, ok := ret.Get(0).(func(int) model.Category); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(model.Category)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *CategoryData) Create(data model.Category) (model.Category, error) {
	ret := _m.Called(data)

	var r0 model.Category
	if rf, ok := ret.Get(0).(func(model.Category) model.Category); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(model.Category)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.Category) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *CategoryData) Update(data model.Category, ID int) error {
	ret := _m.Called(data, ID)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.Category, int) error); ok {
		r0 = rf(data, ID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *CategoryData) Delete(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}