package dto

import (
	"go-mini-project/model"
	"time"
)


type SalesRequest struct {
	UserId uint `json:"user" form:"user"`
	Invoice string `json:"invoice" form:"invoice"`
	Date time.Time `json:"date" form:"date"`
	Total float64 `json:"total" form:"total"`
	Details   []SalesDetailRequest `json:"details" form:"details"`
}

type SalesDetailRequest struct {
	BookId   uint    `json:"book_id" form:"book_id"`
	Price    float64 `json:"price" form:"price"`
	Quantity int     `json:"quantity" form:"quantity"`
	Subtotal float64 `json:"subtotal" form:"subtotal"`
}

type SalesResponse struct {
	ID         uint              `json:"id"`
	UserId     uint              `json:"user_id"`
	Invoice    string            `json:"invoice"`
	Date       string            `json:"date"`
	Total      float64           `json:"total"`
	SalesDetails []SalesDetailResponse `json:"sales_detail"`
}

func ConvertToSalesResponse(sales model.Sales) SalesResponse {
	salesResponse := SalesResponse{
		ID:         sales.ID,
		UserId:     sales.UserId,
		Invoice:    sales.Invoice,
		Date:       sales.Date.Format("2006-01-02 15:04:05"),
		Total:      sales.Total,
	}

	// Convert SalesDetails to SalesDetailResponse
	for _, detail := range sales.SalesDetails {
		detailResponse := ConvertToSalesDetailResponse(detail)
		salesResponse.SalesDetails = append(salesResponse.SalesDetails, detailResponse)
	}

	return salesResponse
}


type SalesDetailResponse struct {
	ID       uint    `json:"id"`
	SalesId  uint    `json:"sales_id"`
	BookId   uint    `json:"book_id"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
	Subtotal float64 `json:"subtotal"`
}

func ConvertToSalesDetailResponse(detail model.SalesDetail) SalesDetailResponse {
	return SalesDetailResponse{
		ID:       detail.ID,
		SalesId:  detail.SalesId,
		BookId:   detail.BookId,
		Price:    detail.Price,
		Quantity: detail.Quantity,
		Subtotal: detail.Subtotal,
	}
}