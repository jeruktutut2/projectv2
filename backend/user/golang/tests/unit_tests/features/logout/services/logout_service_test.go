package services_test

import (
	"context"
	"errors"
	"testing"

	"project-user/features/logout/services"
	mockrepositories "project-user/tests/unit_tests/features/logout/mocks/repositories"
	mockutils "project-user/tests/unit_tests/mocks/utils"

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
	// sut.errTimeout = context.Canceled
	sut.client = &redis.Client{}
	sut.errTimeout = context.Canceled
	sut.errInternalServer = errors.New("internal server error")
}

func (sut *LogoutServiceTestSuite) SetupTest() {
	sut.T().Log("SetupTest")
	sut.redisUtilMock = new(mockutils.RedisUtilMock)
	sut.redisRepositoryMock = new(mockrepositories.RedisRepositoryMock)
	sut.logoutService = services.NewLogoutService(sut.redisUtilMock, sut.redisRepositoryMock)
}

func (sut *LogoutServiceTestSuite) BeforeTest(suiteName, testName string) {
	sut.T().Log("BeforeTest: " + suiteName + " " + testName)
}

func (sut *LogoutServiceTestSuite) TestLogoutRedisRepositoryDelTimeoutError() {
	sut.T().Log("TestLogoutRedisRepositoryDelTimeoutError")
	sut.redisUtilMock.Mock.On("GetClient").Return(sut.client)
	sut.redisRepositoryMock.Mock.On("Del", sut.client, sut.ctx, sut.sessionId).Return(int64(0), sut.errTimeout)
	err := sut.logoutService.Logout(sut.ctx, sut.requestId, sut.sessionId)
	// sut.Equal(err.Error(), `[{"field":"message","message":"time out or user cancel the request"}]`)
	sut.Equal(err.Error(), "time out or user cancel the request")
}

func (sut *LogoutServiceTestSuite) TestLogoutRedisRepositoryDelInternalServerError() {
	sut.T().Log("TestLogoutRedisRepositoryDelInternalServerError")
	sut.redisUtilMock.Mock.On("GetClient").Return(sut.client)
	sut.redisRepositoryMock.Mock.On("Del", sut.client, sut.ctx, sut.sessionId).Return(int64(0), sut.errInternalServer)
	err := sut.logoutService.Logout(sut.ctx, sut.requestId, sut.sessionId)
	// sut.Equal(err.Error(), `[{"field":"message","message":"internal server error"}]`)
	sut.Equal(err.Error(), "internal server error")
}

func (sut *LogoutServiceTestSuite) TestLogoutRedisRepositoryDelRowsAffectedNotOne() {
	sut.T().Log("TestLogoutRedisRepositoryDelRowsAffectedNotOne")
	sut.redisUtilMock.Mock.On("GetClient").Return(sut.client)
	sut.redisRepositoryMock.Mock.On("Del", sut.client, sut.ctx, sut.sessionId).Return(int64(0), nil)
	err := sut.logoutService.Logout(sut.ctx, sut.requestId, sut.sessionId)
	// sut.Equal(err.Error(), `[{"field":"message","message":"internal server error"}]`)
	sut.Equal(err.Error(), "internal server error")
}

func (sut *LogoutServiceTestSuite) TestLogoutSuccess() {
	sut.T().Log("TestLogoutSuccess")
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
}

func (sut *LogoutServiceTestSuite) TearDownSuite() {
	sut.T().Log("TearDownSuite")
}
