package services_test

import (
	"context"
	"fmt"
	"golang-postgres/commons/helpers"
	"golang-postgres/commons/utils"
	"golang-postgres/features/login/models"
	"golang-postgres/features/login/repositories"
	"golang-postgres/features/login/services"
	"golang-postgres/tests/initialize"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/suite"
)

type LoginServiceTestSuite struct {
	suite.Suite
	ctx                      context.Context
	requestId                string
	sessionId                string
	loginUserRequest         models.LoginUserRequest
	postgresUtil             utils.PostgresUtil
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
	fmt.Println()
	sut.T().Log("SetupSuite")
	sut.postgresUtil = utils.NewPostgresConnection()
	sut.redisUtil = utils.NewRedisConnection()
	sut.validate = validator.New()
	helpers.UsernameValidator(sut.validate)
	helpers.PasswordValidator(sut.validate)
	helpers.TelephoneValidator(sut.validate)
	sut.userRepository = repositories.NewUserRepository()
	sut.userPermissionRepository = repositories.NewUserPermissinoRepository()
	sut.bcryptHelper = helpers.NewBcryptHelper()
	sut.redisRepository = repositories.NewRedisRepository()
	sut.uuidHelper = helpers.NewUuidHelper()
	sut.loginService = services.NewLoginService(sut.postgresUtil, sut.redisUtil, sut.validate, sut.userRepository, sut.userPermissionRepository, sut.bcryptHelper, sut.redisRepository, sut.uuidHelper)
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

func (sut *LoginServiceTestSuite) Test1LoginRedisRepositoryDelWithSessionIdInternalServerError() {
	sut.T().Log("Test1LoginRedisRepositoryDelWithSessionIdInternalServerError")
	initialize.DropTableUserPermission(sut.postgresUtil.GetPool(), sut.ctx)
	initialize.DropTablePermission(sut.postgresUtil.GetPool(), sut.ctx)
	initialize.DropTableUser(sut.postgresUtil.GetPool(), sut.ctx)
	sut.loginUserRequest = models.LoginUserRequest{}
	response, err := sut.loginService.Login(sut.ctx, sut.requestId, sut.sessionId, sut.loginUserRequest)
	sut.Equal(response, "")
	sut.Equal(err.Error(), "validation error")
}

func (sut *LoginServiceTestSuite) Test2LoginUserRepositoryFindByEmailInternalServerError() {
	sut.T().Log("Test2LoginUserRepositoryFindByEmailInternalServerError")
	initialize.DropTableUserPermission(sut.postgresUtil.GetPool(), sut.ctx)
	initialize.DropTablePermission(sut.postgresUtil.GetPool(), sut.ctx)
	initialize.DropTableUser(sut.postgresUtil.GetPool(), sut.ctx)
	response, err := sut.loginService.Login(sut.ctx, sut.requestId, sut.sessionId, sut.loginUserRequest)
	sut.Equal(response, "")
	sut.Equal(err.Error(), "internal server error")
}

func (sut *LoginServiceTestSuite) Test3LoginUserRepositoryFindByEmailBadRequestWrongEmailOrPasswordError() {
	sut.T().Log("Test3LoginUserRepositoryFindByEmailBadRequestWrongEmailOrPasswordError")
	initialize.DropTableUserPermission(sut.postgresUtil.GetPool(), sut.ctx)
	initialize.DropTablePermission(sut.postgresUtil.GetPool(), sut.ctx)
	initialize.DropTableUser(sut.postgresUtil.GetPool(), sut.ctx)
	initialize.CreateTableUser(sut.postgresUtil.GetPool(), sut.ctx)
	response, err := sut.loginService.Login(sut.ctx, sut.requestId, sut.sessionId, sut.loginUserRequest)
	sut.Equal(response, "")
	sut.Equal(err.Error(), "wrong email or password")
}

func (sut *LoginServiceTestSuite) Test4LoginBcryptCompareHashAndPasswordBadRequestWrongEmailOrPasswordError() {
	sut.T().Log("Test4LoginBcryptCompareHashAndPasswordBadRequestWrongEmailOrPasswordError")
	initialize.DropTableUserPermission(sut.postgresUtil.GetPool(), sut.ctx)
	initialize.DropTablePermission(sut.postgresUtil.GetPool(), sut.ctx)
	initialize.DropTableUser(sut.postgresUtil.GetPool(), sut.ctx)
	initialize.CreateTableUser(sut.postgresUtil.GetPool(), sut.ctx)
	initialize.CreateDataUser(sut.postgresUtil.GetPool(), sut.ctx)
	sut.loginUserRequest.Password = "password@A1-"
	response, err := sut.loginService.Login(sut.ctx, sut.requestId, sut.sessionId, sut.loginUserRequest)
	sut.Equal(response, "")
	sut.Equal(err.Error(), "wrong email or password")
}

func (sut *LoginServiceTestSuite) Test5LoginUserPermissionRepositoryFindByUserIdInternalServerError() {
	sut.T().Log("Test5LoginUserPermissionRepositoryFindByUserIdInternalServerError")
	initialize.DropTableUserPermission(sut.postgresUtil.GetPool(), sut.ctx)
	initialize.DropTablePermission(sut.postgresUtil.GetPool(), sut.ctx)
	initialize.DropTableUser(sut.postgresUtil.GetPool(), sut.ctx)
	initialize.CreateTableUser(sut.postgresUtil.GetPool(), sut.ctx)
	initialize.CreateDataUser(sut.postgresUtil.GetPool(), sut.ctx)
	response, err := sut.loginService.Login(sut.ctx, sut.requestId, sut.sessionId, sut.loginUserRequest)
	sut.Equal(response, "")
	sut.Equal(err.Error(), "internal server error")
}

func (sut *LoginServiceTestSuite) Test6LoginSuccess() {
	sut.T().Log("Test6LoginSuccess")
	initialize.DropTableUserPermission(sut.postgresUtil.GetPool(), sut.ctx)
	initialize.DropTablePermission(sut.postgresUtil.GetPool(), sut.ctx)
	initialize.DropTableUser(sut.postgresUtil.GetPool(), sut.ctx)
	initialize.CreateTableUser(sut.postgresUtil.GetPool(), sut.ctx)
	initialize.CreateTablePermission(sut.postgresUtil.GetPool(), sut.ctx)
	initialize.CreateTableUserPermission(sut.postgresUtil.GetPool(), sut.ctx)
	initialize.CreateDataUser(sut.postgresUtil.GetPool(), sut.ctx)
	initialize.CreateDataPermission(sut.postgresUtil.GetPool(), sut.ctx)
	initialize.CreateDataUserPermission(sut.postgresUtil.GetPool(), sut.ctx)
	response, err := sut.loginService.Login(sut.ctx, sut.requestId, sut.sessionId, sut.loginUserRequest)
	sut.NotEqual(response, "")
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
