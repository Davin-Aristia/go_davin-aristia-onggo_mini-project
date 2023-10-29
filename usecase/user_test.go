// package usecase

// import (
// 	"testing"

// 	"go-mini-project/dto"
// 	"go-mini-project/model"
// 	"go-mini-project/repository"
// 	"github.com/stretchr/testify/assert"
// )

// func TestCreateUser(t *testing.T) {
// 	data := dto.UserRequest{
// 		Name: "Davin",
// 		Email:      "davin@gmail.com",
// 		Password:   "123456",
// 		PhoneNumber: "1234567890",
// 		Role:       "admin",
// 	}

// 	mockUser := model.User{
// 		Name:      data.Name,
// 		Email:      data.Email,
// 		Password:   data.Password,
// 		PhoneNumber: data.PhoneNumber,
// 		Role:       data.Role,
// 	}

// 	response := dto.UserResponse{
// 		Name:      data.Name,
// 		Email:      data.Email,
// 		PhoneNumber: data.PhoneNumber,
// 		Role:       data.Role,
// 	}

// 	mockUserRepository := repository.NewMockUserRepo()
// 	mockUserRepository.On("Create", mockUser).Return(response)

// 	service := NewUserUsecase(mockUserRepository)

// 	if _, err := service.Create(data); err != nil {
// 		t.Errorf("Got Error %v", err)
// 	}
// }

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


package usecase

import (
	"go-mini-project/model"
	"go-mini-project/dto"
	mocks "go-mini-project/mock"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestCreate(t *testing.T) {
	repo := new(mocks.UserData)
	insertData := dto.UserRequest{
		Name: 		"Davin",
		Email:      "davin@gmail.com",
		Password:   "123456",
		PhoneNumber: "1234567890",
		Role:       "admin",
	}
	returnData := model.User{
		Model:gorm.Model{ID:1},
		Name:      insertData.Name,
		Email:      insertData.Email,
		Password:   insertData.Password,
		PhoneNumber: insertData.PhoneNumber,
		Role:       insertData.Role,
	}
	response := dto.ConvertToUserResponse(returnData)

	t.Run("Success Insert", func(t *testing.T) {
		repo.On("Create", mock.Anything).Return(returnData, nil).Once()
		srv := NewUserUsecase(repo)

		res, err := srv.Create(insertData)
		assert.NoError(t, err)
		assert.Equal(t, response, res)
		repo.AssertExpectations(t)
	})

	t.Run("Error insert to DB", func(t *testing.T) {
		repo.On("Create", mock.Anything).Return(model.User{}, errors.New("there is some error")).Once()
		srv := NewUserUsecase(repo)

		res, err := srv.Create(insertData)
		assert.EqualError(t, err, "there is some error")
		assert.Equal(t, dto.UserResponse{}, res)
		repo.AssertExpectations(t)
	})
}

func TestCheckSignIn(t *testing.T) {
	repo := new(mocks.UserData)
	returnUser := model.User{
		Model:gorm.Model{ID:1},
		Name: 		"Davin",
		Email:      "davin@gmail.com",
		Password:   "$2a$10$0tJQ8HQmVF30N9x2RMLEH.z8kIwT/YVUMRAV2zTUuTOgfLfZZLWcG",
		PhoneNumber: "1234567890",
		Role:       "admin",
	}

	t.Run("Success Login", func(t *testing.T) {
		repo.On("CheckUserExist", "davin@gmail.com", "customer456").Return(dto.UserResponse{
			ID:        returnUser.ID,
			Name:      returnUser.Name,
			Email:     returnUser.Email,
			Role:      returnUser.Role,
		}, nil).Once()

		srv := NewUserUsecase(repo)

		userData, token, err := srv.CheckSignIn("davin@gmail.com", "customer456")

		assert.NoError(t, err)
		assert.Equal(t, "Davin", userData.Name)
		assert.NotEmpty(t, token)
		repo.AssertExpectations(t)
	})

	t.Run("Invalid Login", func(t *testing.T) {
		repo.On("CheckUserExist", "davin@gmail.com", "customer").Return(dto.UserResponse{}, errors.New("Invalid Email or Password")).Once()

		srv := NewUserUsecase(repo)

		userData, token, err := srv.CheckSignIn("davin@gmail.com", "customer")

		assert.Error(t, err)
		assert.EqualError(t, err, "Invalid Email or Password")
		assert.Equal(t, dto.UserResponse{}, userData)
		assert.Empty(t, token)
		repo.AssertExpectations(t)
	})
}