package services_test

import (
	"context"
	"fmt"
	"golang-postgres/commons/helpers"
	"golang-postgres/commons/utils"
	"golang-postgres/features/register/models"
	"golang-postgres/features/register/repositories"
	"golang-postgres/features/register/services"
	"golang-postgres/tests/initialize"
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
	postgresUtil        utils.PostgresUtil
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
	sut.postgresUtil = utils.NewPostgresConnection()
	sut.validate = validator.New()
	helpers.UsernameValidator(sut.validate)
	helpers.PasswordValidator(sut.validate)
	helpers.TelephoneValidator(sut.validate)
	sut.bcryptHelper = helpers.NewBcryptHelper()
	sut.userRepository = repositories.NewUserRepository()
	sut.registerService = services.NewRegisterService(sut.postgresUtil, sut.validate, sut.bcryptHelper, sut.userRepository)
}

func (sut *RegisterServiceTestSuite) SetupTest() {
	fmt.Println()
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

func (sut *RegisterServiceTestSuite) Test1RegisterRegisterUserRequestValidationError() {
	sut.T().Log("Test1RegisterRegisterUserRequestValidationError")
	registerUserRequest := models.RegisterUserRequest{}
	response, err := sut.registerService.Register(sut.ctx, sut.requestId, sut.nowUnixMili, registerUserRequest)
	sut.Equal(response, models.RegisterUserResponse{})
	sut.Equal(err.Error(), "validation error")
}

func (sut *RegisterServiceTestSuite) Test2RegisterPasswordAndConfirmpasswordIsDifferent() {
	sut.T().Log("Test2RegisterPasswordAndConfirmpasswordIsDifferent")
	sut.registerUserRequest.Confirmpassword = "password@A1-"
	response, err := sut.registerService.Register(sut.ctx, sut.requestId, sut.nowUnixMili, sut.registerUserRequest)
	sut.Equal(response, models.RegisterUserResponse{})
	sut.Equal(err.Error(), "password and confirm password is different")
}

func (sut *RegisterServiceTestSuite) Test3RegisterUserRepositoryCountByUsernameInternalServerError() {
	sut.T().Log("Test3RegisterUserRepositoryCountByUsernameInternalServerError")
	initialize.DropTableUserPermission(sut.postgresUtil.GetPool(), sut.ctx)
	initialize.DropTablePermission(sut.postgresUtil.GetPool(), sut.ctx)
	initialize.DropTableUser(sut.postgresUtil.GetPool(), sut.ctx)
	response, err := sut.registerService.Register(sut.ctx, sut.requestId, sut.nowUnixMili, sut.registerUserRequest)
	sut.Equal(response, models.RegisterUserResponse{})
	sut.Equal(err.Error(), "internal server error")
}

func (sut *RegisterServiceTestSuite) Test4RegisterUserRepositoryCountByCountByEmailBadRequestUsernameAlreadyExists() {
	sut.T().Log("Test4RegisterUserRepositoryCountByUsernameBadRequestUsernameAlreadyExists")
	initialize.DropTableUserPermission(sut.postgresUtil.GetPool(), sut.ctx)
	initialize.DropTablePermission(sut.postgresUtil.GetPool(), sut.ctx)
	initialize.DropTableUser(sut.postgresUtil.GetPool(), sut.ctx)
	initialize.CreateTableUser(sut.postgresUtil.GetPool(), sut.ctx)
	initialize.CreateDataUser(sut.postgresUtil.GetPool(), sut.ctx)
	sut.registerUserRequest.Username = "username1"
	response, err := sut.registerService.Register(sut.ctx, sut.requestId, sut.nowUnixMili, sut.registerUserRequest)
	sut.Equal(response, models.RegisterUserResponse{})
	sut.Equal(err.Error(), "email already exists")
}

func (sut *RegisterServiceTestSuite) Test5RegisterSuccess() {
	sut.T().Log("Test5RegisterSuccess")
	initialize.DropTableUserPermission(sut.postgresUtil.GetPool(), sut.ctx)
	initialize.DropTablePermission(sut.postgresUtil.GetPool(), sut.ctx)
	initialize.DropTableUser(sut.postgresUtil.GetPool(), sut.ctx)
	initialize.CreateTableUser(sut.postgresUtil.GetPool(), sut.ctx)
	// initialize.CreateDataUser(sut.postgresUtil.GetPool(), sut.ctx)
	sut.registerUserRequest.Username = "username1"
	sut.registerUserRequest.Email = "email1@email.com"
	response, err := sut.registerService.Register(sut.ctx, sut.requestId, sut.nowUnixMili, sut.registerUserRequest)
	registerUserResponse := models.RegisterUserResponse{}
	registerUserResponse.Username = "username1"
	registerUserResponse.Email = "email1@email.com"
	registerUserResponse.Utc = "+0800"
	sut.Equal(response, registerUserResponse)
	sut.Equal(err, nil)

	user := initialize.GetDataUserByEmail(sut.postgresUtil.GetPool(), sut.ctx, registerUserResponse.Email)
	sut.Equal(user.Username.String, registerUserResponse.Username)
	sut.Equal(user.Email.String, registerUserResponse.Email)
	sut.Equal(user.Utc.String, registerUserResponse.Utc)
}

func (sut *RegisterServiceTestSuite) AfterTest(suiteName, testName string) {
	sut.T().Log("AfterTest: " + suiteName + " " + testName)
}

func (sut *RegisterServiceTestSuite) TearDownTest() {
	sut.T().Log("TearDownTest")
	fmt.Println()
}

func (sut *RegisterServiceTestSuite) TearDownSuite() {
	sut.T().Log("TearDownSuite")
}
