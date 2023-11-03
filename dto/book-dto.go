package dto

import (
	"go-mini-project/model"
)

type BookRequest struct {
	Title  string `json:"title" form:"title"`
	Author string `json:"author" form:"author"`
	Price  float64 `json:"price" form:"price"`
	Stock  int `json:"stock" form:"stock"`
	CategoryId  int `json:"category" form:"category"`
}

type BookResponse struct {
	ID         uint    `json:"id"`
	Title      string  `json:"title"`
	Author     string  `json:"author"`
	Price      float64 `json:"price"`
	Stock      int     `json:"stock"`
	CategoryId int     `json:"category_id"`
	CategoryName   string  `json:"category_name"`
}

func ConvertToBookResponse(book model.Book) BookResponse {
	return BookResponse{
		ID:         book.ID,
		Title:      book.Title,
		Author:     book.Author,
		Price:      book.Price,
		Stock:      book.Stock,
		CategoryId: int(book.CategoryId),
		CategoryName: 	book.Category.Name,
	}
}