package login_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"project-user/commons/helpers"
	"project-user/commons/setups"
	"project-user/commons/utils"
	"project-user/features/login/routes"
	"project-user/tests/initialize"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/suite"
)

type LoginTestSuite struct {
	suite.Suite
	ctx          context.Context
	requestBody  string
	mysqlUtil    utils.MysqlUtil
	redisUtil    utils.RedisUtil
	validate     *validator.Validate
	e            *echo.Echo
	bcryptHelper helpers.BcryptHelper
	uuidHelper   helpers.UuidHelper
}

func TestLoginTestSuite(t *testing.T) {
	suite.Run(t, new(LoginTestSuite))
}

func (sut *LoginTestSuite) SetupSuite() {
	fmt.Println()
	sut.T().Log("SetupSuite")
	sut.mysqlUtil = utils.NewMysqlConnection("root", "12345", "localhost:3309", "user", 10, 10, 10, 10)
	sut.redisUtil = utils.NewRedisConnection("localhost", "6380", 0)
	sut.validate = setups.SetValidator()
	sut.bcryptHelper = helpers.NewBcryptHelper()
	sut.uuidHelper = helpers.NewUuidHelper()
	sut.e = echo.New()
	sut.e.Use(echomiddleware.Recover())
	sut.e.HTTPErrorHandler = setups.CustomHTTPErrorHandler
	routes.LoginRoute(sut.e, sut.mysqlUtil, sut.redisUtil, sut.validate, sut.bcryptHelper, sut.uuidHelper)
}

func (sut *LoginTestSuite) SetupTest() {
	sut.T().Log("SetupTest")
	sut.ctx = context.Background()
	sut.requestBody = `{
		"email": "email@email.com",
		"password": "password@A1"
	}`
}

func (sut *LoginTestSuite) BeforeTest(suiteName, testName string) {
	sut.T().Log("BeforeTest: " + suiteName + " " + testName)
}

func (sut *LoginTestSuite) Test1LoginRedisRepositoryDelWithSessionIdInternalServerError() {
	sut.T().Log("TestLoginRedisRepositoryDelWithSessionIdInternalServerError")
	initialize.DropTableUserPermission(sut.mysqlUtil.GetDb(), sut.ctx)
	initialize.DropTablePermission(sut.mysqlUtil.GetDb(), sut.ctx)
	initialize.DropTableUser(sut.mysqlUtil.GetDb(), sut.ctx)
	sut.requestBody = `{}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/users/login", strings.NewReader(sut.requestBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("X-REQUEST-ID", "requestId")
	req.Header.Set("X-SESSION-USER-ID", "sessionId")
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
	sut.Equal(errorMessage0["field"], "email")
	sut.Equal(errorMessage0["message"], "is required")
	errorMessage1, _ := errorsResponseBody[1].((map[string]interface{}))
	sut.Equal(errorMessage1["field"], "password")
	sut.Equal(errorMessage1["message"], "is required")
}

func (sut *LoginTestSuite) Test2LoginUserRepositoryFindByEmailInternalServerError() {
	sut.T().Log("TestLoginUserRepositoryFindByEmailInternalServerError")
	initialize.DropTableUserPermission(sut.mysqlUtil.GetDb(), sut.ctx)
	initialize.DropTablePermission(sut.mysqlUtil.GetDb(), sut.ctx)
	initialize.DropTableUser(sut.mysqlUtil.GetDb(), sut.ctx)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/users/login", strings.NewReader(sut.requestBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("X-REQUEST-ID", "requestId")
	req.Header.Set("X-SESSION-USER-ID", "sessionId")
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

func (sut *LoginTestSuite) Test3LoginUserRepositoryFindByEmailBadRequestWrongEmailOrPasswordError() {
	sut.T().Log("TestLoginUserRepositoryFindByEmailBadRequestWrongEmailOrPasswordError")
	initialize.DropTableUserPermission(sut.mysqlUtil.GetDb(), sut.ctx)
	initialize.DropTablePermission(sut.mysqlUtil.GetDb(), sut.ctx)
	initialize.DropTableUser(sut.mysqlUtil.GetDb(), sut.ctx)
	initialize.CreateTableUser(sut.mysqlUtil.GetDb(), sut.ctx)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/users/login", strings.NewReader(sut.requestBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("X-REQUEST-ID", "requestId")
	req.Header.Set("X-SESSION-USER-ID", "sessionId")
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
	sut.Equal(errorMessage0["message"], "wrong email or password")
}

func (sut *LoginTestSuite) Test4LoginBcryptCompareHashAndPasswordBadRequestWrongEmailOrPasswordError() {
	sut.T().Log("TestLoginBcryptCompareHashAndPasswordBadRequestWrongEmailOrPasswordError")
	initialize.DropTableUserPermission(sut.mysqlUtil.GetDb(), sut.ctx)
	initialize.DropTablePermission(sut.mysqlUtil.GetDb(), sut.ctx)
	initialize.DropTableUser(sut.mysqlUtil.GetDb(), sut.ctx)
	initialize.CreateTableUser(sut.mysqlUtil.GetDb(), sut.ctx)
	initialize.CreateDataUser(sut.mysqlUtil.GetDb(), sut.ctx)
	sut.requestBody = `{
		"email": "email@email.com",
		"password": "password@A1-"
	}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/users/login", strings.NewReader(sut.requestBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("X-REQUEST-ID", "requestId")
	req.Header.Set("X-SESSION-USER-ID", "sessionId")
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
	sut.Equal(errorMessage0["message"], "wrong email or password")
}

func (sut *LoginTestSuite) Test5LoginUserPermissionRepositoryFindByUserIdInternalServerError() {
	sut.T().Log("TestLoginUserPermissionRepositoryFindByUserIdInternalServerError")
	initialize.DropTableUserPermission(sut.mysqlUtil.GetDb(), sut.ctx)
	initialize.DropTablePermission(sut.mysqlUtil.GetDb(), sut.ctx)
	initialize.DropTableUser(sut.mysqlUtil.GetDb(), sut.ctx)
	initialize.CreateTableUser(sut.mysqlUtil.GetDb(), sut.ctx)
	initialize.CreateDataUser(sut.mysqlUtil.GetDb(), sut.ctx)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/users/login", strings.NewReader(sut.requestBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("X-REQUEST-ID", "requestId")
	req.Header.Set("X-SESSION-USER-ID", "sessionId")
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

func (sut *LoginTestSuite) Test6LoginSuccess() {
	sut.T().Log("TestLoginSuccess")
	initialize.DropTableUserPermission(sut.mysqlUtil.GetDb(), sut.ctx)
	initialize.DropTablePermission(sut.mysqlUtil.GetDb(), sut.ctx)
	initialize.DropTableUser(sut.mysqlUtil.GetDb(), sut.ctx)
	initialize.CreateTableUser(sut.mysqlUtil.GetDb(), sut.ctx)
	initialize.CreateTablePermission(sut.mysqlUtil.GetDb(), sut.ctx)
	initialize.CreateTableUserPermission(sut.mysqlUtil.GetDb(), sut.ctx)
	initialize.CreateDataUser(sut.mysqlUtil.GetDb(), sut.ctx)
	initialize.CreateDataPermission(sut.mysqlUtil.GetDb(), sut.ctx)
	initialize.CreateDataUserPermission(sut.mysqlUtil.GetDb(), sut.ctx)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/users/login", strings.NewReader(sut.requestBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("X-REQUEST-ID", "requestId")
	req.Header.Set("X-SESSION-USER-ID", "sessionId")
	rec := httptest.NewRecorder()
	sut.e.ServeHTTP(rec, req)
	response := rec.Result()
	sut.Equal(response.StatusCode, http.StatusOK)
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	sut.NotEqual(responseBody["data"], "")
	sut.Equal(responseBody["errors"], nil)
}

func (sut *LoginTestSuite) AfterTest(suiteName, testName string) {
	sut.T().Log("AfterTest: " + suiteName + " " + testName)
}

func (sut *LoginTestSuite) TearDownTest() {
	sut.T().Log("TearDownTest")
	fmt.Println()
}

func (sut *LoginTestSuite) TearDownSuite() {
	sut.T().Log("TearDownSuite")
}
