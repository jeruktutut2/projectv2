package setups

import (
	"context"
	"encoding/json"
	"errors"
	"golang-postgres/commons/helpers"
	"golang-postgres/commons/utils"
	"net/http"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"

	loginroutes "golang-postgres/features/login/routes"
	logoutroutes "golang-postgres/features/logout/routes"
	registerroutes "golang-postgres/features/register/routes"
)

func SetEcho(postgresUtil utils.PostgresUtil, redisUtil utils.RedisUtil, validate *validator.Validate, bcryptHelper helpers.BcryptHelper, uuidHelper helpers.UuidHelper) (e *echo.Echo) {
	e = echo.New()
	e.Use(echomiddleware.Recover())
	e.HTTPErrorHandler = CustomHTTPErrorHandler

	registerroutes.RegisterRoute(e, postgresUtil, validate, bcryptHelper)
	loginroutes.LoginRoute(e, postgresUtil, redisUtil, validate, bcryptHelper, uuidHelper)
	logoutroutes.LogoutRoute(e, redisUtil)

	return
}

func StartEcho(e *echo.Echo) {
	host := os.Getenv("PROJECT_USER_APPLICATION_HOST")
	go func() {
		if err := e.Start(host); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()
	println(time.Now().String(), "echo: started at", host)
}

func StopEcho(e *echo.Echo) {
	host := os.Getenv("PROJECT_USER_APPLICATION_HOST")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
	cancel()
	println(time.Now().String(), "echo: shutdown properly at", host)
}

func CustomHTTPErrorHandler(err error, c echo.Context) {
	requestId := c.Request().Header.Get("X-REQUESTID")
	helpers.PrintLogToTerminal(err, requestId)
	he, ok := err.(*echo.HTTPError)
	if !ok {
		err = errors.New("cannot convert error to echo.HTTPError")
		helpers.PrintLogToTerminal(err, requestId)
		errorMessages := helpers.ToErrorMessages("internal server error")
		response := helpers.Response{
			Data:   nil,
			Errors: errorMessages,
		}
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		c.Response().WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(c.Response()).Encode(response)
		return
	}

	var message string
	if he.Code == http.StatusNotFound {
		message = "not found"
	} else if he.Code == http.StatusMethodNotAllowed {
		message = "method not allowed"
	} else {
		message = "internal server error"
	}
	errorMessages := helpers.ToErrorMessages(message)
	response := helpers.Response{
		Data:   "",
		Errors: errorMessages,
	}
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c.Response().WriteHeader(he.Code)
	json.NewEncoder(c.Response()).Encode(response)
}
