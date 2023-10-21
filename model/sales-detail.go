package model

import (
	"gorm.io/gorm"
	"github.com/go-playground/validator/v10"
)

type SalesDetail struct {
	*gorm.Model

	SalesId uint `json:"salesId" validate:"required"`
	BookId uint `json:"bookId" validate:"required"`
	Price float64 `json:"price" validate:"required"`
	Quantity int `json:"quantity" validate:"required"`
	Subtotal float64 `json:"subtotal" validate:"required"`
}

func ValidateSalesDetailRequest(salesDetail *SalesDetail) error {
    validate := validator.New()
    return validate.Struct(salesDetail)
}