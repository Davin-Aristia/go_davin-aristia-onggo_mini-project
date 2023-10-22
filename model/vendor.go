package model

import (
	"gorm.io/gorm"
	"github.com/go-playground/validator/v10"
)

type Vendor struct {
	*gorm.Model

	Name    	string 		`json:"name" validate:"required" gorm:"type:varchar(255)"`
	Address    	string 		`json:"address" validate:"required" gorm:"type:varchar(255)"`
	Email    	string 		`json:"email" validate:"required,email" gorm:"type:varchar(255)"`
	PhoneNumber string 		`json:"phoneNumber" validate:"required" gorm:"type:varchar(20)"`
	Purchases 	[]Purchase  `gorm:foreignKey:VendorId`
}

func ValidateVendorRequest(vendor *Vendor) error {
    validate := validator.New()
    return validate.Struct(vendor)
}
