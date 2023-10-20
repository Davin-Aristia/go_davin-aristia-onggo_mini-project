package controller

import (
	"go-mini-project/dto"
	"go-mini-project/usecase"
	"go-mini-project/middleware"

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
			"message": "failed get books",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success get books",
		"book":    book,
	})
}

func (u *bookController) GetBookById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	
	book, err := u.useCase.GetById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed get book",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success get book",
		"book":    book,
	})
}

func (u *bookController) InsertBook(c echo.Context) error {
	role := middleware.ExtractTokenUserRole(c)
	if role != "admin"{
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message": "unauthorized",
		})
	}

	var payloads dto.BookRequest

	if err := c.Bind(&payloads); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind " + err.Error(),
		})
	}

	book, err := u.useCase.Create(payloads)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed create book",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success create book",
		"book":    book,
	})
}

func (u *bookController) UpdateBook(c echo.Context) error {
	role := middleware.ExtractTokenUserRole(c)
	if role != "admin"{
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message": "unauthorized",
		})
	}
	
	var payloads dto.BookRequest

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error parse id",
			"error": err.Error(),
		})
	}

	if err := c.Bind(&payloads); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind " + err.Error(),
		})
	}

	book, err := u.useCase.Update(payloads, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed update book",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success update book",
		"book":    book,
	})
}

func (u *bookController) DeleteBook(c echo.Context) error {
	role := middleware.ExtractTokenUserRole(c)
	if role != "admin"{
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message": "unauthorized",
		})
	}
	
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error parse id",
			"error": err.Error(),
		})
	}

	err = u.useCase.Delete(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed delete book",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success delete book",
	})
}