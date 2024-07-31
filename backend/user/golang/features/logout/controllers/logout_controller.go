package controllers

import (
	"net/http"
	"project-user/exceptions"
	"project-user/features/logout/services"
	"project-user/helpers"
	"project-user/middlewares"

	"github.com/labstack/echo/v4"
)

type LogoutController interface {
	Logout(c echo.Context) error
}

type LogoutControllerImplementation struct {
	LogoutService services.LogoutService
}

func NewLogoutController(logoutService services.LogoutService) LogoutController {
	return &LogoutControllerImplementation{
		LogoutService: logoutService,
	}
}

func (controller *LogoutControllerImplementation) Logout(c echo.Context) error {
	requestId := c.Request().Context().Value(middlewares.RequestIdKey).(string)
	sessioIdUser := c.Request().Context().Value(middlewares.SessionIdKey).(string)
	err := controller.LogoutService.Logout(c.Request().Context(), requestId, sessioIdUser)
	if err != nil {
		return exceptions.ErrorHandler(c, err, requestId)
	}
	respone := helpers.ResponseMessage{Message: "logout success"}
	return c.JSON(http.StatusOK, helpers.Response{Data: respone, Errors: nil})
}
