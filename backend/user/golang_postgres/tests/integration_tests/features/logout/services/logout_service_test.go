package services_test

import (
	"context"
	"fmt"
	"golang-postgres/commons/utils"
	"golang-postgres/features/login/repositories"
	"golang-postgres/features/logout/services"
	"golang-postgres/tests/initialize"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/suite"
)

type LogoutServiceTestSuite struct {
	suite.Suite
	ctx             context.Context
	requestId       string
	sessionId       string
	redisUtil       utils.RedisUtil
	redisRepository repositories.RedisRepository
	logoutService   services.LogoutService
}

func TestLogoutTestSuite(t *testing.T) {
	suite.Run(t, new(LogoutServiceTestSuite))
}

func (sut *LogoutServiceTestSuite) SetupSuite() {
	sut.T().Log("SetupSuite")
	sut.redisUtil = utils.NewRedisConnection()
	sut.redisRepository = repositories.NewRedisRepository()
}

func (sut *LogoutServiceTestSuite) SetupTest() {
	fmt.Println()
	sut.T().Log("SetupTest")
	sut.ctx = context.Background()
	sut.requestId = "requestId"
	sut.sessionId = "sessionId"
	sut.logoutService = services.NewLogoutService(sut.redisUtil, sut.redisRepository)
}

func (sut *LogoutServiceTestSuite) BeforeTest(suiteName, testName string) {
	sut.T().Log("BeforeTest: " + suiteName + " " + testName)
}

func (sut *LogoutServiceTestSuite) Test1LogoutRowsAffectedNotOneInternalServerError() {
	sut.T().Log("Test1LogoutRowsAffectedNotOneInternalServerError")
	err := sut.logoutService.Logout(sut.ctx, sut.requestId, sut.sessionId)
	sut.Equal(err.Error(), "internal server error")
}

func (sut *LogoutServiceTestSuite) Test2LogoutSuccess() {
	sut.T().Log("Test2LogoutSuccess")
	initialize.DelDataRedis(sut.redisUtil.GetClient(), sut.ctx, sut.sessionId)
	initialize.SetDataRedis(sut.redisUtil.GetClient(), sut.ctx, sut.sessionId, "value", time.Duration(0))
	err := sut.logoutService.Logout(sut.ctx, sut.requestId, sut.sessionId)
	sut.Equal(err, nil)
	_, err = initialize.GetDataRedis(sut.redisUtil.GetClient(), sut.ctx, sut.sessionId)
	sut.Equal(err, redis.Nil)
}

func (sut *LogoutServiceTestSuite) AfterTest(suiteName, testName string) {
	sut.T().Log("AfterTest: " + suiteName + " " + testName)
}

func (sut *LogoutServiceTestSuite) TearDownTest() {
	sut.T().Log("TearDownTest")
	fmt.Println()
}

func (sut *LogoutServiceTestSuite) TearDownSuite() {
	sut.T().Log("TearDownSuite")
}
