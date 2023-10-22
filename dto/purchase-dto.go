package dto

import "time"

type PurchaseRequest struct {
	VendorId uint `json:"vendor" form:"vendor"`
	PurchaseOrder string `json:"invoice" form:"invoice"`
	Date time.Time `json:"date" form:"date"`
	Total float64 `json:"total" form:"total"`
	Details   []PurchaseDetailRequest `json:"details" form:"details"`
}

type PurchaseDetailRequest struct {
	BookID   uint    `json:"bookId" form:"bookId"`
	Price    float64 `json:"price" form:"price"`
	Quantity int     `json:"quantity" form:"quantity"`
	Subtotal float64 `json:"subtotal" form:"subtotal"`
}