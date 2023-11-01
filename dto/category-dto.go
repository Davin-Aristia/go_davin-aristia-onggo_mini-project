package dto

import (
	"go-mini-project/model"
)

type CategoryRequest struct {
	Name  string `json:"name" form:"name"`
}

type CategoryResponse struct {
	ID    uint       `json:"id"`
	Name  string     `json:"name"`
	Books []categoryBookResponse `json:"book"`
}

type categoryBookResponse struct {
	ID         uint    `json:"id"`
	Title      string  `json:"title"`
	Author     string  `json:"author"`
	Price      float64 `json:"price"`
	Stock      int     `json:"stock"`
}

func convertToCategoryBookResponse(book model.Book) categoryBookResponse {
	return categoryBookResponse{
		ID:         book.ID,
		Title:      book.Title,
		Author:     book.Author,
		Price:      book.Price,
		Stock:      book.Stock,
	}
}

func ConvertToCategoryResponse(category model.Category) CategoryResponse {
	categoryResponse := CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}

	for _, book := range category.Books {
		bookResponse := convertToCategoryBookResponse(book)
		categoryResponse.Books = append(categoryResponse.Books, bookResponse)
	}

	return categoryResponse
}