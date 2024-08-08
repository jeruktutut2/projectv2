package services_test

import (
	"context"
	"project-user/features/register/models"
	"project-user/features/register/repositories"
	"project-user/features/register/services"
	"project-user/helpers"
	"project-user/tests/initialize"
	"project-user/utils"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/suite"
)

type RegisterServiceTestSuite struct {
	suite.Suite
	ctx                 context.Context
	requestId           string
	nowUnixMili         int64
	registerUserRequest models.RegisterUserRequest
	mysqlUtil           utils.MysqlUtil
	validate            *validator.Validate
	bcryptHelper        helpers.BcryptHelper
	userRepository      repositories.UserRepository
	registerService     services.RegisterService
}

func TestRegisterTestSuite(t *testing.T) {
	suite.Run(t, new(RegisterServiceTestSuite))
}

func (sut *RegisterServiceTestSuite) SetupSuite() {
	sut.T().Log("SetupSuite")
	sut.mysqlUtil = utils.NewMysqlConnection("root", "12345", "localhost:3309", "user", 10, 10, 10, 10)
	sut.validate = validator.New()
	helpers.UsernameValidator(sut.validate)
	helpers.PasswordValidator(sut.validate)
	helpers.TelephoneValidator(sut.validate)
	sut.bcryptHelper = helpers.NewBcryptHelper()
	sut.userRepository = repositories.NewUserRepository()
	sut.registerService = services.NewRegisterService(sut.mysqlUtil, sut.validate, sut.bcryptHelper, sut.userRepository)
}

func (sut *RegisterServiceTestSuite) SetupTest() {
	sut.T().Log("SetupTest")
	sut.ctx = context.Background()
	sut.requestId = "requestId"
	sut.nowUnixMili = time.Now().UnixMilli()
	sut.registerUserRequest = models.RegisterUserRequest{
		Username:        "username",
		Email:           "email@email.com",
		Password:        "password@A1",
		Confirmpassword: "password@A1",
		Utc:             "+0800",
	}
}

func (sut *RegisterServiceTestSuite) BeforeTest(suiteName, testName string) {
	sut.T().Log("BeforeTest: " + suiteName + " " + testName)
}

func (sut *RegisterServiceTestSuite) TestRegisterRegisterUserRequestValidationError() {
	sut.T().Log("TestRegisterRegisterUserRequestValidationError")
	registerUserRequest := models.RegisterUserRequest{}
	response, err := sut.registerService.Register(sut.ctx, sut.requestId, sut.nowUnixMili, registerUserRequest)
	sut.Equal(response, models.RegisterUserResponse{})
	sut.Equal(err.Error(), "validation error")
}

func (sut *RegisterServiceTestSuite) TestRegisterPasswordAndConfirmpasswordIsDifferent() {
	sut.T().Log("TestRegisterPasswordAndConfirmpasswordIsDifferent")
	sut.registerUserRequest.Confirmpassword = "password@A1-"
	response, err := sut.registerService.Register(sut.ctx, sut.requestId, sut.nowUnixMili, sut.registerUserRequest)
	sut.Equal(response, models.RegisterUserResponse{})
	sut.Equal(err.Error(), "password and confirm password is different")
}

func (sut *RegisterServiceTestSuite) TestRegisterUserRepositoryCountByUsernameInternalServerError() {
	sut.T().Log("TestRegisterUserRepositoryCountByUsernameInternalServerError")
	initialize.DropTableUser(sut.mysqlUtil.GetDb(), sut.ctx)
	response, err := sut.registerService.Register(sut.ctx, sut.requestId, sut.nowUnixMili, sut.registerUserRequest)
	sut.Equal(response, models.RegisterUserResponse{})
	sut.Equal(err.Error(), "internal server error")
}

func (sut *RegisterServiceTestSuite) TestRegisterUserRepositoryCountByUsernameBadRequestUsernameAlreadyExists() {
	sut.T().Log("TestRegisterUserRepositoryCountByUsernameBadRequestUsernameAlreadyExists")
	initialize.DropTableUser(sut.mysqlUtil.GetDb(), sut.ctx)
	initialize.CreateTableUser(sut.mysqlUtil.GetDb(), sut.ctx)
	initialize.CreateDataUser(sut.mysqlUtil.GetDb(), sut.ctx)
	response, err := sut.registerService.Register(sut.ctx, sut.requestId, sut.nowUnixMili, sut.registerUserRequest)
	sut.Equal(response, models.RegisterUserResponse{})
	sut.Equal(err.Error(), "username already exists")
}

func (sut *RegisterServiceTestSuite) TestRegisterUserRepositoryCountByCountByEmailBadRequestUsernameAlreadyExists() {
	sut.T().Log("TestRegisterUserRepositoryCountByUsernameBadRequestUsernameAlreadyExists")
	initialize.DropTableUser(sut.mysqlUtil.GetDb(), sut.ctx)
	initialize.CreateTableUser(sut.mysqlUtil.GetDb(), sut.ctx)
	initialize.CreateDataUser(sut.mysqlUtil.GetDb(), sut.ctx)
	sut.registerUserRequest.Username = "username1"
	response, err := sut.registerService.Register(sut.ctx, sut.requestId, sut.nowUnixMili, sut.registerUserRequest)
	sut.Equal(response, models.RegisterUserResponse{})
	sut.Equal(err.Error(), "email already exists")
}

func (sut *RegisterServiceTestSuite) TestRegisterSuccess() {
	sut.T().Log("TestRegisterSuccess")
	initialize.DropTableUser(sut.mysqlUtil.GetDb(), sut.ctx)
	initialize.CreateTableUser(sut.mysqlUtil.GetDb(), sut.ctx)
	initialize.CreateDataUser(sut.mysqlUtil.GetDb(), sut.ctx)
	sut.registerUserRequest.Username = "username1"
	sut.registerUserRequest.Email = "email1@email.com"
	response, err := sut.registerService.Register(sut.ctx, sut.requestId, sut.nowUnixMili, sut.registerUserRequest)
	registerUserResponse := models.RegisterUserResponse{}
	registerUserResponse.Username = "username1"
	registerUserResponse.Email = "email1@email.com"
	registerUserResponse.Utc = "+0800"
	sut.Equal(response, registerUserResponse)
	sut.Equal(err, nil)

	user := initialize.GetDataUserByEmail(sut.mysqlUtil.GetDb(), sut.ctx, registerUserResponse.Email)
	sut.Equal(user.Username.String, registerUserResponse.Username)
	sut.Equal(user.Email.String, registerUserResponse.Email)
	sut.Equal(user.Utc.String, registerUserResponse.Utc)
}

func (sut *RegisterServiceTestSuite) AfterTest(suiteName, testName string) {
	sut.T().Log("AfterTest: " + suiteName + " " + testName)
}

func (sut *RegisterServiceTestSuite) TearDownTest() {
	sut.T().Log("TearDownTest")
}

func (sut *RegisterServiceTestSuite) TearDownSuite() {
	sut.T().Log("TearDownSuite")
}
