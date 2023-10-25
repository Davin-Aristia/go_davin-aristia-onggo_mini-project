package usecase

import (
	"testing"
	"fmt"

	"go-mini-project/dto"
	"go-mini-project/model"
	"go-mini-project/repository"
	// "github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	data := dto.UserRequest{
		Name: "Davin",
		Email:      "davin@gmail.com",
		Password:   "123456",
		PhoneNumber: "1234567890",
		Role:       "admin",
	}

	mockUser := model.User{
		Name:      data.Name,
		Email:      data.Email,
		Password:   data.Password,
		PhoneNumber: data.PhoneNumber,
		Role:       data.Role,
	}

	fmt.Println("mockUser", mockUser)

	mockUserRepository := repository.NewMockUserRepo()
	mockUserRepository.On("Create", mockUser).Return(model.User{}, nil)

	service := NewUserUsecase(mockUserRepository)

	if _, err := service.Create(data); err != nil {
		t.Errorf("Got Error %v", err)
	}
}

// func TestCheckLogin(t *testing.T) {
// 	mockUserRepository := repository.NewMockUserRepo()

// 	expectedUser := model.User{
// 		Email:    "davin@gmail.com",
// 		Password: "123",
// 		PhoneNumber: "1234567890",
// 		Role:       "user",
// 	}

// 	expectedDTO := dto.UserResponse{
// 		ID:          1,
// 		Email:       expectedUser.Email,
// 		PhoneNumber: expectedUser.PhoneNumber,
// 		Role:        expectedUser.Role,
// 	}

// 	mockUserRepository.On("CheckUserExist", expectedUser.Email, expectedUser.Password).Return(expectedDTO, nil)

// 	userUsecase := NewUserUsecase(mockUserRepository)
// 	userData, token, err := userUsecase.CheckSignIn(expectedUser.Email, expectedUser.Password)

// 	assert.Equal(t, expectedDTO, userData)
// 	assert.NotEmpty(t, token)
// 	assert.NoError(t, err)
// }