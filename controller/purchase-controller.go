package controller

import (
	"go-mini-project/dto"
	"go-mini-project/usecase"
	"go-mini-project/middleware"

	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PurchaseController interface{}

type purchaseController struct {
	useCase usecase.PurchaseUsecase
}

func NewPurchaseController(purchaseUsecase usecase.PurchaseUsecase) *purchaseController {
	return &purchaseController{
		purchaseUsecase,
	}
}

func (u *purchaseController) GetPurchase(c echo.Context) error {
	role := middleware.ExtractTokenUserRole(c)
	if role != "admin"{
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message": "unauthorized",
		})
	}

	purchaseOrder := c.FormValue("purchaseOrder")
	vendorId := c.FormValue("vendor")
	
	purchase, err := u.useCase.Get(purchaseOrder, vendorId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed get purchase",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success get purchase",
		"purchase":    purchase,
	})
}

func (u *purchaseController) GetPurchaseById(c echo.Context) error {
	role := middleware.ExtractTokenUserRole(c)
	if role != "admin"{
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message": "unauthorized",
		})
	}

	id, err := strconv.Atoi(c.Param("id"))
	
	purchase, err := u.useCase.GetById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed get purchase",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success get purchase",
		"purchase":    purchase,
	})
}

func (u *purchaseController) CreatePurchase(c echo.Context) error {
	role := middleware.ExtractTokenUserRole(c)
	if role != "admin"{
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message": "unauthorized",
		})
	}

	var payloads dto.PurchaseRequest

	if err := c.Bind(&payloads); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind " + err.Error(),
		})
	}

	purchase, err := u.useCase.Create(payloads)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed create purchase",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success create purchase",
		"purchase":    purchase,
	})
}