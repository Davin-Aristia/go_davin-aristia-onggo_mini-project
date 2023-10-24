package controller

import (
	"go-mini-project/dto"
	"go-mini-project/middleware"
	"go-mini-project/usecase"

	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type VendorController interface{}

type vendorController struct {
	useCase usecase.VendorUsecase
}

func NewVendorController(vendorUsecase usecase.VendorUsecase) *vendorController {
	return &vendorController{
		vendorUsecase,
	}
}

func (u *vendorController) GetVendors(c echo.Context) error {
	role := middleware.ExtractTokenUserRole(c)
	if role != "admin" {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Only admin roles are allowed to perform this operation.",
		})
	}

	name := c.FormValue("name")

	vendor, err := u.useCase.Get(name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get vendors",
			"response": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success get vendors",
		"response":  vendor,
	})
}

func (u *vendorController) GetVendorById(c echo.Context) error {
	role := middleware.ExtractTokenUserRole(c)
	if role != "admin" {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Only admin roles are allowed to perform this operation.",
		})
	}

	id, err := strconv.Atoi(c.Param("id"))

	vendor, err := u.useCase.GetById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get vendor",
			"response": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success get vendor",
		"response":  vendor,
	})
}

func (u *vendorController) InsertVendor(c echo.Context) error {
	role := middleware.ExtractTokenUserRole(c)
	if role != "admin" {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Only admin roles are allowed to perform this operation.",
		})
	}

	var payloads dto.VendorRequest

	if err := c.Bind(&payloads); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error bind data",
			"response": err.Error(),
		})
	}

	vendor, err := u.useCase.Create(payloads)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed create vendor",
			"response": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success create vendor",
		"response":  vendor,
	})
}

func (u *vendorController) UpdateVendor(c echo.Context) error {
	role := middleware.ExtractTokenUserRole(c)
	if role != "admin" {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Only admin roles are allowed to perform this operation.",
		})
	}

	var payloads dto.VendorRequest

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

	vendor, err := u.useCase.Update(payloads, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed update vendor",
			"response": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success update vendor",
		"response":  vendor,
	})
}

func (u *vendorController) DeleteVendor(c echo.Context) error {
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
			"message":  "failed delete vendor",
			"response": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success delete vendor",
		"response": "success delete vendor with id "+ strconv.Itoa(id),
	})
}
