package logout_test

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"project-user/features/logout/routes"
	"project-user/setups"
	"project-user/tests/initialize"
	"project-user/utils"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/suite"
)

type LogoutTestSuite struct {
	suite.Suite
	// mysqlUtil utils.MysqlUtil
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
	// sut.mysqlUtil = utils.NewMysqlConnection("root", "12345", "localhost:3309", "user", 10, 10, 10, 10)
	sut.redisUtil = utils.NewRedisConnection("localhost", "6380", 0)
	sut.validate = setups.SetValidator()
	sut.e = echo.New()
	sut.e.Use(echomiddleware.Recover())
	sut.e.HTTPErrorHandler = setups.CustomHTTPErrorHandler
	// sut.e.Use(middlewares.PrintRequestResponseLog)
	// sut.e.Use(middlewares.GetSessionIdUser)
	routes.LogoutRoute(sut.e, sut.redisUtil)
}

func (sut *LogoutTestSuite) SetupTest() {
	sut.T().Log("SetupTest")
	sut.sessionId = "sessionId"
	sut.ctx = context.Background()
}

func (sut *LogoutTestSuite) BeforeTest(suiteName, testName string) {
	sut.T().Log("BeforeTest: " + suiteName + " " + testName)
}

func (sut *LogoutTestSuite) TestLogoutRowsAffectedNotOneInternalServerError() {
	sut.T().Log("TestLogoutRowsAffectedNotOneInternalServerError")
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
	// fmt.Println("responseBody:", responseBody)
	sut.Equal(responseBody["data"], nil)
	errorsResponseBody := responseBody["errors"].([]interface{})
	errorMessage0, _ := errorsResponseBody[0].((map[string]interface{}))
	sut.Equal(errorMessage0["field"], "message")
	sut.Equal(errorMessage0["message"], "internal server error")
}

func (sut *LogoutTestSuite) TestLogoutSuccess() {
	sut.T().Log("TestLogoutSuccess")
	initialize.DelDataRedis(sut.redisUtil.GetClient(), sut.ctx, sut.sessionId)
	initialize.SetDataRedis(sut.redisUtil.GetClient(), sut.ctx, sut.sessionId, "value", time.Duration(0))
	req := httptest.NewRequest(http.MethodPost, "/api/v1/users/logout", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	// req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
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
	// fmt.Println("responseBody:", responseBody)
	data := responseBody["data"].(map[string]interface{})
	sut.Equal(data["message"], "logout success")
	sut.Equal(responseBody[""], nil)
}

func (sut *LogoutTestSuite) AfterTest(suiteName, testName string) {
	sut.T().Log("AfterTest: " + suiteName + " " + testName)
}

func (sut *LogoutTestSuite) TearDownTest() {
	sut.T().Log("TearDownTest")
}

func (sut *LogoutTestSuite) TearDownSuite() {
	sut.T().Log("TearDownSuite")
}
