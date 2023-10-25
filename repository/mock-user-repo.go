package repository

import (
	"go-mini-project/model"
	"go-mini-project/dto"
	mock "github.com/stretchr/testify/mock"

	"fmt"
)

type mockUserReposory struct {
	mock.Mock
}

func NewMockUserRepo() *mockUserReposory {
	return &mockUserReposory{}
}

func (m *mockUserReposory) CheckUserExist(email string, password string) (dto.UserResponse, error) {
	ret := m.Called(email, password)
	return ret.Get(0).(dto.UserResponse), ret.Error(1)
}

func (m *mockUserReposory) Get() ([]model.User, error) {
	ret := m.Called()
	return ret.Get(0).([]model.User), ret.Error(1)
}

func (m *mockUserReposory) Create(data model.User) (model.User, error) {
	fmt.Println("data", data)
	ret := m.Called(data)
	fmt.Println("ret", ret)
	
	// Check if the function was called with a model.User argument
	userArg, ok := ret.Get(0).(model.User)
	if !ok {
		return model.User{}, fmt.Errorf("unexpected type for user argument")
	}
	fmt.Println("userArg", userArg)
	
	// Check if the function was called with an error argument
	if errArg := ret.Error(1); errArg != nil {
		return model.User{}, errArg
	}

	return userArg, nil
}