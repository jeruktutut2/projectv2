package services_test

import (
	"context"
	"testing"

	"project-user/features/login/models"
	"project-user/features/login/repositories"
	"project-user/features/login/services"
	"project-user/helpers"
	"project-user/tests/initialize"
	"project-user/utils"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/suite"
)

type LoginServiceTestSuite struct {
	suite.Suite
	ctx                      context.Context
	requestId                string
	sessionId                string
	loginUserRequest         models.LoginUserRequest
	mysqlUtil                utils.MysqlUtil
	redisUtil                utils.RedisUtil
	validate                 *validator.Validate
	userRepository           repositories.UserRepository
	userPermissionRepository repositories.UserPermissionRepository
	bcryptHelper             helpers.BcryptHelper
	redisRepository          repositories.RedisRepository
	uuidHelper               helpers.UuidHelper
	loginService             services.LoginService
}

func TestLoginTestSuite(t *testing.T) {
	suite.Run(t, new(LoginServiceTestSuite))
}

func (sut *LoginServiceTestSuite) SetupSuite() {
	sut.T().Log("SetupSuite")
	sut.mysqlUtil = utils.NewMysqlConnection("root", "12345", "localhost:3309", "user", 10, 10, 10, 10)
	sut.redisUtil = utils.NewRedisConnection("localhost", "6380", 0)
	sut.validate = validator.New()
	helpers.UsernameValidator(sut.validate)
	helpers.PasswordValidator(sut.validate)
	helpers.TelephoneValidator(sut.validate)
	sut.userRepository = repositories.NewUserRepository()
	sut.userPermissionRepository = repositories.NewUserPermissinoRepository()
	sut.bcryptHelper = helpers.NewBcryptHelper()
	sut.redisRepository = repositories.NewRedisRepository()
	sut.uuidHelper = helpers.NewUuidHelper()
	sut.loginService = services.NewLoginService(sut.mysqlUtil, sut.redisUtil, sut.validate, sut.userRepository, sut.userPermissionRepository, sut.bcryptHelper, sut.redisRepository, sut.uuidHelper)
}

func (sut *LoginServiceTestSuite) SetupTest() {
	sut.T().Log("SetupTest")
	sut.ctx = context.Background()
	sut.requestId = "requestId"
	sut.sessionId = "sessionId"
	sut.loginUserRequest = models.LoginUserRequest{
		Email:    "email@email.com",
		Password: "password@A1",
	}
}

func (sut *LoginServiceTestSuite) BeforeTest(suiteName, testName string) {
	sut.T().Log("BeforeTest: " + suiteName + " " + testName)
}

func (sut *LoginServiceTestSuite) TestLoginRedisRepositoryDelWithSessionIdInternalServerError() {
	sut.T().Log("TestLoginRedisRepositoryDelWithSessionIdInternalServerError")
	initialize.DropTableUserPermission(sut.mysqlUtil.GetDb(), sut.ctx)
	initialize.DropTablePermission(sut.mysqlUtil.GetDb(), sut.ctx)
	initialize.DropTableUser(sut.mysqlUtil.GetDb(), sut.ctx)
	sut.loginUserRequest = models.LoginUserRequest{}
	response, err := sut.loginService.Login(sut.ctx, sut.requestId, sut.sessionId, sut.loginUserRequest)
	sut.Equal(response, "")
	// sut.Equal(err.Error(), `[{"field":"email","message":"is required"},{"field":"password","message":"is required"}]`)
	sut.Equal(err.Error(), "validation error")
}

func (sut *LoginServiceTestSuite) TestLoginUserRepositoryFindByEmailInternalServerError() {
	sut.T().Log("TestLoginUserRepositoryFindByEmailInternalServerError")
	initialize.DropTableUserPermission(sut.mysqlUtil.GetDb(), sut.ctx)
	initialize.DropTablePermission(sut.mysqlUtil.GetDb(), sut.ctx)
	initialize.DropTableUser(sut.mysqlUtil.GetDb(), sut.ctx)
	response, err := sut.loginService.Login(sut.ctx, sut.requestId, sut.sessionId, sut.loginUserRequest)
	sut.Equal(response, "")
	// sut.Equal(err.Error(), `[{"field":"message","message":"internal server error"}]`)
	sut.Equal(err.Error(), "internal server error")
}

func (sut *LoginServiceTestSuite) TestLoginUserRepositoryFindByEmailBadRequestWrongEmailOrPasswordError() {
	sut.T().Log("TestLoginUserRepositoryFindByEmailBadRequestWrongEmailOrPasswordError")
	initialize.DropTableUserPermission(sut.mysqlUtil.GetDb(), sut.ctx)
	initialize.DropTablePermission(sut.mysqlUtil.GetDb(), sut.ctx)
	initialize.DropTableUser(sut.mysqlUtil.GetDb(), sut.ctx)
	initialize.CreateTableUser(sut.mysqlUtil.GetDb(), sut.ctx)
	response, err := sut.loginService.Login(sut.ctx, sut.requestId, sut.sessionId, sut.loginUserRequest)
	sut.Equal(response, "")
	// sut.Equal(err.Error(), `[{"field":"message","message":"wrong email or password"}]`)
	sut.Equal(err.Error(), "wrong email or password")
}

func (sut *LoginServiceTestSuite) TestLoginBcryptCompareHashAndPasswordBadRequestWrongEmailOrPasswordError() {
	sut.T().Log("TestLoginBcryptCompareHashAndPasswordBadRequestWrongEmailOrPasswordError")
	initialize.DropTableUserPermission(sut.mysqlUtil.GetDb(), sut.ctx)
	initialize.DropTablePermission(sut.mysqlUtil.GetDb(), sut.ctx)
	initialize.DropTableUser(sut.mysqlUtil.GetDb(), sut.ctx)
	initialize.CreateTableUser(sut.mysqlUtil.GetDb(), sut.ctx)
	initialize.CreateDataUser(sut.mysqlUtil.GetDb(), sut.ctx)
	sut.loginUserRequest.Password = "password@A1-"
	response, err := sut.loginService.Login(sut.ctx, sut.requestId, sut.sessionId, sut.loginUserRequest)
	sut.Equal(response, "")
	// sut.Equal(err.Error(), `[{"field":"message","message":"wrong email or password"}]`)
	sut.Equal(err.Error(), "wrong email or password")
}

func (sut *LoginServiceTestSuite) TestLoginUserPermissionRepositoryFindByUserIdInternalServerError() {
	sut.T().Log("TestLoginUserPermissionRepositoryFindByUserIdInternalServerError")
	initialize.DropTableUserPermission(sut.mysqlUtil.GetDb(), sut.ctx)
	initialize.DropTablePermission(sut.mysqlUtil.GetDb(), sut.ctx)
	initialize.DropTableUser(sut.mysqlUtil.GetDb(), sut.ctx)
	initialize.CreateTableUser(sut.mysqlUtil.GetDb(), sut.ctx)
	initialize.CreateDataUser(sut.mysqlUtil.GetDb(), sut.ctx)
	response, err := sut.loginService.Login(sut.ctx, sut.requestId, sut.sessionId, sut.loginUserRequest)
	sut.Equal(response, "")
	// sut.Equal(err.Error(), `[{"field":"message","message":"internal server error"}]`)
	sut.Equal(err.Error(), "internal server error")
}

func (sut *LoginServiceTestSuite) TestLoginSuccess() {
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
	response, err := sut.loginService.Login(sut.ctx, sut.requestId, sut.sessionId, sut.loginUserRequest)
	sut.NotEqual(response, "")
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
