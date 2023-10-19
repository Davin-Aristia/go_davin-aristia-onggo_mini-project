package dto

type UserRequest struct {
	Email    string `json:"email" form:"email"`
	Name    string `json:"name" form:"name"`
	Password string `json:"password" form:"password"`
	PhoneNumber string `json:"phoneNumber" form:"phoneNumber"`
	Role string `json:"role" form:"role"`
}