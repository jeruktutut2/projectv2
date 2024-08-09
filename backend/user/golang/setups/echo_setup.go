package setups

// import (
// 	"context"
// 	"encoding/json"
// 	"errors"
// 	"net/http"
// 	loginroute "project-user/features/login/routes"
// 	logoutroute "project-user/features/logout/routes"
// 	regiterroute "project-user/features/register/routes"
// 	"project-user/helpers"
// 	"project-user/utils"
// 	"time"

// 	"github.com/go-playground/validator/v10"
// 	"github.com/labstack/echo/v4"
// 	echomiddleware "github.com/labstack/echo/v4/middleware"
// )

// func SetEcho(mysqlUtil utils.MysqlUtil, redisUtil utils.RedisUtil, validate *validator.Validate, bcryptHelper helpers.BcryptHelper, uuidHelper helpers.UuidHelper) (e *echo.Echo) {
// 	e = echo.New()
// 	e.Use(echomiddleware.Recover())
// 	e.HTTPErrorHandler = CustomHTTPErrorHandler

// 	regiterroute.RegisterRoute(e, mysqlUtil, validate, bcryptHelper)
// 	loginroute.LoginRoute(e, mysqlUtil, redisUtil, validate, bcryptHelper, uuidHelper)
// 	logoutroute.LogoutRoute(e, redisUtil)

// 	return
// }

// func StartEcho(e *echo.Echo, host string) {
// 	go func() {
// 		if err := e.Start(host); err != nil && err != http.ErrServerClosed {
// 			e.Logger.Fatal("shutting down the server")
// 		}
// 	}()
// 	println(time.Now().String(), "echo: started at", host)
// }

// func StopEcho(e *echo.Echo, host string) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	// defer cancel()
// 	if err := e.Shutdown(ctx); err != nil {
// 		e.Logger.Fatal(err)
// 	}
// 	cancel()
// 	println(time.Now().String(), "echo: shutdown properly at", host)
// }

// func CustomHTTPErrorHandler(err error, c echo.Context) {
// 	requestId := c.Request().Header.Get("X-REQUESTID")
// 	helpers.PrintLogToTerminal(err, requestId)
// 	he, ok := err.(*echo.HTTPError)
// 	if !ok {
// 		err = errors.New("cannot convert error to echo.HTTPError")
// 		helpers.PrintLogToTerminal(err, requestId)
// 		errorMessages := helpers.ToErrorMessages("internal server error")
// 		response := helpers.Response{
// 			Data:   nil,
// 			Errors: errorMessages,
// 		}
// 		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 		c.Response().WriteHeader(http.StatusInternalServerError)
// 		json.NewEncoder(c.Response()).Encode(response)
// 		return
// 	}

// 	var message string
// 	if he.Code == http.StatusNotFound {
// 		message = "not found"
// 	} else if he.Code == http.StatusMethodNotAllowed {
// 		message = "method not allowed"
// 	} else {
// 		message = "internal server error"
// 	}
// 	errorMessages := helpers.ToErrorMessages(message)
// 	response := helpers.Response{
// 		Data:   "",
// 		Errors: errorMessages,
// 	}
// 	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	c.Response().WriteHeader(he.Code)
// 	json.NewEncoder(c.Response()).Encode(response)
// }
