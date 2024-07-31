package exceptions

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"project-user/helpers"
)

type ResponseException struct {
	Code         int
	ErrorMessage interface{}
	Message      string
}

func NewResponseException(code int, errorMessage interface{}, message string) ResponseException {
	return ResponseException{Code: code, ErrorMessage: errorMessage, Message: message}
}

func (exception ResponseException) Error() string {
	return exception.Message
}

// type ErrorMessage struct {
// 	Field   string `json:"field"`
// 	Message string `json:"message"`
// }

// func toErrorMessage(field string, message string) ([]byte, error) {
// 	var errorMessages []helpers.ErrorMessage
// 	var errorMessage helpers.ErrorMessage
// 	errorMessage.Field = field
// 	errorMessage.Message = message
// 	errorMessages = append(errorMessages, errorMessage)
// 	return json.Marshal(errorMessages)
// }

func toErrorMessageRequestTimeout(requestId string) error {
	// errorMessagesByte, err := helpers.ToErrorMessagesString("message", "time out or user cancel the request")
	// if err != nil {
	// 	helpers.PrintLogToTerminal(err, requestId)
	// 	return toErrorMessageInternalServerError(requestId)
	// }
	message := "time out or user cancel the request"
	errorMessages := helpers.ToErrorMessages(message)
	// return NewResponseException(http.StatusRequestTimeout, errorMessages, string(errorMessagesByte))
	return NewResponseException(http.StatusRequestTimeout, errorMessages, message)
}

func toErrorMessageInternalServerError(requestId string) error {
	// errorMessagesByte, err := helpers.ToErrorMessagesString("message", "internal server error")
	// if err != nil {
	// 	helpers.PrintLogToTerminal(err, requestId)
	// }
	message := "internal server error"
	errorMessages := helpers.ToErrorMessages(message)
	return NewResponseException(http.StatusInternalServerError, errorMessages, message)
}

func CheckError(err error, requestId string) error {
	helpers.PrintLogToTerminal(err, requestId)
	if err == context.Canceled || err == context.DeadlineExceeded {
		return toErrorMessageRequestTimeout(requestId)
	} else {
		return toErrorMessageInternalServerError(requestId)
	}
}

func ToResponseExceptionRequestValidation(requestId string, validationErrorMessages []helpers.ErrorMessage) error {
	var validationErrorMessageByte []byte
	validationErrorMessageByte, err := json.Marshal(validationErrorMessages)
	if err != nil {
		return CheckError(err, requestId)
	}
	err = errors.New(string(validationErrorMessageByte))
	helpers.PrintLogToTerminal(err, requestId)
	// return NewResponseException(http.StatusBadRequest, validationErrorMessages, string(validationErrorMessageByte))
	return NewResponseException(http.StatusBadRequest, validationErrorMessages, "validation error")
}

func ToResponseException(err error, requestId string, httpCode int, message string) error {
	helpers.PrintLogToTerminal(err, requestId)
	// field := "message"
	// errorMessagesByte, err := helpers.ToErrorMessagesString(field, message)
	// if err != nil {
	// 	helpers.PrintLogToTerminal(err, requestId)
	// 	return CheckError(err, requestId)
	// }
	errorMessages := helpers.ToErrorMessages(message)
	// return NewResponseException(httpCode, errorMessages, string(errorMessagesByte))
	return NewResponseException(httpCode, errorMessages, message)
}
