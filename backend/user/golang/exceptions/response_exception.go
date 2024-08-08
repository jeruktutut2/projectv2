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

func toErrorMessageRequestTimeout(requestId string) error {
	message := "time out or user cancel the request"
	errorMessages := helpers.ToErrorMessages(message)
	return NewResponseException(http.StatusRequestTimeout, errorMessages, message)
}

func toErrorMessageInternalServerError(requestId string) error {
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
	return NewResponseException(http.StatusBadRequest, validationErrorMessages, "validation error")
}

func ToResponseException(err error, requestId string, httpCode int, message string) error {
	helpers.PrintLogToTerminal(err, requestId)
	errorMessages := helpers.ToErrorMessages(message)
	return NewResponseException(httpCode, errorMessages, message)
}
