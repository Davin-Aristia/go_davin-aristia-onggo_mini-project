package controller

import (
	"go-mini-project/dto"
	"go-mini-project/middleware"
	"go-mini-project/usecase"

	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BookController interface{}

type bookController struct {
	useCase usecase.BookUsecase
}

func NewBookController(bookUsecase usecase.BookUsecase) *bookController {
	return &bookController{
		bookUsecase,
	}
}

func (u *bookController) GetBooks(c echo.Context) error {
	title := c.FormValue("title")
	author := c.FormValue("author")
	category := c.FormValue("category")

	book, err := u.useCase.Get(title, author, category)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get books",
			"response": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success get books",
		"response": book,
	})
}

func (u *bookController) GetBookById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	book, err := u.useCase.GetById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get book",
			"response": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success get book",
		"response": book,
	})
}

func (u *bookController) InsertBook(c echo.Context) error {
	role := middleware.ExtractTokenUserRole(c)
	if role != "admin" {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Only admin roles are allowed to perform this operation.",
		})
	}

	var payloads dto.BookRequest

	if err := c.Bind(&payloads); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error bind data",
			"response": err.Error(),
		})
	}

	book, err := u.useCase.Create(payloads)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed create book",
			"response": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success create book",
		"response": book,
	})
}

func (u *bookController) UpdateBook(c echo.Context) error {
	role := middleware.ExtractTokenUserRole(c)
	if role != "admin" {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Only admin roles are allowed to perform this operation.",
		})
	}

	var payloads dto.BookRequest

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error parse id",
			"response":   err.Error(),
		})
	}

	if err := c.Bind(&payloads); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error bind data",
			"response": err.Error(),
		})
	}

	book, err := u.useCase.Update(payloads, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed update book",
			"response":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success update book",
		"response":    book,
	})
}

func (u *bookController) DeleteBook(c echo.Context) error {
	role := middleware.ExtractTokenUserRole(c)
	if role != "admin" {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Only admin roles are allowed to perform this operation.",
		})
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error parse id",
			"response":   err.Error(),
		})
	}

	err = u.useCase.Delete(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed delete book",
			"response":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success delete book",
		"response": "success delete book with id "+strconv.Itoa(id),
	})
}
