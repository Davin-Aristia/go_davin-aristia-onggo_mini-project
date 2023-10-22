package model

import (
	"gorm.io/gorm"
	"time"
	"github.com/go-playground/validator/v10"
)

type Purchase struct {
	*gorm.Model

	VendorId 		uint 			 `json:"vendorId" validate:"required"`
	PurchaseOrder 	string 			 `json:"purchaseOrder" validate:"required" gorm:"type:char(19);unique_index"`
	Date 			time.Time 		 `json:"date" validate:"required" gorm:"type:datetime"`
	Total 			float64 		 `json:"total" gorm:"type:decimal(15,2)"`
	PurchaseDetails []PurchaseDetail `gorm:foreignKey:PurchaseId`
}

func ValidatePurchaseRequest(purchase *Purchase) error {
    validate := validator.New()
    return validate.Struct(purchase)
}


type PurchaseDetail struct {
	*gorm.Model

	PurchaseId 	uint `json:"purchaseId" validate:"required"`
	BookId 		uint `json:"bookId" validate:"required"`
	Price 		float64 `json:"price" validate:"required" gorm:"type:decimal(15,2)"`
	Quantity 	int `json:"quantity" validate:"required"`
	Subtotal 	float64 `json:"subtotal" validate:"required" gorm:"type:decimal(15,2)"`
}

func ValidatePurchaseDetailRequest(purchaseDetail *PurchaseDetail) error {
    validate := validator.New()
    return validate.Struct(purchaseDetail)
}