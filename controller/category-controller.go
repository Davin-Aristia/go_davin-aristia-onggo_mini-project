package controller

import (
	"go-mini-project/dto"
	"go-mini-project/usecase"
	"go-mini-project/middleware"

	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CategoryController interface{}

type categoryController struct {
	useCase usecase.CategoryUsecase
}

func NewCategoryController(categoryUsecase usecase.CategoryUsecase) *categoryController {
	return &categoryController{
		categoryUsecase,
	}
}

func (u *categoryController) GetCategories(c echo.Context) error {
	name := c.FormValue("name")
	
	category, err := u.useCase.Get(name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed get categories",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success get categories",
		"category":    category,
	})
}

func (u *categoryController) GetCategoryById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	
	category, err := u.useCase.GetById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed get category",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success get category",
		"category":    category,
	})
}

func (u *categoryController) InsertCategory(c echo.Context) error {
	role := middleware.ExtractTokenUserRole(c)
	if role != "admin"{
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message": "unauthorized",
		})
	}

	var payloads dto.CategoryRequest

	if err := c.Bind(&payloads); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind " + err.Error(),
		})
	}

	category, err := u.useCase.Create(payloads)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed create category",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success create category",
		"category":    category,
	})
}

func (u *categoryController) UpdateCategory(c echo.Context) error {
	role := middleware.ExtractTokenUserRole(c)
	if role != "admin"{
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message": "unauthorized",
		})
	}
	
	var payloads dto.CategoryRequest

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

	category, err := u.useCase.Update(payloads, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed update category",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success update category",
		"category":    category,
	})
}

func (u *categoryController) DeleteCategory(c echo.Context) error {
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
			"message": "failed delete category",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success delete category",
	})
}