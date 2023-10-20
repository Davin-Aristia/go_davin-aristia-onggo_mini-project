package model

import (
	"gorm.io/gorm"
	"time"
	"github.com/go-playground/validator/v10"
)

type Sales struct {
	*gorm.Model

	UserId uint `json:"userId" validate:"required"`
	Invoice string `json:"invoice" validate:"required" gorm:"unique_index"`
	Date time.Time `json:"date" validate:"required"`
	Total float64 `json:"total" validate:"required"`
}

func ValidateSalesRequest(sales *Sales) error {
    validate := validator.New()
    return validate.Struct(sales)
}