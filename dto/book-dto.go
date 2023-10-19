package dto

type BookRequest struct {
	Title  string `json:"title" form:"required"`
	Author string `json:"author" form:"required"`
	Price  int `json:"price" form:"required"`
	Stock  int `json:"stock" form:"required"`
}

type BookResponse struct {
	ID 	uint `json:"id" form:"id"`
	Title  string `json:"title" form:"required"`
	Author string `json:"author" form:"required"`
	Price  int `json:"price" form:"required"`
	Stock  int `json:"stock" form:"required"`
}