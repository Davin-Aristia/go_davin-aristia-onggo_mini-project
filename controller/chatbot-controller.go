package controller

import (
	"go-mini-project/dto"
	"go-mini-project/usecase"

	"net/http"

	"github.com/labstack/echo/v4"
)

type ChatbotController interface{}

type chatbotController struct {
	useCase usecase.ChatbotUsecase
}

func NewChatbotController(chatbotUsecase usecase.ChatbotUsecase) *chatbotController {
	return &chatbotController{
		chatbotUsecase,
	}
}

func (u *chatbotController) BookRecommendation(c echo.Context) error {
	var payloads dto.ChatbotRequest
	if err := c.Bind(&payloads); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error bind data",
			"response": err.Error(),
		})
	}

	bookRecommendation, err := u.useCase.GetLaptopRecommendation(payloads)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed",
			"response": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "success",
		"response": bookRecommendation,
	})
}