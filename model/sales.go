package model

import (
	"gorm.io/gorm"
	"time"
	"github.com/go-playground/validator/v10"
)

type Sales struct {
	gorm.Model

	UserId 		 uint 			`json:"userId" validate:"required"`
	Invoice 	 string 		`json:"invoice" validate:"required" gorm:"type:char(19);unique_index"`
	Date 		 time.Time 		`json:"date" validate:"required" gorm:"type:datetime"`
	Total 		 float64 		`json:"total" gorm:"type:decimal(15,2)"`
	SalesDetails []SalesDetail  `gorm:foreignKey:SalesId`
}

func ValidateSalesRequest(sales *Sales) error {
    validate := validator.New()
    return validate.Struct(sales)
}

type SalesDetail struct {
	gorm.Model

	SalesId   uint `json:"salesId" validate:"required"`
	BookId 	  uint `json:"bookId" validate:"required"`
	Price 	  float64 `json:"price" validate:"required" gorm:"type:decimal(15,2)"`
	Quantity  int `json:"quantity" validate:"required"`
	Subtotal  float64 `json:"subtotal" validate:"required" gorm:"type:decimal(15,2)"`
	Book 	  Book `gorm:"foreignKey:BookId" validate:"-"`
}

func ValidateSalesDetailRequest(salesDetail *SalesDetail) error {
    validate := validator.New()
    return validate.Struct(salesDetail)
}