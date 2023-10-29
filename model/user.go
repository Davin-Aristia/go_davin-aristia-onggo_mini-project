package model

import (
	"gorm.io/gorm"
	"github.com/go-playground/validator/v10"
)

type User struct {
	gorm.Model

	Name    	string  `json:"name" validate:"required" gorm:"type:varchar(255)"`
	Email    	string  `json:"email" validate:"required,email" gorm:"type:varchar(255)"`
	Password 	string  `json:"password" validate:"required" gorm:"type:varchar(100)"`
	PhoneNumber string  `json:"phoneNumber" validate:"required" gorm:"type:varchar(20)"`
	Role 		string  `json:"role" gorm:"type:enum('customer','admin');default:'customer'"`
	Sales 		[]Sales `gorm:foreignKey:UserId`
}

func ValidateUserRequest(user *User) error {
    validate := validator.New()
    return validate.Struct(user)
}
