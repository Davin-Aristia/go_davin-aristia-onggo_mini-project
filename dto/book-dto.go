package dto

type BookRequest struct {
	Title  string `json:"title" form:"title"`
	Author string `json:"author" form:"author"`
	Price  float64 `json:"price" form:"price"`
	Stock  int `json:"stock" form:"stock"`
	CategoryId  int `json:"category" form:"category"`
}