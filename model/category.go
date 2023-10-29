package model

import (
	"gorm.io/gorm"
	"github.com/go-playground/validator/v10"
)

type Category struct {
	gorm.Model

	Name  string `json:"name" validate:"required" gorm:"type:varchar(255)"`
	Books []Book `gorm:foreignKey:CategoryId`
}

func ValidateCategoryRequest(category *Category) error {
    validate := validator.New()
    return validate.Struct(category)
}