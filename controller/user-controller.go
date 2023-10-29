package controller

import (
	"go-mini-project/config"
	"go-mini-project/dto"
	"go-mini-project/template"
	"go-mini-project/usecase"

	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type UserController interface{}

type userController struct {
	useCase usecase.UserUsecase
}

func NewUserController(userUsecase usecase.UserUsecase) *userController {
	return &userController{
		userUsecase,
	}
}

func (u *userController) SignUp(c echo.Context) error {
	var payloads dto.UserRequest

	if err := c.Bind(&payloads); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "error bind data",
			"response": err.Error(),
		})
	}

	user, err := u.useCase.Create(payloads)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed sign up",
			"response": err.Error(),
		})
	}

	//email
	currentDate := time.Now().Format("2006-01-02 15:04:05")

	emailBody, err := template.RenderSignupTemplate(currentDate, payloads.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message":  "failed render mail template",
			"response": err.Error(),
		})
	}

	err = config.SendMail(payloads.Email, "Sign up activity to Book Store API", emailBody)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message":  "failed send email",
			"response": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success sign up",
		"response":    user,
	})
}

func (u *userController) SignIn(c echo.Context) error {
	var signInReq dto.UserRequest
	errBind := c.Bind(&signInReq)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data",
			"response": errBind.Error(),
		})
	}

	data, token, err := u.useCase.CheckSignIn(signInReq.Email, signInReq.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message":  "fail signIn",
			"response": err.Error(),
		})
	}
	response := map[string]any{
		"user_id": data.ID,
		"name":    data.Name,
		"email":   data.Email,
		"role":    data.Role,
		"token":   token,
	}

	//email
	currentDate := time.Now().Format("2006-01-02 15:04:05")

	emailBody, err := template.RenderSigninTemplate(currentDate, data.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message":  "failed render mail template",
			"response": err.Error(),
		})
	}

	err = config.SendMail(data.Email, "Sign in activity to Book Store API", emailBody)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message":  "failed send email",
			"response": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "Success receive user data",
		"response": response,
	})
}
