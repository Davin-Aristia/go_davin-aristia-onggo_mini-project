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
	Books []BookResponse `json:"book"`
}

func ConvertToCategoryResponse(category model.Category) CategoryResponse {
	categoryResponse := CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}

	for _, book := range category.Books {
		bookResponse := ConvertToBookResponse(book)
		categoryResponse.Books = append(categoryResponse.Books, bookResponse)
	}

	return categoryResponse
}