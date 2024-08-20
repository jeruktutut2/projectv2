package controllers

import (
	"golang-postgres/commons/exceptions"
	"golang-postgres/commons/helpers"
	"golang-postgres/commons/middlewares"
	"golang-postgres/features/login/models"
	"golang-postgres/features/login/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type LoginController interface {
	Login(c echo.Context) error
}

type LoginControllerImplementation struct {
	LoginService services.LoginService
}

func NewLoginController(loginService services.LoginService) LoginController {
	return &LoginControllerImplementation{
		LoginService: loginService,
	}
}

func (controller *LoginControllerImplementation) Login(c echo.Context) error {
	requestId := c.Request().Context().Value(middlewares.RequestIdKey).(string)
	sessionIdUser := c.Request().Context().Value(middlewares.SessionIdKey).(string)
	var loginUserRequest models.LoginUserRequest
	err := c.Bind(&loginUserRequest)
	if err != nil {
		return exceptions.ErrorHandler(c, err, requestId)
	}
	loginResponse, err := controller.LoginService.Login(c.Request().Context(), requestId, sessionIdUser, loginUserRequest)
	if err != nil {
		return exceptions.ErrorHandler(c, err, requestId)
	}
	return c.JSON(http.StatusOK, helpers.Response{Data: loginResponse, Errors: nil})
}
