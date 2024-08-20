package exceptions

import (
	"golang-postgres/commons/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ErrorHandler(c echo.Context, err error, requestId string) error {
	helpers.PrintLogToTerminal(err, requestId)
	var httpStatusCode int
	var errorMessage interface{}
	if exception, ok := err.(ResponseException); ok {
		httpStatusCode = exception.Code
		errorMessage = exception.ErrorMessage
	} else {
		httpStatusCode = http.StatusInternalServerError
		errorMessage = helpers.ToErrorMessages("internal server error")
	}
	return c.JSON(httpStatusCode, helpers.Response{Data: nil, Errors: errorMessage})
}
