package services_test

import (
	"context"
	"errors"
	"fmt"
	"golang-postgres/commons/helpers"
	"golang-postgres/features/login/models"
	"golang-postgres/features/login/services"
	mockhelpers "golang-postgres/tests/unit_tests/commons/helpers/mocks"
	mockutils "golang-postgres/tests/unit_tests/commons/utils/mocks"
	mockrepositories "golang-postgres/tests/unit_tests/features/login/mocks/repositories"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/suite"
)

type LoginServiceTestSuite struct {
	suite.Suite
	ctx                          context.Context
	requestId                    string
	sessionId                    string
	pool                         *pgxpool.Pool
	client                       *redis.Client
	errTimeout                   error
	errInternalServer            error
	loginUserRequest             models.LoginUserRequest
	user                         models.User
	postgresUtilMock             *mockutils.PostgresUtilMock
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
	sut.pool = &pgxpool.Pool{}
	sut.client = &redis.Client{}
	sut.errTimeout = context.Canceled
	sut.errInternalServer = errors.New("internal server error")
}

func (sut *LoginServiceTestSuite) SetupTest() {
	fmt.Println()
	sut.T().Log("SetupTest")
	sut.loginUserRequest = models.LoginUserRequest{
		Email:    "email17@email.com",
		Password: "password@A1",
	}
	sut.user = models.User{
		Id:        pgtype.Int4{Valid: true, Int32: 1},
		Username:  pgtype.Text{Valid: true, String: "username17"},
		Email:     pgtype.Text{Valid: true, String: "email17@email.com"},
		Password:  pgtype.Text{Valid: true, String: "$2a$10$MvEM5qcQFk39jC/3fYzJzOIy7M/xQiGv/PAkkoarCMgsx/rO0UaPG"},
		Utc:       pgtype.Text{Valid: true, String: "+0800"},
		CreatedAt: pgtype.Int8{Valid: true, Int64: 1719496855216},
	}
	sut.postgresUtilMock = new(mockutils.PostgresUtilMock)
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
	sut.loginService = services.NewLoginService(sut.postgresUtilMock, sut.redisUtilMock, sut.validate, sut.userRepositoryMock, sut.userPermissionRepositoryMock, sut.bcryptHelperMock, sut.redisRepositoryMock, sut.uuidHelperMock)
}

func (sut *LoginServiceTestSuite) BeforeTest(suiteName, testName string) {
	sut.T().Log("BeforeTest: " + suiteName + " " + testName)
}

func (sut *LoginServiceTestSuite) Test01LoginRedisRepositoryDelWithSessionIdTimeoutError() {
	sut.T().Log("Test01LoginRedisRepositoryDelWithSessionIdTimeoutError")
	sut.loginUserRequest = models.LoginUserRequest{}
	sut.redisUtilMock.Mock.On("GetClient").Return(sut.client)
	sut.redisRepositoryMock.Mock.On("Del", sut.client, sut.ctx, sut.sessionId).Return(int64(0), sut.errTimeout)
	response, err := sut.loginService.Login(sut.ctx, sut.requestId, sut.sessionId, sut.loginUserRequest)
	sut.Equal(response, "")
	sut.Equal(err.Error(), "validation error")
}

func (sut *LoginServiceTestSuite) Test02LoginRedisRepositoryDelWithSessionIdInternalServerError() {
	sut.T().Log("Test02LoginRedisRepositoryDelWithSessionIdInternalServerError")
	sut.loginUserRequest = models.LoginUserRequest{}
	sut.redisUtilMock.Mock.On("GetClient").Return(sut.client)
	sut.redisRepositoryMock.Mock.On("Del", sut.client, sut.ctx, sut.sessionId).Return(int64(0), sut.errInternalServer)
	response, err := sut.loginService.Login(sut.ctx, sut.requestId, sut.sessionId, sut.loginUserRequest)
	sut.Equal(response, "")
	sut.Equal(err.Error(), "validation error")
}

func (sut *LoginServiceTestSuite) Test03LoginRedisRepositoryDelWithoutSessionIdTimeoutError() {
	sut.T().Log("Test03LoginRedisRepositoryDelWithoutSessionIdTimeoutError")
	sut.loginUserRequest = models.LoginUserRequest{}
	sut.redisUtilMock.Mock.On("GetClient").Return(sut.client)
	sut.redisRepositoryMock.Mock.On("Del", sut.client, sut.ctx, sut.sessionId).Return(int64(0), sut.errTimeout)
	response, err := sut.loginService.Login(sut.ctx, sut.requestId, "", sut.loginUserRequest)
	sut.Equal(response, "")
	sut.Equal(err.Error(), "validation error")
}

func (sut *LoginServiceTestSuite) Test04LoginRedisRepositoryDelWithoutSessionIdInternalServerError() {
	sut.T().Log("Test04LoginRedisRepositoryDelWithoutSessionIdInternalServerError")
	sut.loginUserRequest = models.LoginUserRequest{}
	sut.redisUtilMock.Mock.On("GetClient").Return(sut.client)
	sut.redisRepositoryMock.Mock.On("Del", sut.client, sut.ctx, sut.sessionId).Return(int64(0), sut.errInternalServer)
	response, err := sut.loginService.Login(sut.ctx, sut.requestId, "", sut.loginUserRequest)
	sut.Equal(response, "")
	sut.Equal(err.Error(), "validation error")
}

func (sut *LoginServiceTestSuite) Test05LoginUserRepositoryFindByEmailTimeoutError() {
	sut.T().Log("Test05LoginUserRepositoryFindByEmailTimeoutError")

	sut.redisUtilMock.Mock.On("GetClient").Return(sut.client)
	sut.redisRepositoryMock.Mock.On("Del", sut.client, sut.ctx, sut.sessionId).Return(int64(0), sut.errInternalServer)
	sut.postgresUtilMock.Mock.On("GetPool").Return(sut.pool)
	sut.userRepositoryMock.Mock.On("FindByEmail", sut.pool, sut.ctx, sut.loginUserRequest.Email).Return(models.User{}, sut.errTimeout)
	response, err := sut.loginService.Login(sut.ctx, sut.requestId, "", sut.loginUserRequest)
	sut.Equal(response, "")
	sut.Equal(err.Error(), "time out or user cancel the request")
}

func (sut *LoginServiceTestSuite) Test06LoginUserRepositoryFindByEmailInternalServerError() {
	sut.T().Log("Test06LoginUserRepositoryFindByEmailInternalServerError")
	sut.redisUtilMock.Mock.On("GetClient").Return(sut.client)
	sut.redisRepositoryMock.Mock.On("Del", sut.client, sut.ctx, sut.sessionId).Return(int64(0), sut.errInternalServer)
	sut.postgresUtilMock.Mock.On("GetPool").Return(sut.pool)
	sut.userRepositoryMock.Mock.On("FindByEmail", sut.pool, sut.ctx, sut.loginUserRequest.Email).Return(models.User{}, sut.errInternalServer)
	response, err := sut.loginService.Login(sut.ctx, sut.requestId, "", sut.loginUserRequest)
	sut.Equal(response, "")
	sut.Equal(err.Error(), "internal server error")
}

func (sut *LoginServiceTestSuite) Test07LoginUserRepositoryFindByEmailBadRequestWrongEmailPassword() {
	sut.T().Log("Test07LoginUserRepositoryFindByEmailBadRequestWrongEmailPassword")
	sut.redisUtilMock.Mock.On("GetClient").Return(sut.client)
	sut.redisRepositoryMock.Mock.On("Del", sut.client, sut.ctx, sut.sessionId).Return(int64(0), sut.errInternalServer)
	sut.postgresUtilMock.Mock.On("GetPool").Return(sut.pool)
	sut.userRepositoryMock.Mock.On("FindByEmail", sut.pool, sut.ctx, sut.loginUserRequest.Email).Return(models.User{}, pgx.ErrNoRows)
	response, err := sut.loginService.Login(sut.ctx, sut.requestId, "", sut.loginUserRequest)
	sut.Equal(response, "")
	sut.Equal(err.Error(), "wrong email or password")
}

func (sut *LoginServiceTestSuite) Test08LoginBcryptCompareHashAndPasswordBadRequestWrongEmailPassword() {
	sut.T().Log("Test08LoginBcryptCompareHashAndPasswordBadRequestWrongEmailPassword")
	sut.redisUtilMock.Mock.On("GetClient").Return(sut.client)
	sut.redisRepositoryMock.Mock.On("Del", sut.client, sut.ctx, sut.sessionId).Return(int64(0), sut.errInternalServer)
	sut.postgresUtilMock.Mock.On("GetPool").Return(sut.pool)
	sut.user.Password.String = ""
	sut.userRepositoryMock.Mock.On("FindByEmail", sut.pool, sut.ctx, sut.loginUserRequest.Email).Return(sut.user, nil)
	response, err := sut.loginService.Login(sut.ctx, sut.requestId, "", sut.loginUserRequest)
	sut.Equal(response, "")
	sut.Equal(err.Error(), "wrong email or password")
}

func (sut *LoginServiceTestSuite) Test09LoginUserPermissionRepositoryFindByUserIdTimeoutError() {
	sut.T().Log("Test09LoginUserPermissionRepositoryFindByUserIdTimeoutError")
	sut.redisUtilMock.Mock.On("GetClient").Return(sut.client)
	sut.redisRepositoryMock.Mock.On("Del", sut.client, sut.ctx, sut.sessionId).Return(int64(0), sut.errInternalServer)
	sut.postgresUtilMock.Mock.On("GetPool").Return(sut.pool)
	sut.userRepositoryMock.Mock.On("FindByEmail", sut.pool, sut.ctx, sut.loginUserRequest.Email).Return(sut.user, nil)
	sut.userPermissionRepositoryMock.Mock.On("FindByUserId", sut.pool, sut.ctx, sut.user.Id.Int32).Return([]models.UserPermission{}, sut.errTimeout)
	response, err := sut.loginService.Login(sut.ctx, sut.requestId, "", sut.loginUserRequest)
	sut.Equal(response, "")
	sut.Equal(err.Error(), "time out or user cancel the request")
}

func (sut *LoginServiceTestSuite) Test10LoginUserPermissionRepositoryFindByUserIdInternalServerError() {
	sut.T().Log("Test10LoginUserPermissionRepositoryFindByUserIdInternalServerError")
	sut.redisUtilMock.Mock.On("GetClient").Return(sut.client)
	sut.redisRepositoryMock.Mock.On("Del", sut.client, sut.ctx, sut.sessionId).Return(int64(0), sut.errInternalServer)
	sut.postgresUtilMock.Mock.On("GetPool").Return(sut.pool)
	sut.userRepositoryMock.Mock.On("FindByEmail", sut.pool, sut.ctx, sut.loginUserRequest.Email).Return(sut.user, nil)
	sut.userPermissionRepositoryMock.Mock.On("FindByUserId", sut.pool, sut.ctx, sut.user.Id.Int32).Return([]models.UserPermission{}, sut.errInternalServer)
	response, err := sut.loginService.Login(sut.ctx, sut.requestId, "", sut.loginUserRequest)
	sut.Equal(response, "")
	sut.Equal(err.Error(), "internal server error")
}

func (sut *LoginServiceTestSuite) Test11LoginRedisRepositorySetTimeoutError() {
	sut.T().Log("Test11LoginRedisRepositorySetTimeoutError")
	sut.redisUtilMock.Mock.On("GetClient").Return(sut.client)
	sut.redisRepositoryMock.Mock.On("Del", sut.client, sut.ctx, sut.sessionId).Return(int64(0), sut.errInternalServer)
	sut.postgresUtilMock.Mock.On("GetPool").Return(sut.pool)
	sut.userRepositoryMock.Mock.On("FindByEmail", sut.pool, sut.ctx, sut.loginUserRequest.Email).Return(sut.user, nil)
	sut.userPermissionRepositoryMock.Mock.On("FindByUserId", sut.pool, sut.ctx, sut.user.Id.Int32).Return([]models.UserPermission{}, nil)
	sut.uuidHelperMock.Mock.On("String").Return(sut.sessionId)
	session := `{"email":"email17@email.com","id":1,"idPermissions":null,"username":"username17"}`
	sut.redisRepositoryMock.Mock.On("Set", sut.client, sut.ctx, sut.sessionId, session, time.Duration(0)).Return("", sut.errTimeout)
	response, err := sut.loginService.Login(sut.ctx, sut.requestId, "", sut.loginUserRequest)
	sut.Equal(response, sut.sessionId)
	sut.Equal(err.Error(), "time out or user cancel the request")
}

func (sut *LoginServiceTestSuite) Test12LoginRedisRepositorySetInternalServerError() {
	sut.T().Log("Test12LoginRedisRepositorySetInternalServerError")
	sut.redisUtilMock.Mock.On("GetClient").Return(sut.client)
	sut.redisRepositoryMock.Mock.On("Del", sut.client, sut.ctx, sut.sessionId).Return(int64(0), sut.errInternalServer)
	sut.postgresUtilMock.Mock.On("GetPool").Return(sut.pool)
	sut.userRepositoryMock.Mock.On("FindByEmail", sut.pool, sut.ctx, sut.loginUserRequest.Email).Return(sut.user, nil)
	sut.userPermissionRepositoryMock.Mock.On("FindByUserId", sut.pool, sut.ctx, sut.user.Id.Int32).Return([]models.UserPermission{}, nil)
	sut.uuidHelperMock.Mock.On("String").Return(sut.sessionId)
	session := `{"email":"email17@email.com","id":1,"idPermissions":null,"username":"username17"}`
	sut.redisRepositoryMock.Mock.On("Set", sut.client, sut.ctx, sut.sessionId, session, time.Duration(0)).Return("", sut.errInternalServer)
	response, err := sut.loginService.Login(sut.ctx, sut.requestId, "", sut.loginUserRequest)
	sut.Equal(response, sut.sessionId)
	sut.Equal(err.Error(), "internal server error")
}

func (sut *LoginServiceTestSuite) Test13LoginSuccess() {
	sut.T().Log("Test13LoginSuccess")
	sut.redisUtilMock.Mock.On("GetClient").Return(sut.client)
	sut.redisRepositoryMock.Mock.On("Del", sut.client, sut.ctx, sut.sessionId).Return(int64(0), sut.errInternalServer)
	sut.postgresUtilMock.Mock.On("GetPool").Return(sut.pool)
	sut.userRepositoryMock.Mock.On("FindByEmail", sut.pool, sut.ctx, sut.loginUserRequest.Email).Return(sut.user, nil)
	sut.userPermissionRepositoryMock.Mock.On("FindByUserId", sut.pool, sut.ctx, sut.user.Id.Int32).Return([]models.UserPermission{}, nil)
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
	fmt.Println()
}

func (sut *LoginServiceTestSuite) TearDownSuite() {
	sut.T().Log("TearDownSuite")
}
