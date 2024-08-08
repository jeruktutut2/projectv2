package controllers

import (
	"net/http"
	"project-user/exceptions"
	"project-user/features/register/models"
	"project-user/features/register/services"
	"project-user/helpers"
	"project-user/middlewares"
	"time"

	"github.com/labstack/echo/v4"
)

type RegisterController interface {
	Register(c echo.Context) error
}

type RegisterControllerImplementation struct {
	RegisterService services.RegisterService
}

func NewRegisterController(registerService services.RegisterService) RegisterController {
	return &RegisterControllerImplementation{
		RegisterService: registerService,
	}
}

func (controller *RegisterControllerImplementation) Register(c echo.Context) error {
	requestId := c.Request().Context().Value(middlewares.RequestIdKey).(string)
	var registerUserRequest models.RegisterUserRequest
	err := c.Bind(&registerUserRequest)
	if err != nil {
		return exceptions.ErrorHandler(c, err, requestId)
	}
	nowUnixMili := time.Now().UnixMilli()
	registerResponse, err := controller.RegisterService.Register(c.Request().Context(), requestId, nowUnixMili, registerUserRequest)
	if err != nil {
		return exceptions.ErrorHandler(c, err, requestId)
	}
	return c.JSON(http.StatusCreated, helpers.Response{Data: registerResponse, Errors: nil})
}
