package register_test

import (
	"context"
	"encoding/json"
	"fmt"
	"golang-postgres/commons/helpers"
	"golang-postgres/commons/setups"
	"golang-postgres/commons/utils"
	"golang-postgres/features/register/routes"
	"golang-postgres/tests/initialize"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/suite"
)

type RegisterTestSuite struct {
	suite.Suite
	ctx          context.Context
	postgresUtil utils.PostgresUtil
	validate     *validator.Validate
	bcryptHelper helpers.BcryptHelper
	e            *echo.Echo
	requestBody  string
}

func TestRegisterTestSuite(t *testing.T) {
	suite.Run(t, new(RegisterTestSuite))
}

func (sut *RegisterTestSuite) SetupSuite() {
	sut.T().Log("SetupSuite")
	sut.postgresUtil = utils.NewPostgresConnection()
	sut.validate = setups.SetValidator()
	sut.bcryptHelper = helpers.NewBcryptHelper()
	sut.e = echo.New()
	sut.e.Use(echomiddleware.Recover())
	sut.e.HTTPErrorHandler = setups.CustomHTTPErrorHandler
	routes.RegisterRoute(sut.e, sut.postgresUtil, sut.validate, sut.bcryptHelper)
}

func (sut *RegisterTestSuite) SetupTest() {
	fmt.Println()
	sut.T().Log("SetupTest")
	sut.ctx = context.Background()
	sut.requestBody = `{
		"username": "username",
		"email": "email@email.com",
		"password": "password@A1",
		"confirmpassword":"password@A1",
		"utc": "+0800"
	}`
}

func (sut *RegisterTestSuite) BeforeTest(suiteName, testName string) {
	sut.T().Log("BeforeTest: " + suiteName + " " + testName)
}

func (sut *RegisterTestSuite) Test1RegisterRegisterUserCannotFindRequestId() {
	sut.T().Log("Test1RegisterRegisterUserCannotFindRequestId")
	requestBody := `{
		"username": "",
		"email": "",
		"password": "",
		"confirmpassword":"",
		"utc": ""
	}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/users/register", strings.NewReader(requestBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	sut.e.ServeHTTP(rec, req)
	response := rec.Result()
	sut.Equal(response.StatusCode, http.StatusBadRequest)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	sut.Equal(responseBody["data"], nil)
	errorsResponseBody := responseBody["errors"].([]interface{})
	errorMessage0, _ := errorsResponseBody[0].((map[string]interface{}))
	sut.Equal(errorMessage0["field"], "message")
	sut.Equal(errorMessage0["message"], "cannot find requestId")
}

func (sut *RegisterTestSuite) Test2RegisterRegisterUserEmptyRequestBody2() {
	sut.T().Log("Test2RegisterRegisterUserEmptyRequestBody2")
	sut.requestBody = ``
	req := httptest.NewRequest(http.MethodPost, "/api/v1/users/register", strings.NewReader(sut.requestBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("X-REQUEST-ID", "requestId")
	rec := httptest.NewRecorder()
	sut.e.ServeHTTP(rec, req)
	response := rec.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	sut.Equal(responseBody["data"], nil)
	errorsResponseBody := responseBody["errors"].([]interface{})
	errorMessage0, _ := errorsResponseBody[0].((map[string]interface{}))
	sut.Equal(errorMessage0["field"], "message")
	sut.Equal(errorMessage0["message"], "internal server error")
}

func (sut *RegisterTestSuite) Test3RegisterRegisterUserEmptyRequestBody() {
	sut.T().Log("Test3RegisterRegisterUserEmptyRequestBody")
	sut.requestBody = `{}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/users/register", strings.NewReader(sut.requestBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("X-REQUEST-ID", "requestId")
	rec := httptest.NewRecorder()
	sut.e.ServeHTTP(rec, req)
	response := rec.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	sut.Equal(responseBody["data"], nil)
	errorsResponseBody := responseBody["errors"].([]interface{})
	errorMessage0, _ := errorsResponseBody[0].((map[string]interface{}))
	sut.Equal(errorMessage0["field"], "username")
	sut.Equal(errorMessage0["message"], "is required")
	errorMessage1, _ := errorsResponseBody[1].((map[string]interface{}))
	sut.Equal(errorMessage1["field"], "email")
	sut.Equal(errorMessage1["message"], "is required")
	errorMessage2, _ := errorsResponseBody[2].((map[string]interface{}))
	sut.Equal(errorMessage2["field"], "password")
	sut.Equal(errorMessage2["message"], "is required")
	errorMessage3, _ := errorsResponseBody[3].((map[string]interface{}))
	sut.Equal(errorMessage3["field"], "confirmpassword")
	sut.Equal(errorMessage3["message"], "is required")
	errorMessage4, _ := errorsResponseBody[4].((map[string]interface{}))
	sut.Equal(errorMessage4["field"], "utc")
	sut.Equal(errorMessage4["message"], "is required")
}

func (sut *RegisterTestSuite) Test4RegisterRegisterUserRequestValidationError() {
	sut.T().Log("Test4RegisterRegisterUserRequestValidationError")
	sut.requestBody = `{
		"username": "",
		"email": "",
		"password": "",
		"confirmpassword":"",
		"utc": ""
	}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/users/register", strings.NewReader(sut.requestBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("X-REQUEST-ID", "requestId")
	rec := httptest.NewRecorder()
	sut.e.ServeHTTP(rec, req)
	response := rec.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	sut.Equal(responseBody["data"], nil)
	errorsResponseBody := responseBody["errors"].([]interface{})
	errorMessage0, _ := errorsResponseBody[0].((map[string]interface{}))
	sut.Equal(errorMessage0["field"], "username")
	sut.Equal(errorMessage0["message"], "is required")
	errorMessage1, _ := errorsResponseBody[1].((map[string]interface{}))
	sut.Equal(errorMessage1["field"], "email")
	sut.Equal(errorMessage1["message"], "is required")
	errorMessage2, _ := errorsResponseBody[2].((map[string]interface{}))
	sut.Equal(errorMessage2["field"], "password")
	sut.Equal(errorMessage2["message"], "is required")
	errorMessage3, _ := errorsResponseBody[3].((map[string]interface{}))
	sut.Equal(errorMessage3["field"], "confirmpassword")
	sut.Equal(errorMessage3["message"], "is required")
	errorMessage4, _ := errorsResponseBody[4].((map[string]interface{}))
	sut.Equal(errorMessage4["field"], "utc")
	sut.Equal(errorMessage4["message"], "is required")
}

func (sut *RegisterTestSuite) Test5RegisterPasswordAndConfirmpasswordIsDifferent() {
	sut.T().Log("Test5RegisterPasswordAndConfirmpasswordIsDifferent")
	sut.requestBody = `{
  		"username": "username1",
  		"email": "email1@email.com",
  		"password": "password@A1",
  		"confirmpassword":"password@A1-",
  		"utc": "+0800"
  	}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/users/register", strings.NewReader(sut.requestBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("X-REQUEST-ID", "requestId")
	rec := httptest.NewRecorder()
	sut.e.ServeHTTP(rec, req)
	response := rec.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	sut.Equal(responseBody["data"], nil)
	errorsResponseBody := responseBody["errors"].([]interface{})
	errorMessage0, _ := errorsResponseBody[0].((map[string]interface{}))
	sut.Equal(errorMessage0["field"], "message")
	sut.Equal(errorMessage0["message"], "password and confirm password is different")
}

func (sut *RegisterTestSuite) Test6RegisterUserRepositoryCountByUsernameInternalServerError() {
	sut.T().Log("Test6RegisterUserRepositoryCountByUsernameInternalServerError")
	initialize.DropTableUserPermission(sut.postgresUtil.GetPool(), sut.ctx)
	initialize.DropTablePermission(sut.postgresUtil.GetPool(), sut.ctx)
	initialize.DropTableUser(sut.postgresUtil.GetPool(), sut.ctx)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/users/register", strings.NewReader(sut.requestBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("X-REQUEST-ID", "requestId")
	rec := httptest.NewRecorder()
	sut.e.ServeHTTP(rec, req)
	response := rec.Result()
	sut.Equal(response.StatusCode, http.StatusInternalServerError)
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	sut.Equal(responseBody["data"], nil)
	errorsResponseBody := responseBody["errors"].([]interface{})
	errorMessage0, _ := errorsResponseBody[0].((map[string]interface{}))
	sut.Equal(errorMessage0["field"], "message")
	sut.Equal(errorMessage0["message"], "internal server error")
}

func (sut *RegisterTestSuite) Test7RegisterUserRepositoryCountByCountByEmailBadRequestUsernameAlreadyExists() {
	sut.T().Log("Test8RegisterUserRepositoryCountByCountByEmailBadRequestUsernameAlreadyExists")
	initialize.DropTableUserPermission(sut.postgresUtil.GetPool(), sut.ctx)
	initialize.DropTablePermission(sut.postgresUtil.GetPool(), sut.ctx)
	initialize.DropTableUser(sut.postgresUtil.GetPool(), sut.ctx)
	initialize.CreateTableUser(sut.postgresUtil.GetPool(), sut.ctx)
	initialize.CreateDataUser(sut.postgresUtil.GetPool(), sut.ctx)
	sut.requestBody = `{
		"username": "username1",
		"email": "email@email.com",
		"password": "password@A1",
		"confirmpassword":"password@A1",
		"utc": "+0800"
	}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/users/register", strings.NewReader(sut.requestBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("X-REQUEST-ID", "requestId")
	rec := httptest.NewRecorder()
	sut.e.ServeHTTP(rec, req)
	response := rec.Result()
	sut.Equal(response.StatusCode, http.StatusBadRequest)
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	sut.Equal(responseBody["data"], nil)
	errorsResponseBody := responseBody["errors"].([]interface{})
	errorMessage0, _ := errorsResponseBody[0].((map[string]interface{}))
	sut.Equal(errorMessage0["field"], "message")
	sut.Equal(errorMessage0["message"], "email already exists")
}

func (sut *RegisterTestSuite) Test8RegisterSuccess() {
	sut.T().Log("Test8RegisterSuccess")
	initialize.DropTableUserPermission(sut.postgresUtil.GetPool(), sut.ctx)
	initialize.DropTablePermission(sut.postgresUtil.GetPool(), sut.ctx)
	initialize.DropTableUser(sut.postgresUtil.GetPool(), sut.ctx)
	initialize.CreateTableUser(sut.postgresUtil.GetPool(), sut.ctx)
	initialize.CreateDataUser(sut.postgresUtil.GetPool(), sut.ctx)
	sut.requestBody = `{
		"username": "username1",
		"email": "email1@email.com",
		"password": "password@A1",
		"confirmpassword":"password@A1",
		"utc": "+0800"
	}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/users/register", strings.NewReader(sut.requestBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("X-REQUEST-ID", "requestId")
	rec := httptest.NewRecorder()
	sut.e.ServeHTTP(rec, req)
	response := rec.Result()
	sut.Equal(response.StatusCode, http.StatusCreated)
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	data := responseBody["data"].(map[string]interface{})
	sut.Equal(data["username"], "username1")
	sut.Equal(data["email"], "email1@email.com")
	sut.Equal(data["utc"], "+0800")
	sut.Equal(responseBody["errors"], nil)

	user := initialize.GetDataUserByEmail(sut.postgresUtil.GetPool(), sut.ctx, data["email"].(string))
	sut.Equal(user.Username.String, data["username"])
	sut.Equal(user.Email.String, data["email"])
	sut.Equal(user.Utc.String, data["utc"])
}

func (sut *RegisterTestSuite) AfterTest(suiteName, testName string) {
	sut.T().Log("AfterTest: " + suiteName + " " + testName)
}

func (sut *RegisterTestSuite) TearDownTest() {
	sut.T().Log("TearDownTest")
	fmt.Println()
}

func (sut *RegisterTestSuite) TearDownSuite() {
	sut.T().Log("TearDownSuite")
}
