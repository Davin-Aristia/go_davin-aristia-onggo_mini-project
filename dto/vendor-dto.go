package dto

import (
	"go-mini-project/model"
)

type VendorRequest struct {
	Name    string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Address string `json:"address" form:"address"`
	PhoneNumber string `json:"phoneNumber" form:"phoneNumber"`
}

type VendorResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
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