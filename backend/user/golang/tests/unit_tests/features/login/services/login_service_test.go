package services_test

import (
	"context"
	"database/sql"
	"errors"
	"project-user/commons/helpers"
	"project-user/features/login/models"
	"project-user/features/login/services"
	"testing"
	"time"

	mockhelpers "project-user/tests/unit_tests/mocks/helpers"
	mockutils "project-user/tests/unit_tests/mocks/utils"

	mockrepositories "project-user/tests/unit_tests/features/login/mocks/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/suite"
)

type LoginServiceTestSuite struct {
	suite.Suite
	ctx                          context.Context
	requestId                    string
	sessionId                    string
	db                           *sql.DB
	client                       *redis.Client
	errTimeout                   error
	errInternalServer            error
	loginUserRequest             models.LoginUserRequest
	user                         models.User
	mysqlUtilMock                *mockutils.MysqlUtilMock
	redisUtilMock                *mockutils.RedisUtilMock
	validate                     *validator.Validate
	userRepositoryMock           *mockrepositories.UserRepositoryMock
	userPermissionRepositoryMock *mockrepositories.UserPermissionRepositoryMock
	bcryptHelperMock             *mockhelpers.BcryptHelperMock
	redisRepositoryMock          *mockrepositories.RedisRepositoryMock
	uuidHelperMock               *mockhelpers.UuidHelperMock
	loginService                 services.LoginService
}

func TestLoginTestSuite(t *testing.T) {
	suite.Run(t, new(LoginServiceTestSuite))
}

func (sut *LoginServiceTestSuite) SetupSuite() {
	sut.T().Log("SetupSuite")
	sut.ctx = context.Background()
	sut.requestId = "requestId"
	sut.sessionId = "sessionId"
	sut.db = &sql.DB{}
	sut.client = &redis.Client{}
	sut.errTimeout = context.Canceled
	sut.errInternalServer = errors.New("internal server error")
}

func (sut *LoginServiceTestSuite) SetupTest() {
	sut.T().Log("SetupTest")
	sut.loginUserRequest = models.LoginUserRequest{
		Email:    "email17@email.com",
		Password: "password@A1",
	}
	sut.user = models.User{
		Id:        sql.NullInt32{Valid: true, Int32: 1},
		Username:  sql.NullString{Valid: true, String: "username17"},
		Email:     sql.NullString{Valid: true, String: "email17@email.com"},
		Password:  sql.NullString{Valid: true, String: "$2a$10$MvEM5qcQFk39jC/3fYzJzOIy7M/xQiGv/PAkkoarCMgsx/rO0UaPG"},
		Utc:       sql.NullString{Valid: true, String: "+0800"},
		CreatedAt: sql.NullInt64{Valid: true, Int64: 1719496855216},
	}
	sut.mysqlUtilMock = new(mockutils.MysqlUtilMock)
	sut.redisUtilMock = new(mockutils.RedisUtilMock)
	sut.validate = validator.New()
	helpers.UsernameValidator(sut.validate)
	helpers.PasswordValidator(sut.validate)
	helpers.TelephoneValidator(sut.validate)
	sut.userRepositoryMock = new(mockrepositories.UserRepositoryMock)
	sut.userPermissionRepositoryMock = new(mockrepositories.UserPermissionRepositoryMock)
	sut.bcryptHelperMock = new(mockhelpers.BcryptHelperMock)
	sut.redisRepositoryMock = new(mockrepositories.RedisRepositoryMock)
	sut.uuidHelperMock = new(mockhelpers.UuidHelperMock)
	sut.loginService = services.NewLoginService(sut.mysqlUtilMock, sut.redisUtilMock, sut.validate, sut.userRepositoryMock, sut.userPermissionRepositoryMock, sut.bcryptHelperMock, sut.redisRepositoryMock, sut.uuidHelperMock)
}

func (sut *LoginServiceTestSuite) BeforeTest(suiteName, testName string) {
	sut.T().Log("BeforeTest: " + suiteName + " " + testName)
}

func (sut *LoginServiceTestSuite) TestLoginRedisRepositoryDelWithSessionIdTimeoutError() {
	sut.T().Log("TestLoginRedisRepositoryDelWithSessionIdTimeoutError")
	sut.loginUserRequest = models.LoginUserRequest{}
	sut.redisUtilMock.Mock.On("GetClient").Return(sut.client)
	sut.redisRepositoryMock.Mock.On("Del", sut.client, sut.ctx, sut.sessionId).Return(int64(0), sut.errTimeout)
	response, err := sut.loginService.Login(sut.ctx, sut.requestId, sut.sessionId, sut.loginUserRequest)
	sut.Equal(response, "")
	sut.Equal(err.Error(), "validation error")
}

func (sut *LoginServiceTestSuite) TestLoginRedisRepositoryDelWithSessionIdInternalServerError() {
	sut.T().Log("TestLoginRedisRepositoryDelWithSessionIdInternalServerError")
	sut.loginUserRequest = models.LoginUserRequest{}
	sut.redisUtilMock.Mock.On("GetClient").Return(sut.client)
	sut.redisRepositoryMock.Mock.On("Del", sut.client, sut.ctx, sut.sessionId).Return(int64(0), sut.errInternalServer)
	response, err := sut.loginService.Login(sut.ctx, sut.requestId, sut.sessionId, sut.loginUserRequest)
	sut.Equal(response, "")
	sut.Equal(err.Error(), "validation error")
}

func (sut *LoginServiceTestSuite) TestLoginRedisRepositoryDelWithoutSessionIdTimeoutError() {
	sut.T().Log("TestLoginRedisRepositoryDelWithoutSessionIdTimeoutError")
	sut.loginUserRequest = models.LoginUserRequest{}
	sut.redisUtilMock.Mock.On("GetClient").Return(sut.client)
	sut.redisRepositoryMock.Mock.On("Del", sut.client, sut.ctx, sut.sessionId).Return(int64(0), sut.errTimeout)
	response, err := sut.loginService.Login(sut.ctx, sut.requestId, "", sut.loginUserRequest)
	sut.Equal(response, "")
	sut.Equal(err.Error(), "validation error")
}

func (sut *LoginServiceTestSuite) TestLoginRedisRepositoryDelWithoutSessionIdInternalServerError() {
	sut.T().Log("TestLoginRedisRepositoryDelWithoutSessionIdInternalServerError")
	sut.loginUserRequest = models.LoginUserRequest{}
	sut.redisUtilMock.Mock.On("GetClient").Return(sut.client)
	sut.redisRepositoryMock.Mock.On("Del", sut.client, sut.ctx, sut.sessionId).Return(int64(0), sut.errInternalServer)
	response, err := sut.loginService.Login(sut.ctx, sut.requestId, "", sut.loginUserRequest)
	sut.Equal(response, "")
	sut.Equal(err.Error(), "validation error")
}

func (sut *LoginServiceTestSuite) TestLoginUserRepositoryFindByEmailTimeoutError() {
	sut.T().Log("TestLoginUserRepositoryFindByEmailTimeoutError")

	sut.redisUtilMock.Mock.On("GetClient").Return(sut.client)
	sut.redisRepositoryMock.Mock.On("Del", sut.client, sut.ctx, sut.sessionId).Return(int64(0), sut.errInternalServer)
	sut.mysqlUtilMock.Mock.On("GetDb").Return(sut.db)
	sut.userRepositoryMock.Mock.On("FindByEmail", sut.db, sut.ctx, sut.loginUserRequest.Email).Return(models.User{}, sut.errTimeout)
	response, err := sut.loginService.Login(sut.ctx, sut.requestId, "", sut.loginUserRequest)
	sut.Equal(response, "")
	sut.Equal(err.Error(), "time out or user cancel the request")
}

func (sut *LoginServiceTestSuite) TestLoginUserRepositoryFindByEmailInternalServerError() {
	sut.T().Log("TestLoginUserRepositoryFindByEmailInternalServerError")
	sut.redisUtilMock.Mock.On("GetClient").Return(sut.client)
	sut.redisRepositoryMock.Mock.On("Del", sut.client, sut.ctx, sut.sessionId).Return(int64(0), sut.errInternalServer)
	sut.mysqlUtilMock.Mock.On("GetDb").Return(sut.db)
	sut.userRepositoryMock.Mock.On("FindByEmail", sut.db, sut.ctx, sut.loginUserRequest.Email).Return(models.User{}, sut.errInternalServer)
	response, err := sut.loginService.Login(sut.ctx, sut.requestId, "", sut.loginUserRequest)
	sut.Equal(response, "")
	sut.Equal(err.Error(), "internal server error")
}

func (sut *LoginServiceTestSuite) TestLoginUserRepositoryFindByEmailBadRequestWrongEmailPassword() {
	sut.T().Log("TestLoginUserRepositoryFindByEmailBadRequestWrongEmailPassword")
	sut.redisUtilMock.Mock.On("GetClient").Return(sut.client)
	sut.redisRepositoryMock.Mock.On("Del", sut.client, sut.ctx, sut.sessionId).Return(int64(0), sut.errInternalServer)
	sut.mysqlUtilMock.Mock.On("GetDb").Return(sut.db)
	sut.userRepositoryMock.Mock.On("FindByEmail", sut.db, sut.ctx, sut.loginUserRequest.Email).Return(models.User{}, sql.ErrNoRows)
	response, err := sut.loginService.Login(sut.ctx, sut.requestId, "", sut.loginUserRequest)
	sut.Equal(response, "")
	sut.Equal(err.Error(), "wrong email or password")
}

func (sut *LoginServiceTestSuite) TestLoginBcryptCompareHashAndPasswordBadRequestWrongEmailPassword() {
	sut.T().Log("TestLoginBcryptCompareHashAndPasswordBadRequestWrongEmailPassword")
	sut.redisUtilMock.Mock.On("GetClient").Return(sut.client)
	sut.redisRepositoryMock.Mock.On("Del", sut.client, sut.ctx, sut.sessionId).Return(int64(0), sut.errInternalServer)
	sut.mysqlUtilMock.Mock.On("GetDb").Return(sut.db)
	sut.user.Password.String = ""
	sut.userRepositoryMock.Mock.On("FindByEmail", sut.db, sut.ctx, sut.loginUserRequest.Email).Return(sut.user, nil)
	response, err := sut.loginService.Login(sut.ctx, sut.requestId, "", sut.loginUserRequest)
	sut.Equal(response, "")
	sut.Equal(err.Error(), "wrong email or password")
}

func (sut *LoginServiceTestSuite) TestLoginUserPermissionRepositoryFindByUserIdTimeoutError() {
	sut.T().Log("TestLoginUserPermissionRepositoryFindByUserIdTimeoutError")
	sut.redisUtilMock.Mock.On("GetClient").Return(sut.client)
	sut.redisRepositoryMock.Mock.On("Del", sut.client, sut.ctx, sut.sessionId).Return(int64(0), sut.errInternalServer)
	sut.mysqlUtilMock.Mock.On("GetDb").Return(sut.db)
	sut.userRepositoryMock.Mock.On("FindByEmail", sut.db, sut.ctx, sut.loginUserRequest.Email).Return(sut.user, nil)
	sut.userPermissionRepositoryMock.Mock.On("FindByUserId", sut.db, sut.ctx, sut.user.Id.Int32).Return([]models.UserPermission{}, sut.errTimeout)
	response, err := sut.loginService.Login(sut.ctx, sut.requestId, "", sut.loginUserRequest)
	sut.Equal(response, "")
	sut.Equal(err.Error(), "time out or user cancel the request")
}

func (sut *LoginServiceTestSuite) TestLoginUserPermissionRepositoryFindByUserIdInternalServerError() {
	sut.T().Log("TestLoginUserPermissionRepositoryFindByUserIdInternalServerError")
	sut.redisUtilMock.Mock.On("GetClient").Return(sut.client)
	sut.redisRepositoryMock.Mock.On("Del", sut.client, sut.ctx, sut.sessionId).Return(int64(0), sut.errInternalServer)
	sut.mysqlUtilMock.Mock.On("GetDb").Return(sut.db)
	sut.userRepositoryMock.Mock.On("FindByEmail", sut.db, sut.ctx, sut.loginUserRequest.Email).Return(sut.user, nil)
	sut.userPermissionRepositoryMock.Mock.On("FindByUserId", sut.db, sut.ctx, sut.user.Id.Int32).Return([]models.UserPermission{}, sut.errInternalServer)
	response, err := sut.loginService.Login(sut.ctx, sut.requestId, "", sut.loginUserRequest)
	sut.Equal(response, "")
	sut.Equal(err.Error(), "internal server error")
}

func (sut *LoginServiceTestSuite) TestLoginRedisRepositorySetTimeoutError() {
	sut.T().Log("TestLoginRedisRepositorySetTimeoutError")
	sut.redisUtilMock.Mock.On("GetClient").Return(sut.client)
	sut.redisRepositoryMock.Mock.On("Del", sut.client, sut.ctx, sut.sessionId).Return(int64(0), sut.errInternalServer)
	sut.mysqlUtilMock.Mock.On("GetDb").Return(sut.db)
	sut.userRepositoryMock.Mock.On("FindByEmail", sut.db, sut.ctx, sut.loginUserRequest.Email).Return(sut.user, nil)
	sut.userPermissionRepositoryMock.Mock.On("FindByUserId", sut.db, sut.ctx, sut.user.Id.Int32).Return([]models.UserPermission{}, nil)
	sut.uuidHelperMock.Mock.On("String").Return(sut.sessionId)
	session := `{"email":"email17@email.com","id":1,"idPermissions":null,"username":"username17"}`
	sut.redisRepositoryMock.Mock.On("Set", sut.client, sut.ctx, sut.sessionId, session, time.Duration(0)).Return("", sut.errTimeout)
	response, err := sut.loginService.Login(sut.ctx, sut.requestId, "", sut.loginUserRequest)
	sut.Equal(response, sut.sessionId)
	sut.Equal(err.Error(), "time out or user cancel the request")
}

func (sut *LoginServiceTestSuite) TestLoginRedisRepositorySetInternalServerError() {
	sut.T().Log("TestLoginRedisRepositorySetInternalServerError")
	sut.redisUtilMock.Mock.On("GetClient").Return(sut.client)
	sut.redisRepositoryMock.Mock.On("Del", sut.client, sut.ctx, sut.sessionId).Return(int64(0), sut.errInternalServer)
	sut.mysqlUtilMock.Mock.On("GetDb").Return(sut.db)
	sut.userRepositoryMock.Mock.On("FindByEmail", sut.db, sut.ctx, sut.loginUserRequest.Email).Return(sut.user, nil)
	sut.userPermissionRepositoryMock.Mock.On("FindByUserId", sut.db, sut.ctx, sut.user.Id.Int32).Return([]models.UserPermission{}, nil)
	sut.uuidHelperMock.Mock.On("String").Return(sut.sessionId)
	session := `{"email":"email17@email.com","id":1,"idPermissions":null,"username":"username17"}`
	sut.redisRepositoryMock.Mock.On("Set", sut.client, sut.ctx, sut.sessionId, session, time.Duration(0)).Return("", sut.errInternalServer)
	response, err := sut.loginService.Login(sut.ctx, sut.requestId, "", sut.loginUserRequest)
	sut.Equal(response, sut.sessionId)
	sut.Equal(err.Error(), "internal server error")
}

func (sut *LoginServiceTestSuite) TestLoginSuccess() {
	sut.T().Log("TestLoginSuccess")
	sut.redisUtilMock.Mock.On("GetClient").Return(sut.client)
	sut.redisRepositoryMock.Mock.On("Del", sut.client, sut.ctx, sut.sessionId).Return(int64(0), sut.errInternalServer)
	sut.mysqlUtilMock.Mock.On("GetDb").Return(sut.db)
	sut.userRepositoryMock.Mock.On("FindByEmail", sut.db, sut.ctx, sut.loginUserRequest.Email).Return(sut.user, nil)
	sut.userPermissionRepositoryMock.Mock.On("FindByUserId", sut.db, sut.ctx, sut.user.Id.Int32).Return([]models.UserPermission{}, nil)
	sut.uuidHelperMock.Mock.On("String").Return(sut.sessionId)
	session := `{"email":"email17@email.com","id":1,"idPermissions":null,"username":"username17"}`
	sut.redisRepositoryMock.Mock.On("Set", sut.client, sut.ctx, sut.sessionId, session, time.Duration(0)).Return("", nil)
	response, err := sut.loginService.Login(sut.ctx, sut.requestId, "", sut.loginUserRequest)
	sut.Equal(response, sut.sessionId)
	sut.Equal(err, nil)
}

func (sut *LoginServiceTestSuite) AfterTest(suiteName, testName string) {
	sut.T().Log("AfterTest: " + suiteName + " " + testName)
}

func (sut *LoginServiceTestSuite) TearDownTest() {
	sut.T().Log("TearDownTest")
}

func (sut *LoginServiceTestSuite) TearDownSuite() {
	sut.T().Log("TearDownSuite")
}
