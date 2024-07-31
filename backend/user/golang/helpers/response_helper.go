package helpers

type Response struct {
	Data   interface{} `json:"data"`
	Errors interface{} `json:"errors"`
}

type ResponseMessage struct {
	Message string `json:"message"`
}

// type ErrorMessage struct {
// 	Field   string `json:"field"`
// 	Message string `json:"message"`
// }

// func ToErrorMessages(field string, message string) (errorMessages []ErrorMessage) {
// 	// var errorMessages []ErrorMessage
// 	var errorMessage ErrorMessage
// 	errorMessage.Field = field
// 	errorMessage.Message = message
// 	errorMessages = append(errorMessages, errorMessage)
// 	return
// }

// func ToErrorMessagesString(field string, message string) ([]byte, error) {
// 	var errorMessages []ErrorMessage
// 	var errorMessage ErrorMessage
// 	errorMessage.Field = field
// 	errorMessage.Message = message
// 	errorMessages = append(errorMessages, errorMessage)
// 	return json.Marshal(errorMessages)
// }

// func ToResponse(c echo.Context, httpStatusCode int, requestId string, responseData interface{}, responseError interface{}) error {
// func ToResponse(c echo.Context, httpStatusCode int, responseData interface{}, responseError interface{}) error {
// 	r := Response{
// 		Data:   responseData,
// 		Errors: responseError,
// 	}
// 	// respByte, err := json.Marshal(r)
// 	// if err != nil {
// 	// 	PrintLogToTerminal(err, requestId)
// 	// 	response := `{"data": null, "error": "internal server error"}`
// 	// 	return c.JSON(http.StatusInternalServerError, response)
// 	// }
// 	// responseBody := string(respByte)
// 	// log := `{"responseTime": "` + time.Now().String() + `", "app": "project-user", "requestId": "` + requestId + `", "response": ` + responseBody + `}`
// 	// fmt.Println(log)
// 	return c.JSON(httpStatusCode, r)
// }
