package middlewares

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"project-user/helpers"
	"time"

	"github.com/labstack/echo/v4"
)

type responseBodyWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w *responseBodyWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func PrintRequestResponseLog(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		datetimeNowRequest := time.Now()
		requestMethod := c.Request().Method
		// requestId := c.Request().Context().Value(RequestIdKey).(string)
		var err error
		var errorMessages []helpers.ErrorMessage
		var httpResponseCode int

		// why put this, because i wanna log request when there is no request id
		requestId := c.Request().Header.Get("X-REQUEST-ID")
		if requestId == "" {
			errRequestId := errors.New("cannot find requestId")
			helpers.PrintLogToTerminal(errRequestId, "requestId")
			httpResponseCode = http.StatusBadRequest
			// fmt.Println()
			// fmt.Println("httpResponseCode:", httpResponseCode)
			// fmt.Println()
			errorMessages = helpers.ToErrorMessages("cannot find requestId")
			err = errRequestId
		}
		ctx := context.WithValue(c.Request().Context(), RequestIdKey, requestId)
		c.SetRequest(c.Request().WithContext(ctx))

		var requestBody string
		requestBody = `""`
		// fmt.Println()
		// fmt.Println("c.Request().GetBody:", c.Request().GetBody)
		// body, err := io.ReadAll(c.Request().Body)
		// fmt.Println("body, err:", string(body), err)

		// fmt.Println()
		// body, err := io.ReadAll(c.Request().Body)
		// if err != nil {
		// 	helpers.PrintLogToTerminal(err, requestId)
		// 	errorMessages := helpers.ToErrorMessages("internal server error")
		// 	return c.JSON(http.StatusInternalServerError, helpers.Response{Data: nil, Errors: errorMessages})
		// }
		// jsonRequestBodyMap := make(map[string]interface{})
		// errJsonREquestBodyMap := json.NewDecoder(c.Request().Body).Decode(&jsonRequestBodyMap)
		// // fmt.Println("jsonRequestBodyMap1:", jsonRequestBodyMap)
		// if errJsonREquestBodyMap != nil {
		// 	helpers.PrintLogToTerminal(errJsonREquestBodyMap, requestId)
		// 	errorMessages = helpers.ToErrorMessages("internal server error")
		// 	err = errJsonREquestBodyMap
		// 	httpResponseCode = http.StatusInternalServerError
		// 	// fmt.Println()
		// 	// fmt.Println("httpResponseCode:", httpResponseCode)
		// 	// fmt.Println()
		// 	// return c.JSON(http.StatusInternalServerError, helpers.Response{Data: nil, Errors: errorMessages})
		// }
		// jsonRequestBodyMap2 := make(map[string]interface{})
		// json.NewDecoder(c.Request().Body).Decode(&jsonRequestBodyMap2)
		// fmt.Println("jsonRequestBodyMap2:", jsonRequestBodyMap2)
		// fmt.Println("jsonRequestBodyMap:", jsonRequestBodyMap)
		// if c.Request().GetBody != nil {
		// if len(body) != 0 {
		body, errJsonRequestBody := io.ReadAll(c.Request().Body)
		if errJsonRequestBody != nil {
			helpers.PrintLogToTerminal(errJsonRequestBody, requestId)
			errorMessages = helpers.ToErrorMessages("internal server error")
			err = errJsonRequestBody
			httpResponseCode = http.StatusInternalServerError
		}
		if len(body) == 0 {
			errLenBody := errors.New("json len body equal to 0")
			helpers.PrintLogToTerminal(errLenBody, requestId)
			err = errLenBody
			helpers.ToErrorMessages("internal server error")
		}
		jsonRequestBodyMap := make(map[string]interface{})
		errJsonRequestBodyMap := json.Unmarshal(body, &jsonRequestBodyMap)
		if errJsonRequestBodyMap != nil {
			helpers.PrintLogToTerminal(errJsonRequestBodyMap, requestId)
			err = errJsonRequestBodyMap
			errorMessages = helpers.ToErrorMessages("internal server error")
		}
		c.Request().Body = io.NopCloser(bytes.NewBuffer(body))

		if jsonRequestBodyMap != nil {

			// jsonRequestBodyMap := make(map[string]interface{})
			// err := json.NewDecoder(c.Request().Body).Decode(&jsonRequestBodyMap)
			// if err != nil {
			// 	helpers.PrintLogToTerminal(err, requestId)
			// 	// return helpers.ToResponse(c, http.StatusInternalServerError, nil, "internal server error")
			// 	errorMessages := helpers.ToErrorMessages("internal server error")
			// 	return c.JSON(http.StatusInternalServerError, helpers.Response{Data: nil, Errors: errorMessages})
			// }

			// jsonRequestBodyMap := make(map[string]interface{})
			// var jsonRequestBodyMap map[string]interface{}
			// err = json.Unmarshal(body, &jsonRequestBodyMap)
			// if err != nil {
			// 	helpers.PrintLogToTerminal(err, requestId)
			// 	errorMessages := helpers.ToErrorMessages("internal server error")
			// 	return c.JSON(http.StatusInternalServerError, helpers.Response{Data: nil, Errors: errorMessages})
			// }
			// fmt.Println("jsonRequestBodyMap:", jsonRequestBodyMap)

			if c.Request().URL.Path == "/api/v1/users/register" {
				// delete(jsonRequestBodyMap, "password")
				// delete(jsonRequestBodyMap, "confirmpassword")
				// if _, ok := jsonRequestBodyMap["password"]; ok {
				delete(jsonRequestBodyMap, "password")
				// }
				// if _, ok := jsonRequestBodyMap["confirmpassword"]; ok {
				delete(jsonRequestBodyMap, "confirmpassword")
				// }
			} else if c.Request().URL.Path == "/api/v1/users/login" {
				// delete(jsonRequestBodyMap, "password")
				// if _, ok := jsonRequestBodyMap["password"]; ok {
				delete(jsonRequestBodyMap, "password")
				// }
			}

			// var jsonRequestBodyByte []byte
			jsonRequestBodyByte, errJsonRequestBodyByte := json.Marshal(jsonRequestBodyMap)
			if errJsonRequestBodyByte != nil {
				helpers.PrintLogToTerminal(errJsonRequestBodyByte, requestId)
				// return helpers.ToResponse(c, http.StatusInternalServerError, nil, "cannot convert request body to json")
				errorMessages = helpers.ToErrorMessages("internal server error")
				httpResponseCode = http.StatusInternalServerError
				err = errJsonRequestBodyByte
				// fmt.Println()
				// fmt.Println("httpResponseCode:", httpResponseCode)
				// fmt.Println()
				// return c.JSON(http.StatusInternalServerError, helpers.Response{Data: nil, Errors: errorMessages})
			}
			requestBody = string(jsonRequestBodyByte)
		}

		host := c.Request().Host
		protocol := ""
		if c.Request().TLS == nil {
			protocol = "http"
		} else {
			protocol = "https"
		}
		urlPath := c.Request().URL.Path
		userAgent := c.Request().Header.Get("User-Agent")
		remoteAddr := c.Request().RemoteAddr
		forwardedFor := c.Request().Header.Get("X-Forwarded-For")

		requestLog := `{"requestTime": "` + datetimeNowRequest.String() + `", "app": "project-backend-user", "method": "` + requestMethod + `","requestId":"` + requestId + `","host": "` + host + `","urlPath":"` + urlPath + `","protocol":"` + protocol + `","body": ` + requestBody + `, "userAgent": "` + userAgent + `", "remoteAddr": "` + remoteAddr + `", "forwardedFor": "` + forwardedFor + `"}`
		fmt.Println(requestLog)

		// fmt.Println()
		// fmt.Println("httpResponseCode1:", httpResponseCode, err)
		// fmt.Println()
		if err != nil {
			// fmt.Println()
			// fmt.Println("httpResponseCode2:", httpResponseCode, err)
			// fmt.Println()
			return c.JSON(httpResponseCode, helpers.Response{Data: nil, Errors: errorMessages})
		}

		// so i can catch the response body
		resBody := new(bytes.Buffer)
		mw := io.MultiWriter(c.Response().Writer, resBody)
		writer := &responseBodyWriter{Writer: mw, ResponseWriter: c.Response().Writer}
		c.Response().Writer = writer

		err = next(c)
		if err != nil {
			c.Error(err)
		}

		responseBody := resBody.String()
		log := `{"responseTime": "` + time.Now().String() + `", "app": "project-backend-user", "requestId": "` + requestId + `", "response": ` + responseBody + `}`
		fmt.Println(log)
		return nil
	}
}

func PrintRequestResponseLogWithNoRequestBody(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		datetimeNowRequest := time.Now()
		requestMethod := c.Request().Method
		// requestId := c.Request().Context().Value(RequestIdKey).(string)
		var err error
		var errorMessages []helpers.ErrorMessage
		var httpResponseCode int

		// why put this, because i wanna log request when there is no request id
		requestId := c.Request().Header.Get("X-REQUEST-ID")
		if requestId == "" {
			errRequestId := errors.New("cannot find requestId")
			helpers.PrintLogToTerminal(errRequestId, "requestId")
			httpResponseCode = http.StatusBadRequest
			// fmt.Println()
			// fmt.Println("httpResponseCode:", httpResponseCode)
			// fmt.Println()
			errorMessages = helpers.ToErrorMessages("cannot find requestId")
			err = errRequestId
		}
		ctx := context.WithValue(c.Request().Context(), RequestIdKey, requestId)
		c.SetRequest(c.Request().WithContext(ctx))

		var requestBody string
		requestBody = `""`
		// fmt.Println()
		// fmt.Println("c.Request().GetBody:", c.Request().GetBody)
		// body, err := io.ReadAll(c.Request().Body)
		// fmt.Println("body, err:", string(body), err)

		// fmt.Println()
		// body, err := io.ReadAll(c.Request().Body)
		// if err != nil {
		// 	helpers.PrintLogToTerminal(err, requestId)
		// 	errorMessages := helpers.ToErrorMessages("internal server error")
		// 	return c.JSON(http.StatusInternalServerError, helpers.Response{Data: nil, Errors: errorMessages})
		// }
		// jsonRequestBodyMap := make(map[string]interface{})
		// errJsonREquestBodyMap := json.NewDecoder(c.Request().Body).Decode(&jsonRequestBodyMap)
		// // fmt.Println("jsonRequestBodyMap1:", jsonRequestBodyMap)
		// if errJsonREquestBodyMap != nil {
		// 	helpers.PrintLogToTerminal(errJsonREquestBodyMap, requestId)
		// 	errorMessages = helpers.ToErrorMessages("internal server error")
		// 	err = errJsonREquestBodyMap
		// 	httpResponseCode = http.StatusInternalServerError
		// 	// fmt.Println()
		// 	// fmt.Println("httpResponseCode:", httpResponseCode)
		// 	// fmt.Println()
		// 	// return c.JSON(http.StatusInternalServerError, helpers.Response{Data: nil, Errors: errorMessages})
		// }
		// jsonRequestBodyMap2 := make(map[string]interface{})
		// json.NewDecoder(c.Request().Body).Decode(&jsonRequestBodyMap2)
		// fmt.Println("jsonRequestBodyMap2:", jsonRequestBodyMap2)
		// fmt.Println("jsonRequestBodyMap:", jsonRequestBodyMap)
		// if c.Request().GetBody != nil {
		// if len(body) != 0 {
		// body, errJsonRequestBody := io.ReadAll(c.Request().Body)
		// if errJsonRequestBody != nil {
		// 	helpers.PrintLogToTerminal(errJsonRequestBody, requestId)
		// 	errorMessages = helpers.ToErrorMessages("internal server error")
		// 	err = errJsonRequestBody
		// 	httpResponseCode = http.StatusInternalServerError
		// }
		// if len(body) == 0 {
		// 	errLenBody := errors.New("json len body equal to 0")
		// 	helpers.PrintLogToTerminal(errLenBody, requestId)
		// 	err = errLenBody
		// 	helpers.ToErrorMessages("internal server error")
		// }
		// jsonRequestBodyMap := make(map[string]interface{})
		// errJsonRequestBodyMap := json.Unmarshal(body, &jsonRequestBodyMap)
		// if errJsonRequestBodyMap != nil {
		// 	helpers.PrintLogToTerminal(errJsonRequestBodyMap, requestId)
		// 	err = errJsonRequestBodyMap
		// 	errorMessages = helpers.ToErrorMessages("internal server error")
		// }
		// c.Request().Body = io.NopCloser(bytes.NewBuffer(body))

		// if jsonRequestBodyMap != nil {

		// 	// jsonRequestBodyMap := make(map[string]interface{})
		// 	// err := json.NewDecoder(c.Request().Body).Decode(&jsonRequestBodyMap)
		// 	// if err != nil {
		// 	// 	helpers.PrintLogToTerminal(err, requestId)
		// 	// 	// return helpers.ToResponse(c, http.StatusInternalServerError, nil, "internal server error")
		// 	// 	errorMessages := helpers.ToErrorMessages("internal server error")
		// 	// 	return c.JSON(http.StatusInternalServerError, helpers.Response{Data: nil, Errors: errorMessages})
		// 	// }

		// 	// jsonRequestBodyMap := make(map[string]interface{})
		// 	// var jsonRequestBodyMap map[string]interface{}
		// 	// err = json.Unmarshal(body, &jsonRequestBodyMap)
		// 	// if err != nil {
		// 	// 	helpers.PrintLogToTerminal(err, requestId)
		// 	// 	errorMessages := helpers.ToErrorMessages("internal server error")
		// 	// 	return c.JSON(http.StatusInternalServerError, helpers.Response{Data: nil, Errors: errorMessages})
		// 	// }
		// 	// fmt.Println("jsonRequestBodyMap:", jsonRequestBodyMap)

		// 	if c.Request().URL.Path == "/api/v1/users/register" {
		// 		// delete(jsonRequestBodyMap, "password")
		// 		// delete(jsonRequestBodyMap, "confirmpassword")
		// 		// if _, ok := jsonRequestBodyMap["password"]; ok {
		// 		delete(jsonRequestBodyMap, "password")
		// 		// }
		// 		// if _, ok := jsonRequestBodyMap["confirmpassword"]; ok {
		// 		delete(jsonRequestBodyMap, "confirmpassword")
		// 		// }
		// 	} else if c.Request().URL.Path == "/api/v1/users/login" {
		// 		// delete(jsonRequestBodyMap, "password")
		// 		// if _, ok := jsonRequestBodyMap["password"]; ok {
		// 		delete(jsonRequestBodyMap, "password")
		// 		// }
		// 	}

		// 	// var jsonRequestBodyByte []byte
		// 	jsonRequestBodyByte, errJsonRequestBodyByte := json.Marshal(jsonRequestBodyMap)
		// 	if errJsonRequestBodyByte != nil {
		// 		helpers.PrintLogToTerminal(errJsonRequestBodyByte, requestId)
		// 		// return helpers.ToResponse(c, http.StatusInternalServerError, nil, "cannot convert request body to json")
		// 		errorMessages = helpers.ToErrorMessages("internal server error")
		// 		httpResponseCode = http.StatusInternalServerError
		// 		err = errJsonRequestBodyByte
		// 		// fmt.Println()
		// 		// fmt.Println("httpResponseCode:", httpResponseCode)
		// 		// fmt.Println()
		// 		// return c.JSON(http.StatusInternalServerError, helpers.Response{Data: nil, Errors: errorMessages})
		// 	}
		// 	requestBody = string(jsonRequestBodyByte)
		// }

		host := c.Request().Host
		protocol := ""
		if c.Request().TLS == nil {
			protocol = "http"
		} else {
			protocol = "https"
		}
		urlPath := c.Request().URL.Path
		userAgent := c.Request().Header.Get("User-Agent")
		remoteAddr := c.Request().RemoteAddr
		forwardedFor := c.Request().Header.Get("X-Forwarded-For")

		requestLog := `{"requestTime": "` + datetimeNowRequest.String() + `", "app": "project-backend-user", "method": "` + requestMethod + `","requestId":"` + requestId + `","host": "` + host + `","urlPath":"` + urlPath + `","protocol":"` + protocol + `","body": ` + requestBody + `, "userAgent": "` + userAgent + `", "remoteAddr": "` + remoteAddr + `", "forwardedFor": "` + forwardedFor + `"}`
		fmt.Println(requestLog)

		// fmt.Println()
		// fmt.Println("httpResponseCode1:", httpResponseCode, err)
		// fmt.Println()
		if err != nil {
			// fmt.Println()
			// fmt.Println("httpResponseCode2:", httpResponseCode, err)
			// fmt.Println()
			return c.JSON(httpResponseCode, helpers.Response{Data: nil, Errors: errorMessages})
		}

		// so i can catch the response body
		resBody := new(bytes.Buffer)
		mw := io.MultiWriter(c.Response().Writer, resBody)
		writer := &responseBodyWriter{Writer: mw, ResponseWriter: c.Response().Writer}
		c.Response().Writer = writer

		err = next(c)
		if err != nil {
			c.Error(err)
		}

		responseBody := resBody.String()
		log := `{"responseTime": "` + time.Now().String() + `", "app": "project-backend-user", "requestId": "` + requestId + `", "response": ` + responseBody + `}`
		fmt.Println(log)
		return nil
	}
}
