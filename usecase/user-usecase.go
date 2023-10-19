package usecase

import (
	"errors"

	"go-mini-project/dto"
	"go-mini-project/middleware"
	"go-mini-project/model"
	"go-mini-project/repository"
)

type UserUsecase interface {
	CheckSignIn(email string, password string) (dto.UserResponse, string, error)
	Create(payloads dto.UserRequest) (model.User, error)
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) *userUsecase {
	return &userUsecase{userRepository: userRepo}
}

func (s *userUsecase) CheckSignIn(email string, password string) (dto.UserResponse, string, error) {
	userData, err := s.userRepository.CheckUserExist(email, password)
	if err != nil {
		return dto.UserResponse{}, "", err
	}

	var token string
	var customError error
	if userData.ID != 0 {
		var errToken error
		token, errToken = middleware.CreateToken(int(userData.ID), userData.Name, userData.Email, userData.Role)
		if errToken != nil {
			return dto.UserResponse{}, "", errToken
		}
	} else {
		customError = errors.New("Invalid Email or Password")
	}
	return userData, token, customError
}

func (s *userUsecase) Create(payloads dto.UserRequest) (model.User, error) {
	userData := model.User{
		Name : payloads.Name,
		Email : payloads.Email,
		Password : payloads.Password,
		PhoneNumber : payloads.PhoneNumber,
		Role : payloads.Role,
	}

	userData, err := s.userRepository.Create(userData)
	if err != nil {
		return model.User{}, err
	}
	return userData, nil
}
