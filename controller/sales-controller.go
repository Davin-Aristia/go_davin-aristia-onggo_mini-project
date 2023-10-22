package controller

import (
	"go-mini-project/dto"
	"go-mini-project/usecase"
	"go-mini-project/middleware"

	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type SalesController interface{}

type salesController struct {
	useCase usecase.SalesUsecase
}

func NewSalesController(salesUsecase usecase.SalesUsecase) *salesController {
	return &salesController{
		salesUsecase,
	}
}

func (u *salesController) GetSales(c echo.Context) error {
	role := middleware.ExtractTokenUserRole(c)
	user := middleware.ExtractTokenUserId(c)
	if user == 0{
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message": "unauthorized",
		})
	}

	invoice := c.FormValue("invoice")
	var userId string

	if role == "admin"{
		userId = c.FormValue("user")
	}else{
		userId = strconv.Itoa(user)
	}
	
	sales, err := u.useCase.Get(invoice, userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed get sales",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success get sales",
		"sales":    sales,
	})
}

func (u *salesController) GetSalesById(c echo.Context) error {
	role := middleware.ExtractTokenUserRole(c)
	user := middleware.ExtractTokenUserId(c)

	id, err := strconv.Atoi(c.Param("id"))
	
	sales, err := u.useCase.GetById(id, user, role)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed get sales",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success get sales",
		"sales":    sales,
	})
}

func (u *salesController) Checkout(c echo.Context) error {
	user := middleware.ExtractTokenUserId(c)
	if user == 0{
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message": "unauthorized",
		})
	}

	var payloads dto.SalesRequest

	if err := c.Bind(&payloads); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind " + err.Error(),
		})
	}

	sales, err := u.useCase.Create(payloads, user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed create sales",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success create sales",
		"sales":    sales,
	})
}