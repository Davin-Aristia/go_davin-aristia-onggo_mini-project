package model

import (
	"gorm.io/gorm"
	"github.com/go-playground/validator/v10"
)

type User struct {
	*gorm.Model

	Name    string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	PhoneNumber string `json:"phoneNumber" validate:"required"`
	Role string `json:"role" gorm:"type:enum('customer','admin');default:'customer'"`
	Sales []Sales `gorm:foreignKey:UserId`
}

func ValidateUserRequest(user *User) error {
    validate := validator.New()
    return validate.Struct(user)
}
