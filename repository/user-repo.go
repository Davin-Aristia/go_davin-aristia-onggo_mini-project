package repository

import (
	"errors"
	
	"go-mini-project/dto"
	"go-mini-project/model"

	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	CheckUserExist(email string, password string) (dto.UserResponse, error)
	Create(data model.User) (model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) CheckUserExist(email string, password string) (dto.UserResponse, error) {
	var data model.User
	tx := r.db.Where("email = ?", email).First(&data)
	if tx.Error != nil {
		return dto.UserResponse{}, errors.New("Invalid Email or Password")
	}

	err := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(password))
	if err != nil {
		return dto.UserResponse{}, tx.Error
	}

	userData := dto.UserResponse{
		ID:    data.ID,
		Name:  data.Name,
		Email: data.Email,
		Role:  data.Role,
	}
	return userData, nil
}

func (r *userRepository) Create(data model.User) (model.User, error) {
	err := model.ValidateUserRequest(&data)
	if err != nil {
		return model.User{}, err
	}
	
	hashPassword,err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return model.User{}, err
	}
	data.Password = string(hashPassword)

	tx := r.db.Save(&data)
	if tx.Error != nil {
		return model.User{}, tx.Error
	}
	return data, nil
}