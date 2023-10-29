package dto

import (
	"go-mini-project/model"
)

type UserRequest struct {
	Email    string `json:"email" form:"email"`
	Name    string `json:"name" form:"name"`
	Password string `json:"password" form:"password"`
	PhoneNumber string `json:"phoneNumber" form:"phoneNumber"`
	Role string `json:"role" form:"role"`
}

type UserResponse struct {
	ID 		 uint `json:"id"`
	Name    string `json:"name"`
	Email    string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Role string `json:"role"`
}

func ConvertToUserResponse(user model.User) UserResponse {
	return UserResponse{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Role:        user.Role,
	}
}