package middlewares

import (
	"context"
	"errors"
	"net/http"
	"project-user/helpers"

	"github.com/labstack/echo/v4"
)

func GetRequestId(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// c.Request().Context().Value(middleware.RequestIdKey).(string)
		// fmt.Println()
		// fmt.Println("ked dini")
		// fmt.Println()
		requestId := c.Request().Header.Get("X-REQUEST-ID")
		if requestId == "" {
			err := errors.New("cannot find requestId")
			helpers.PrintLogToTerminal(err, "requestId")
			errorMessages := helpers.ToErrorMessages("cannot find requestId")
			return c.JSON(http.StatusBadRequest, helpers.Response{Data: nil, Errors: errorMessages})
		}
		// fmt.Println()
		// fmt.Println("requestId:", requestId)
		// fmt.Println()
		ctx := context.WithValue(c.Request().Context(), RequestIdKey, requestId)
		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}
