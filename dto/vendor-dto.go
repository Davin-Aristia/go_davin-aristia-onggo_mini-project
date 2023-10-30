package dto

import (
	"go-mini-project/model"
)

type VendorRequest struct {
	Name    string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Address string `json:"address" form:"address"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
}

type VendorResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

func ConvertToVendorResponse(vendor model.Vendor) VendorResponse {
	return VendorResponse{
		ID:          vendor.ID,
		Name:        vendor.Name,
		Address:     vendor.Address,
		Email:       vendor.Email,
		PhoneNumber: vendor.PhoneNumber,
	}
}