package model

import (
	"gorm.io/gorm"
	"github.com/go-playground/validator/v10"
)

type Book struct {
	*gorm.Model

	Title  string `json:"title" validate:"required"`
	Author string `json:"author" validate:"required"`
	Price  float64 `json:"price" validate:"required"`
	Stock  int `json:"stock"`
	CategoryId uint `json:"categoryId" validate:"required"`
}

func ValidateBookRequest(book *Book) error {
    validate := validator.New()
    return validate.Struct(book)
}