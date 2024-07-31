package helpers

type ErrorMessage struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// func ToErrorMessages(field string, message string) (errorMessages []ErrorMessage) {
func ToErrorMessages(message string) (errorMessages []ErrorMessage) {
	// var errorMessages []ErrorMessage
	var errorMessage ErrorMessage
	errorMessage.Field = "message"
	errorMessage.Message = message
	errorMessages = append(errorMessages, errorMessage)
	return
}

// func ToErrorMessagesString(field string, message string) ([]byte, error) {
// 	var errorMessages []ErrorMessage
// 	var errorMessage ErrorMessage
// 	errorMessage.Field = field
// 	errorMessage.Message = message
// 	errorMessages = append(errorMessages, errorMessage)
// 	return json.Marshal(errorMessages)
// }
