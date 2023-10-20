package dto

import "time"

type SalesRequest struct {
	UserId uint `json:"user" form:"user"`
	Invoice string `json:"invoice" form:"invoice"`
	Date time.Time `json:"date" form:"date"`
	Total float64 `json:"total" form:"total"`
	Details   []SalesDetailRequest `json:"details" form:"details"`
}

type SalesDetailRequest struct {
	BookID   uint    `json:"bookId" form:"bookId"`
	Price    float64 `json:"price" form:"price"`
	Quantity int     `json:"quantity" form:"quantity"`
	Subtotal float64 `json:"subtotal" form:"subtotal"`
}