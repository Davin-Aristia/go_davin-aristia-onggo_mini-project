package controller

import (
	"go-mini-project/dto"
	"go-mini-project/usecase"
	"go-mini-project/middleware"

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
	if role != "admin"{
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message": "unauthorized",
		})
	}
	
	name := c.FormValue("name")
	
	vendor, err := u.useCase.Get(name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed get vendors",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success get vendors",
		"vendor":    vendor,
	})
}

func (u *vendorController) GetVendorById(c echo.Context) error {
	role := middleware.ExtractTokenUserRole(c)
	if role != "admin"{
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message": "unauthorized",
		})
	}

	id, err := strconv.Atoi(c.Param("id"))
	
	vendor, err := u.useCase.GetById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed get vendor",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success get vendor",
		"vendor":    vendor,
	})
}

func (u *vendorController) InsertVendor(c echo.Context) error {
	role := middleware.ExtractTokenUserRole(c)
	if role != "admin"{
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message": "unauthorized",
		})
	}

	var payloads dto.VendorRequest

	if err := c.Bind(&payloads); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind " + err.Error(),
		})
	}

	vendor, err := u.useCase.Create(payloads)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed create vendor",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success create vendor",
		"vendor":    vendor,
	})
}

func (u *vendorController) UpdateVendor(c echo.Context) error {
	role := middleware.ExtractTokenUserRole(c)
	if role != "admin"{
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message": "unauthorized",
		})
	}
	
	var payloads dto.VendorRequest

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

	vendor, err := u.useCase.Update(payloads, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed update vendor",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success update vendor",
		"vendor":    vendor,
	})
}

func (u *vendorController) DeleteVendor(c echo.Context) error {
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
			"message": "failed delete vendor",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success delete vendor",
	})
}