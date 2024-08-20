package services_test

import (
	"context"
	"errors"
	"fmt"
	"golang-postgres/features/logout/services"
	mockutils "golang-postgres/tests/unit_tests/commons/utils/mocks"
	mockrepositories "golang-postgres/tests/unit_tests/features/logout/mocks/repositories"
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/suite"
)

type LogoutServiceTestSuite struct {
	suite.Suite
	ctx                 context.Context
	requestId           string
	sessionId           string
	client              *redis.Client
	errTimeout          error
	errInternalServer   error
	redisUtilMock       *mockutils.RedisUtilMock
	redisRepositoryMock *mockrepositories.RedisRepositoryMock
	logoutService       services.LogoutService
}

func TestLogoutTestSuite(t *testing.T) {
	suite.Run(t, new(LogoutServiceTestSuite))
}

func (sut *LogoutServiceTestSuite) SetupSuite() {
	sut.T().Log("SetupSuite")
	sut.ctx = context.Background()
	sut.requestId = "requestId"
	sut.sessionId = "sessionId"
	sut.client = &redis.Client{}
	sut.errTimeout = context.Canceled
	sut.errInternalServer = errors.New("internal server error")
}

func (sut *LogoutServiceTestSuite) SetupTest() {
	fmt.Println()
	sut.T().Log("SetupTest")
	sut.redisUtilMock = new(mockutils.RedisUtilMock)
	sut.redisRepositoryMock = new(mockrepositories.RedisRepositoryMock)
	sut.logoutService = services.NewLogoutService(sut.redisUtilMock, sut.redisRepositoryMock)
}

func (sut *LogoutServiceTestSuite) BeforeTest(suiteName, testName string) {
	sut.T().Log("BeforeTest: " + suiteName + " " + testName)
}

func (sut *LogoutServiceTestSuite) Test1LogoutRedisRepositoryDelTimeoutError() {
	sut.T().Log("Test1LogoutRedisRepositoryDelTimeoutError")
	sut.redisUtilMock.Mock.On("GetClient").Return(sut.client)
	sut.redisRepositoryMock.Mock.On("Del", sut.client, sut.ctx, sut.sessionId).Return(int64(0), sut.errTimeout)
	err := sut.logoutService.Logout(sut.ctx, sut.requestId, sut.sessionId)
	sut.Equal(err.Error(), "time out or user cancel the request")
}

func (sut *LogoutServiceTestSuite) Test2LogoutRedisRepositoryDelInternalServerError() {
	sut.T().Log("Test2LogoutRedisRepositoryDelInternalServerError")
	sut.redisUtilMock.Mock.On("GetClient").Return(sut.client)
	sut.redisRepositoryMock.Mock.On("Del", sut.client, sut.ctx, sut.sessionId).Return(int64(0), sut.errInternalServer)
	err := sut.logoutService.Logout(sut.ctx, sut.requestId, sut.sessionId)
	sut.Equal(err.Error(), "internal server error")
}

func (sut *LogoutServiceTestSuite) Test3LogoutRedisRepositoryDelRowsAffectedNotOne() {
	sut.T().Log("Test3LogoutRedisRepositoryDelRowsAffectedNotOne")
	sut.redisUtilMock.Mock.On("GetClient").Return(sut.client)
	sut.redisRepositoryMock.Mock.On("Del", sut.client, sut.ctx, sut.sessionId).Return(int64(0), nil)
	err := sut.logoutService.Logout(sut.ctx, sut.requestId, sut.sessionId)
	sut.Equal(err.Error(), "internal server error")
}

func (sut *LogoutServiceTestSuite) Test4LogoutSuccess() {
	sut.T().Log("Test4LogoutSuccess")
	sut.redisUtilMock.Mock.On("GetClient").Return(sut.client)
	sut.redisRepositoryMock.Mock.On("Del", sut.client, sut.ctx, sut.sessionId).Return(int64(1), nil)
	err := sut.logoutService.Logout(sut.ctx, sut.requestId, sut.sessionId)
	sut.Equal(err, nil)
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
