package mock

import (
	"go-mini-project/model"

	mock "github.com/stretchr/testify/mock"
)

type BookData struct {
	mock.Mock
}

func (_m *BookData) Get(title, author string, category int) ([]model.Book, error) {
	ret := _m.Called(title, author, category)

	var r0 []model.Book
	if rf, ok := ret.Get(0).(func(string, string, int) []model.Book); ok {
		r0 = rf(title, author, category)
	} else {
		r0 = ret.Get(0).([]model.Book)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, int) error); ok {
		r1 = rf(title, author, category)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *BookData) GetById(id int) (model.Book, error) {
	ret := _m.Called(id)

	var r0 model.Book
	if rf, ok := ret.Get(0).(func(int) model.Book); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(model.Book)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *BookData) Create(data model.Book) (model.Book, error) {
	ret := _m.Called(data)

	var r0 model.Book
	if rf, ok := ret.Get(0).(func(model.Book) model.Book); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(model.Book)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.Book) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *BookData) Update(data model.Book, ID int) error {
	ret := _m.Called(data, ID)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.Book, int) error); ok {
		r0 = rf(data, ID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *BookData) Delete(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}