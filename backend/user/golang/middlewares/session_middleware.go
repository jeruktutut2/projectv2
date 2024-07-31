package middlewares

import (
	"context"

	"github.com/labstack/echo/v4"
)

func GetSessionIdUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// this session id user can be empty or not, becouse login not necessarily need it, log out need it for redis, but there is condition if rows affected is not one
		// requestId := c.Request().Context().Value(SessionIdKey).(string)
		sessionIdUser := c.Request().Header.Get("X-SESSION-USER-ID")
		// if sessionIdUser == "" {
		// 	errSessionIdUser := errors.New("cannot find sessionId")
		// 	helpers.PrintLogToTerminal(errSessionIdUser, requestId)
		// 	errorsMessage := helpers.ToErrorMessages("cannot find sessionId")
		// 	// return c.JSON(http.StatusBadRequest, helpers.Response{Data: nil, Errors: errorsMessage})
		// 	// return c.JSON(httpStatusCode, helpers.Response{Data: nil, Errors: errorMessage})
		// }
		ctx := context.WithValue(c.Request().Context(), SessionIdKey, sessionIdUser)
		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}
