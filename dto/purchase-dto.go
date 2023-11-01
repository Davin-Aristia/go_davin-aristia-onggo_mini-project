package dto

import (
	"go-mini-project/model"
)

import "time"

type PurchaseRequest struct {
	VendorId uint `json:"vendor" form:"vendor"`
	PurchaseOrder string `json:"invoice" form:"invoice"`
	Date time.Time `json:"date" form:"date"`
	Total float64 `json:"total" form:"total"`
	Details   []PurchaseDetailRequest `json:"details" form:"details"`
}

type PurchaseDetailRequest struct {
	BookID   uint    `json:"book_id" form:"book_id"`
	Price    float64 `json:"price" form:"price"`
	Quantity int     `json:"quantity" form:"quantity"`
	Subtotal float64 `json:"subtotal" form:"subtotal"`
}

type PurchaseResponse struct {
	ID         		uint     `json:"id"`
	VendorId     	uint     `json:"vendor_id"`
	Vendor	     	string   `json:"vendor"`
	PurchaseOrder   string   `json:"purchase_order"`
	Date       		string   `json:"date"`
	Total      		float64  `json:"total"`
	PurchaseDetails []PurchaseDetailResponse `json:"purchase_detail"`
}

func ConvertToPurchaseResponse(purchase model.Purchase) PurchaseResponse {
	purchaseResponse := PurchaseResponse{
		ID:         	purchase.ID,
		VendorId:     	purchase.VendorId,
		Vendor:     	purchase.Vendor.Name,
		PurchaseOrder:  purchase.PurchaseOrder,
		Date:       	purchase.Date.Format("2006-01-02 15:04:05"),
		Total:      	purchase.Total,
	}

	// Convert PurchaseDetails to PurchaseDetailResponse
	for _, detail := range purchase.PurchaseDetails {
		detailResponse := ConvertToPurchaseDetailResponse(detail)
		purchaseResponse.PurchaseDetails = append(purchaseResponse.PurchaseDetails, detailResponse)
	}

	return purchaseResponse
}


type PurchaseDetailResponse struct {
	ID       uint    `json:"id"`
	PurchaseId  uint `json:"purchase_id"`
	BookId   uint    `json:"book_id"`
	Title  	   string  `json:"title"`
	Author 	   string  `json:"author"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
	Subtotal float64 `json:"subtotal"`
}

func ConvertToPurchaseDetailResponse(detail model.PurchaseDetail) PurchaseDetailResponse {
	return PurchaseDetailResponse{
		ID:       	detail.ID,
		PurchaseId: detail.PurchaseId,
		BookId:   	detail.BookId,
		Title:    	detail.Book.Title,
		Author:   	detail.Book.Author,
		Price:    	detail.Price,
		Quantity: 	detail.Quantity,
		Subtotal: 	detail.Subtotal,
	}
}