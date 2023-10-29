package controller

import (
	"go-mini-project/dto"
	"go-mini-project/middleware"
	"go-mini-project/usecase"

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
			"message":  "failed get categories",
			"response": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success get categories",
		"response": category,
	})
}

func (u *categoryController) GetCategoryById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	category, err := u.useCase.GetById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get category",
			"response": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success get category",
		"response": category,
	})
}

func (u *categoryController) InsertCategory(c echo.Context) error {
	role := middleware.ExtractTokenUserRole(c)
	if role != "admin" {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Only admin roles are allowed to perform this operation.",
		})
	}

	var payloads dto.CategoryRequest

	if err := c.Bind(&payloads); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error bind data",
			"response": err.Error(),
		})
	}

	category, err := u.useCase.Create(payloads)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed create category",
			"response": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success create category",
		"response": category,
	})
}

func (u *categoryController) UpdateCategory(c echo.Context) error {
	role := middleware.ExtractTokenUserRole(c)
	if role != "admin" {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Only admin roles are allowed to perform this operation.",
		})
	}

	var payloads dto.CategoryRequest

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error parse id",
			"response": err.Error(),
		})
	}

	if err := c.Bind(&payloads); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error bind data",
			"response": err.Error(),
		})
	}

	category, err := u.useCase.Update(payloads, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed update category",
			"response": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success update category",
		"response": category,
	})
}

func (u *categoryController) DeleteCategory(c echo.Context) error {
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
			"message":  "error parse id",
			"response": err.Error(),
		})
	}

	err = u.useCase.Delete(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed delete category",
			"response": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success delete category",
		"response": "success delete category with id "+ strconv.Itoa(id),
	})
}
