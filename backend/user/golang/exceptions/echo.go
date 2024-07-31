package exceptions

// func CustomHTTPErrorHandler(err error, c echo.Context) {
// 	// requestId := c.Request().Context().Value(middleware.RequestIdKey).(string)
// 	requestId := c.Request().Header.Get("X-REQUESTID")
// 	helpers.PrintLogToTerminal(err, requestId)
// 	he, ok := err.(*echo.HTTPError)
// 	if !ok {
// 		err = errors.New("cannot convert error to echo.HTTPError")
// 		helpers.PrintLogToTerminal(err, requestId)

// 		response := helpers.Response{
// 			Data:  "",
// 			Error: "internal server error",
// 		}
// 		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 		c.Response().WriteHeader(http.StatusInternalServerError)
// 		json.NewEncoder(c.Response()).Encode(response)
// 		return
// 	}

// 	var errorMessage string
// 	if he.Code == http.StatusNotFound {
// 		errorMessage = "not found"
// 	} else if he.Code == http.StatusMethodNotAllowed {
// 		errorMessage = "method not allowed"
// 	} else {
// 		errorMessage = "internal server error"
// 	}
// 	response := helpers.Response{
// 		Data:  "",
// 		Error: errorMessage,
// 	}
// 	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	c.Response().WriteHeader(he.Code)
// 	json.NewEncoder(c.Response()).Encode(response)
// }
