package dto

type UserRequest struct {
	Email    string `json:"email" form:"email"`
	Name    string `json:"name" form:"name"`
	Password string `json:"password" form:"password"`
	PhoneNumber string `json:"phoneNumber" form:"phoneNumber"`
	Role string `json:"role" form:"role"`
}

type UserResponse struct {
	ID 		 uint `json:"id" form:"id"`
	Name    string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	PhoneNumber string `json:"phoneNumber" form:"phoneNumber"`
	Role string `json:"role" form:"role"`
}