package model

import (
	"gorm.io/gorm"
	"github.com/go-playground/validator/v10"
)

type Book struct {
	gorm.Model

	Title  		string 	`json:"title" validate:"required" gorm:"type:varchar(255)"`
	Author 		string 	`json:"author" validate:"required" gorm:"type:varchar(255)"`
	Price  		float64 `json:"price" validate:"required" gorm:"type:decimal(15,2)"`
	Stock  		int 	`json:"stock"`
	CategoryId 	uint 	`json:"categoryId" validate:"required"`
	Category 	Category `gorm:"foreignKey:CategoryId" validate:"-"`
}

func ValidateBookRequest(book *Book) error {
    validate := validator.New()
    return validate.Struct(book)
}