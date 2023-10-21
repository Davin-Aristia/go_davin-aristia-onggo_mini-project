package dto

type VendorRequest struct {
	Name    string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Address string `json:"address" form:"address"`
	PhoneNumber string `json:"phoneNumber" form:"phoneNumber"`
}