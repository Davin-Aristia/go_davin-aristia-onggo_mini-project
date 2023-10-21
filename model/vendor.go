package model

import (
	"gorm.io/gorm"
	"github.com/go-playground/validator/v10"
)

type Vendor struct {
	*gorm.Model

	Name    string `json:"name" validate:"required"`
	Address    string `json:"address" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	PhoneNumber string `json:"phoneNumber" validate:"required"`
}

func ValidateVendorRequest(vendor *Vendor) error {
    validate := validator.New()
    return validate.Struct(vendor)
}
