package logout_test

import (
	"context"
	"encoding/json"
	"fmt"
	"golang-postgres/commons/setups"
	"golang-postgres/commons/utils"
	"golang-postgres/features/logout/routes"
	"golang-postgres/tests/initialize"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/suite"
)

type LogoutTestSuite struct {
	suite.Suite
	ctx       context.Context
	sessionId string
	redisUtil utils.RedisUtil
	validate  *validator.Validate
	e         *echo.Echo
}

func TestLogoutTestSuite(t *testing.T) {
	suite.Run(t, new(LogoutTestSuite))
}

func (sut *LogoutTestSuite) SetupSuite() {
	sut.T().Log("SetupSuite")
	sut.redisUtil = utils.NewRedisConnection()
	sut.validate = setups.SetValidator()
	sut.e = echo.New()
	sut.e.Use(echomiddleware.Recover())
	sut.e.HTTPErrorHandler = setups.CustomHTTPErrorHandler
	routes.LogoutRoute(sut.e, sut.redisUtil)
}

func (sut *LogoutTestSuite) SetupTest() {
	fmt.Println()
	sut.T().Log("SetupTest")
	sut.sessionId = "sessionId"
	sut.ctx = context.Background()
}

func (sut *LogoutTestSuite) BeforeTest(suiteName, testName string) {
	sut.T().Log("BeforeTest: " + suiteName + " " + testName)
}

func (sut *LogoutTestSuite) Test1LogoutRowsAffectedNotOneInternalServerError() {
	sut.T().Log("Test1LogoutRowsAffectedNotOneInternalServerError")
	req := httptest.NewRequest(http.MethodPost, "/api/v1/users/logout", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("X-REQUEST-ID", "requestId")
	req.Header.Set("X-SESSION-USER-ID", "")
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

func (sut *LogoutTestSuite) Test2LogoutSuccess() {
	sut.T().Log("Test2LogoutSuccess")
	initialize.DelDataRedis(sut.redisUtil.GetClient(), sut.ctx, sut.sessionId)
	initialize.SetDataRedis(sut.redisUtil.GetClient(), sut.ctx, sut.sessionId, "value", time.Duration(0))
	req := httptest.NewRequest(http.MethodPost, "/api/v1/users/logout", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("X-REQUEST-ID", "requestId")
	req.Header.Set("X-SESSION-USER-ID", sut.sessionId)
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
	data := responseBody["data"].(map[string]interface{})
	sut.Equal(data["message"], "logout success")
	sut.Equal(responseBody[""], nil)
}

func (sut *LogoutTestSuite) AfterTest(suiteName, testName string) {
	sut.T().Log("AfterTest: " + suiteName + " " + testName)
}

func (sut *LogoutTestSuite) TearDownTest() {
	sut.T().Log("TearDownTest")
	fmt.Println()
}

func (sut *LogoutTestSuite) TearDownSuite() {
	sut.T().Log("TearDownSuite")
}
