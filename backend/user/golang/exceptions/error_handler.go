package exceptions

import (
	"net/http"
	"project-user/helpers"

	"github.com/labstack/echo/v4"
)

// func ErrorHandler(c echo.Context, requestId string, err interface{}) error {
// func ErrorHandler(c echo.Context, err interface{}) error {
func ErrorHandler(c echo.Context, err error, requestId string) error {
	helpers.PrintLogToTerminal(err, requestId)
	var httpStatusCode int
	var errorMessage interface{}
	// if exception, ok := err.(BadRequestException); ok {
	// 	httpStatusCode = exception.Code
	// 	errorMessage = exception.Error()
	// } else if exception, ok := err.(NotFoundException); ok {
	// 	httpStatusCode = exception.Code
	// 	errorMessage = exception.Error()
	// } else if exception, ok := err.(ValidationException); ok {
	// 	httpStatusCode = exception.Code
	// 	var exceptionError []helper.Result
	// 	json.Unmarshal([]byte(exception.Error()), &exceptionError)
	// 	errorMessage = exceptionError
	// } else if exception, ok := err.(TimeoutCancelException); ok {
	// 	httpStatusCode = exception.Code
	// 	errorMessage = exception.Error()
	// } else if exception, ok := err.(InternalServerErrorException); ok {
	// 	httpStatusCode = exception.Code
	// 	errorMessage = exception.Error()
	// } else {
	// 	httpStatusCode = http.StatusInternalServerError
	// 	errorMessage = "internal server error"
	// }

	if exception, ok := err.(ResponseException); ok {
		// if exception {

		// }
		httpStatusCode = exception.Code
		// var exceptionErrors []helpers.ErrorMessage
		// json.Unmarshal([]byte(exception.Error()), &exceptionErrors)
		errorMessage = exception.ErrorMessage
	} else {
		httpStatusCode = http.StatusInternalServerError
		// var exceptionErrors []helpers.ErrorMessage
		// var exceptionError helpers.ErrorMessage
		// exceptionError.Field = "message"
		// exceptionError.Message = "internal server error"
		// exceptionErrors = append(exceptionErrors, exceptionError)
		errorMessage = helpers.ToErrorMessages("internal server error")
		// errorMessage = "internal server error"
		// errorMessage = exceptionErrors
	}
	// return helpers.ToResponse(c, httpStatusCode, nil, errorMessage)
	return c.JSON(httpStatusCode, helpers.Response{Data: nil, Errors: errorMessage})
}

// func CheckError(err error) error {
// 	if err == context.Canceled || err == context.DeadlineExceeded {
// 		return NewTimeoutCancelException()
// 	}
// 	return NewInternalServerErrorException()

// }
