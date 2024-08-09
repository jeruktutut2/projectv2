package middlewares

// import (
// 	"context"

// 	"github.com/labstack/echo/v4"
// )

// func GetSessionIdUser(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		// this session id user can be empty or not, because login not necessarily need it, log out need it for redis, but there is condition if rows affected is not one
// 		sessionIdUser := c.Request().Header.Get("X-SESSION-USER-ID")
// 		ctx := context.WithValue(c.Request().Context(), SessionIdKey, sessionIdUser)
// 		c.SetRequest(c.Request().WithContext(ctx))
// 		return next(c)
// 	}
// }
