package services_test

import (
	"context"
	"database/sql"
	"errors"
	"project-user/features/register/models"
	"project-user/features/register/services"
	"project-user/helpers"
	mockrepositories "project-user/tests/unit_tests/features/register/mocks/repositories"
	mockhelpers "project-user/tests/unit_tests/mocks/helpers"
	mockutils "project-user/tests/unit_tests/mocks/utils"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"
)

type RegisterServiceTestSuite struct {
	suite.Suite
	ctx                 context.Context
	requestId           string
	db                  *sql.DB
	errTimeout          error
	errInternalServer   error
	nowUnixMili         int64
	registerUserRequest models.RegisterUserRequest
	validate            *validator.Validate
	bcryptHelperMock    *mockhelpers.BcryptHelperMock
	mysqlUtilMock       *mockutils.MysqlUtilMock
	userRepositoryMock  *mockrepositories.UserRepositoryMock
	registerService     services.RegisterService
}

func TestRegisterTestSuite(t *testing.T) {
	suite.Run(t, new(RegisterServiceTestSuite))
}

func (sut *RegisterServiceTestSuite) SetupSuite() {
	sut.T().Log("SetupSuite")
	sut.ctx = context.Background()
	sut.requestId = "requestId"
	sut.db = &sql.DB{}
	sut.errTimeout = context.Canceled
	sut.errInternalServer = errors.New("internal server error")
	sut.nowUnixMili = 1720886992508
	sut.validate = validator.New()
	helpers.UsernameValidator(sut.validate)
	helpers.PasswordValidator(sut.validate)
	helpers.TelephoneValidator(sut.validate)
	// sut.bcryptHelperMock = new(mockhelpers.BcryptHelperMock)
	// sut.userRepositoryMock = new(mockrepositories.UserRepositoryMock)
	// sut.mysqlUtilMock = new(mockutils.MysqlUtilMock)
	// sut.registerService = services.NewRegisterService(sut.mysqlUtilMock, sut.validate, sut.bcryptHelperMock, sut.userRepositoryMock)
}

func (sut *RegisterServiceTestSuite) SetupTest() {
	sut.T().Log("SetupTest")
	// sut.registerService = services.NewRegisterService(sut.mysqlUtilMock, sut.validate, sut.bcryptHelperMock, sut.userRepositoryMock)
	sut.registerUserRequest = models.RegisterUserRequest{
		Username:        "username17",
		Email:           "email17@email.com",
		Password:        "password@A1",
		Confirmpassword: "password@A1",
		Utc:             "+0800",
	}
	sut.bcryptHelperMock = new(mockhelpers.BcryptHelperMock)
	sut.userRepositoryMock = new(mockrepositories.UserRepositoryMock)
	sut.mysqlUtilMock = new(mockutils.MysqlUtilMock)
	sut.registerService = services.NewRegisterService(sut.mysqlUtilMock, sut.validate, sut.bcryptHelperMock, sut.userRepositoryMock)
}

func (sut *RegisterServiceTestSuite) BeforeTest(suiteName, testName string) {
	sut.T().Log("BeforeTest: " + suiteName + " " + testName)
}

func (sut *RegisterServiceTestSuite) TestRegisterRegisterUserRequestValidationError() {
	sut.T().Log("TestRegisterRegisterUserRequestValidationError")
	registerUserRequest := models.RegisterUserRequest{}
	response, err := sut.registerService.Register(sut.ctx, sut.requestId, sut.nowUnixMili, registerUserRequest)
	sut.Equal(response, models.RegisterUserResponse{})
	// sut.Equal(err.Error(), `[{"field":"username","message":"is required"},{"field":"email","message":"is required"},{"field":"password","message":"is required"},{"field":"confirmpassword","message":"is required"},{"field":"utc","message":"is required"}]`)
	sut.Equal(err.Error(), "validation error")
}

func (sut *RegisterServiceTestSuite) TestRegisterPasswordAndConfirmpasswordIsDifferent() {
	sut.T().Log("TestRegisterPasswordAndConfirmpasswordIsDifferent")
	sut.registerUserRequest.Confirmpassword = "password@A1-"
	response, err := sut.registerService.Register(sut.ctx, sut.requestId, sut.nowUnixMili, sut.registerUserRequest)
	sut.Equal(response, models.RegisterUserResponse{})
	// sut.Equal(err.Error(), `[{"field":"message","message":"password and confirm password is different"}]`)
	sut.Equal(err.Error(), "password and confirm password is different")
}

func (sut *RegisterServiceTestSuite) TestRegisterUserRepositoryCountByUsernameTimeoutError() {
	sut.T().Log("TestRegisterUserRepositoryCountByUsernameTimeoutError")
	sut.mysqlUtilMock.Mock.On("GetDb").Return(sut.db)
	sut.userRepositoryMock.Mock.On("CountByUsername", sut.db, sut.ctx, sut.registerUserRequest.Username).Return(0, sut.errTimeout)
	response, err := sut.registerService.Register(sut.ctx, sut.requestId, sut.nowUnixMili, sut.registerUserRequest)
	sut.Equal(response, models.RegisterUserResponse{})
	// sut.Equal(err.Error(), `[{"field":"message","message":"time out or user cancel the request"}]`)
	sut.Equal(err.Error(), "time out or user cancel the request")
}

func (sut *RegisterServiceTestSuite) TestRegisterUserRepositoryCountByUsernameInternalServerError() {
	sut.T().Log("TestRegisterUserRepositoryCountByUsernameInternalServerError")
	sut.mysqlUtilMock.Mock.On("GetDb").Return(sut.db)
	sut.userRepositoryMock.Mock.On("CountByUsername", sut.db, sut.ctx, sut.registerUserRequest.Username).Return(0, sut.errInternalServer)
	response, err := sut.registerService.Register(sut.ctx, sut.requestId, sut.nowUnixMili, sut.registerUserRequest)
	sut.Equal(response, models.RegisterUserResponse{})
	// sut.Equal(err.Error(), `[{"field":"message","message":"internal server error"}]`)
	sut.Equal(err.Error(), "internal server error")
}

func (sut *RegisterServiceTestSuite) TestRegisterUserRepositoryCountByUsernameUsernameAlreadyExists() {
	sut.T().Log("TestRegisterUserRepositoryCountByUsernameUsernameAlreadyExists")
	sut.mysqlUtilMock.Mock.On("GetDb").Return(sut.db)
	sut.userRepositoryMock.Mock.On("CountByUsername", sut.db, sut.ctx, sut.registerUserRequest.Username).Return(1, nil)
	response, err := sut.registerService.Register(sut.ctx, sut.requestId, sut.nowUnixMili, sut.registerUserRequest)
	sut.Equal(response, models.RegisterUserResponse{})
	// sut.Equal(err.Error(), `[{"field":"message","message":"username already exists"}]`)
	sut.Equal(err.Error(), "username already exists")
}

func (sut *RegisterServiceTestSuite) TestRegisterUserRepositoryCountByEmailTimeoutError() {
	sut.T().Log("TestRegisterUserRepositoryCountByEmailTimeoutError")
	sut.mysqlUtilMock.Mock.On("GetDb").Return(sut.db)
	sut.userRepositoryMock.Mock.On("CountByUsername", sut.db, sut.ctx, sut.registerUserRequest.Username).Return(0, nil)
	sut.userRepositoryMock.Mock.On("CountByEmail", sut.db, sut.ctx, sut.registerUserRequest.Email).Return(0, sut.errTimeout)
	response, err := sut.registerService.Register(sut.ctx, sut.requestId, sut.nowUnixMili, sut.registerUserRequest)
	sut.Equal(response, models.RegisterUserResponse{})
	// sut.Equal(err.Error(), `[{"field":"message","message":"time out or user cancel the request"}]`)
	sut.Equal(err.Error(), "time out or user cancel the request")
}

func (sut *RegisterServiceTestSuite) TestRegisterUserRepositoryCountByEmailInternalServerError() {
	sut.T().Log("TestRegisterUserRepositoryCountByEmailInternalServerError")
	sut.mysqlUtilMock.Mock.On("GetDb").Return(sut.db)
	sut.userRepositoryMock.Mock.On("CountByUsername", sut.db, sut.ctx, sut.registerUserRequest.Username).Return(0, nil)
	sut.userRepositoryMock.Mock.On("CountByEmail", sut.db, sut.ctx, sut.registerUserRequest.Email).Return(0, sut.errInternalServer)
	response, err := sut.registerService.Register(sut.ctx, sut.requestId, sut.nowUnixMili, sut.registerUserRequest)
	sut.Equal(response, models.RegisterUserResponse{})
	// sut.Equal(err.Error(), `[{"field":"message","message":"internal server error"}]`)
	sut.Equal(err.Error(), "internal server error")
}

func (sut *RegisterServiceTestSuite) TestRegisterUserRepositoryCountByEmailEmailAlreadyExists() {
	sut.T().Log("TestRegisterUserRepositoryCountByEmailEmailAlreadyExists")
	sut.mysqlUtilMock.Mock.On("GetDb").Return(sut.db)
	sut.userRepositoryMock.Mock.On("CountByUsername", sut.db, sut.ctx, sut.registerUserRequest.Username).Return(0, nil)
	sut.userRepositoryMock.Mock.On("CountByEmail", sut.db, sut.ctx, sut.registerUserRequest.Email).Return(1, nil)
	response, err := sut.registerService.Register(sut.ctx, sut.requestId, sut.nowUnixMili, sut.registerUserRequest)
	sut.Equal(response, models.RegisterUserResponse{})
	// sut.Equal(err.Error(), `[{"field":"message","message":"email already exists"}]`)
	sut.Equal(err.Error(), "email already exists")
}

func (sut *RegisterServiceTestSuite) TestRegisterBcryptGenerateFromPasswordTimeoutError() {
	sut.T().Log("TestRegisterBcryptGenerateFromPasswordTimeoutError")
	sut.mysqlUtilMock.Mock.On("GetDb").Return(sut.db)
	sut.userRepositoryMock.Mock.On("CountByUsername", sut.db, sut.ctx, sut.registerUserRequest.Username).Return(0, nil)
	sut.userRepositoryMock.Mock.On("CountByEmail", sut.db, sut.ctx, sut.registerUserRequest.Email).Return(0, nil)
	sut.bcryptHelperMock.Mock.On("GenerateFromPassword", []byte(sut.registerUserRequest.Password), bcrypt.DefaultCost).Return([]uint8{}, sut.errTimeout)
	response, err := sut.registerService.Register(sut.ctx, sut.requestId, sut.nowUnixMili, sut.registerUserRequest)
	sut.Equal(response, models.RegisterUserResponse{})
	// sut.Equal(err.Error(), `[{"field":"message","message":"time out or user cancel the request"}]`)
	sut.Equal(err.Error(), "time out or user cancel the request")
}

func (sut *RegisterServiceTestSuite) TestRegisterBcryptGenerateFromPasswordInternalServerError() {
	sut.T().Log("TestRegisterBcryptGenerateFromPasswordInternalServerError")
	sut.mysqlUtilMock.Mock.On("GetDb").Return(sut.db)
	sut.userRepositoryMock.Mock.On("CountByUsername", sut.db, sut.ctx, sut.registerUserRequest.Username).Return(0, nil)
	sut.userRepositoryMock.Mock.On("CountByEmail", sut.db, sut.ctx, sut.registerUserRequest.Email).Return(0, nil)
	sut.bcryptHelperMock.Mock.On("GenerateFromPassword", []byte(sut.registerUserRequest.Password), bcrypt.DefaultCost).Return([]uint8{}, sut.errInternalServer)
	response, err := sut.registerService.Register(sut.ctx, sut.requestId, sut.nowUnixMili, sut.registerUserRequest)
	sut.Equal(response, models.RegisterUserResponse{})
	// sut.Equal(err.Error(), `[{"field":"message","message":"internal server error"}]`)
	sut.Equal(err.Error(), "internal server error")
}

func (sut *RegisterServiceTestSuite) TestRegisterUserRepositoryCreateTimeoutError() {
	sut.T().Log("TestRegisterUserRepositoryCreateTimeoutError")
	sut.mysqlUtilMock.Mock.On("GetDb").Return(sut.db)
	sut.userRepositoryMock.Mock.On("CountByUsername", sut.db, sut.ctx, sut.registerUserRequest.Username).Return(0, nil)
	sut.userRepositoryMock.Mock.On("CountByEmail", sut.db, sut.ctx, sut.registerUserRequest.Email).Return(0, nil)
	hashedPassword := []uint8{36, 50, 97, 36, 49, 48, 36, 121, 50, 52, 100, 109, 102, 86, 117, 108, 116, 68, 115, 118, 56, 55, 97, 77, 52, 105, 52, 89, 101, 76, 73, 74, 70, 118, 57, 104, 85, 54, 104, 119, 112, 97, 112, 65, 82, 102, 104, 84, 103, 121, 48, 116, 111, 75, 72, 78, 108, 114, 76, 83}
	sut.bcryptHelperMock.Mock.On("GenerateFromPassword", []byte(sut.registerUserRequest.Password), bcrypt.DefaultCost).Return(hashedPassword, nil)
	// var nowUnixMili int64 = 1719496855216
	// sut.timeHelperMock.Mock.On("NowUnixMili").Return(nowUnixMili)
	var user models.User
	user.Username = sql.NullString{Valid: true, String: sut.registerUserRequest.Username}
	user.Email = sql.NullString{Valid: true, String: sut.registerUserRequest.Email}
	user.Password = sql.NullString{Valid: true, String: string(hashedPassword)}
	user.Utc = sql.NullString{Valid: true, String: sut.registerUserRequest.Utc}
	user.CreatedAt = sql.NullInt64{Valid: true, Int64: sut.nowUnixMili}
	sut.userRepositoryMock.Mock.On("Create", sut.db, sut.ctx, user).Return(int64(0), sut.errTimeout)
	response, err := sut.registerService.Register(sut.ctx, sut.requestId, sut.nowUnixMili, sut.registerUserRequest)
	sut.Equal(response, models.RegisterUserResponse{})
	// sut.Equal(err.Error(), `[{"field":"message","message":"time out or user cancel the request"}]`)
	sut.Equal(err.Error(), "time out or user cancel the request")
}

func (sut *RegisterServiceTestSuite) TestRegisterUserRepositoryCreateInternalServerError() {
	sut.T().Log("TestRegisterUserRepositoryCreateInternalServerError")
	sut.mysqlUtilMock.Mock.On("GetDb").Return(sut.db)
	sut.userRepositoryMock.Mock.On("CountByUsername", sut.db, sut.ctx, sut.registerUserRequest.Username).Return(0, nil)
	sut.userRepositoryMock.Mock.On("CountByEmail", sut.db, sut.ctx, sut.registerUserRequest.Email).Return(0, nil)
	hashedPassword := []uint8{36, 50, 97, 36, 49, 48, 36, 121, 50, 52, 100, 109, 102, 86, 117, 108, 116, 68, 115, 118, 56, 55, 97, 77, 52, 105, 52, 89, 101, 76, 73, 74, 70, 118, 57, 104, 85, 54, 104, 119, 112, 97, 112, 65, 82, 102, 104, 84, 103, 121, 48, 116, 111, 75, 72, 78, 108, 114, 76, 83}
	sut.bcryptHelperMock.Mock.On("GenerateFromPassword", []byte(sut.registerUserRequest.Password), bcrypt.DefaultCost).Return(hashedPassword, nil)
	// var nowUnixMili int64 = 1719496855216
	// sut.timeHelperMock.Mock.On("NowUnixMili").Return(nowUnixMili)
	var user models.User
	user.Username = sql.NullString{Valid: true, String: sut.registerUserRequest.Username}
	user.Email = sql.NullString{Valid: true, String: sut.registerUserRequest.Email}
	user.Password = sql.NullString{Valid: true, String: string(hashedPassword)}
	user.Utc = sql.NullString{Valid: true, String: sut.registerUserRequest.Utc}
	user.CreatedAt = sql.NullInt64{Valid: true, Int64: sut.nowUnixMili}
	sut.userRepositoryMock.Mock.On("Create", sut.db, sut.ctx, user).Return(int64(0), sut.errInternalServer)
	response, err := sut.registerService.Register(sut.ctx, sut.requestId, sut.nowUnixMili, sut.registerUserRequest)
	sut.Equal(response, models.RegisterUserResponse{})
	// sut.Equal(err.Error(), `[{"field":"message","message":"internal server error"}]`)
	sut.Equal(err.Error(), "internal server error")
}

func (sut *RegisterServiceTestSuite) TestRegisterUserRepositoryCreateRowsAffectedNotOne() {
	sut.T().Log("TestRegisterUserRepositoryCreateRowsAffectedNotOne")
	sut.mysqlUtilMock.Mock.On("GetDb").Return(sut.db)
	sut.userRepositoryMock.Mock.On("CountByUsername", sut.db, sut.ctx, sut.registerUserRequest.Username).Return(0, nil)
	sut.userRepositoryMock.Mock.On("CountByEmail", sut.db, sut.ctx, sut.registerUserRequest.Email).Return(0, nil)
	hashedPassword := []uint8{36, 50, 97, 36, 49, 48, 36, 121, 50, 52, 100, 109, 102, 86, 117, 108, 116, 68, 115, 118, 56, 55, 97, 77, 52, 105, 52, 89, 101, 76, 73, 74, 70, 118, 57, 104, 85, 54, 104, 119, 112, 97, 112, 65, 82, 102, 104, 84, 103, 121, 48, 116, 111, 75, 72, 78, 108, 114, 76, 83}
	sut.bcryptHelperMock.Mock.On("GenerateFromPassword", []byte(sut.registerUserRequest.Password), bcrypt.DefaultCost).Return(hashedPassword, nil)
	// var nowUnixMili int64 = 1719496855216
	// sut.timeHelperMock.Mock.On("NowUnixMili").Return(nowUnixMili)
	var user models.User
	user.Username = sql.NullString{Valid: true, String: sut.registerUserRequest.Username}
	user.Email = sql.NullString{Valid: true, String: sut.registerUserRequest.Email}
	user.Password = sql.NullString{Valid: true, String: string(hashedPassword)}
	user.Utc = sql.NullString{Valid: true, String: sut.registerUserRequest.Utc}
	user.CreatedAt = sql.NullInt64{Valid: true, Int64: sut.nowUnixMili}
	sut.userRepositoryMock.Mock.On("Create", sut.db, sut.ctx, user).Return(int64(0), nil)
	response, err := sut.registerService.Register(sut.ctx, sut.requestId, sut.nowUnixMili, sut.registerUserRequest)
	sut.Equal(response, models.RegisterUserResponse{})
	// sut.Equal(err.Error(), `[{"field":"message","message":"internal server error"}]`)
	sut.Equal(err.Error(), "internal server error")
}

func (sut *RegisterServiceTestSuite) TestRegisterSuccess() {
	sut.T().Log("TestRegisterSuccess")
	sut.mysqlUtilMock.Mock.On("GetDb").Return(sut.db)
	sut.userRepositoryMock.Mock.On("CountByUsername", sut.db, sut.ctx, sut.registerUserRequest.Username).Return(0, nil)
	sut.userRepositoryMock.Mock.On("CountByEmail", sut.db, sut.ctx, sut.registerUserRequest.Email).Return(0, nil)
	hashedPassword := []uint8{36, 50, 97, 36, 49, 48, 36, 121, 50, 52, 100, 109, 102, 86, 117, 108, 116, 68, 115, 118, 56, 55, 97, 77, 52, 105, 52, 89, 101, 76, 73, 74, 70, 118, 57, 104, 85, 54, 104, 119, 112, 97, 112, 65, 82, 102, 104, 84, 103, 121, 48, 116, 111, 75, 72, 78, 108, 114, 76, 83}
	sut.bcryptHelperMock.Mock.On("GenerateFromPassword", []byte(sut.registerUserRequest.Password), bcrypt.DefaultCost).Return(hashedPassword, nil)
	// var nowUnixMili int64 = 1719496855216
	// sut.timeHelperMock.Mock.On("NowUnixMili").Return(nowUnixMili)
	var user models.User
	user.Username = sql.NullString{Valid: true, String: sut.registerUserRequest.Username}
	user.Email = sql.NullString{Valid: true, String: sut.registerUserRequest.Email}
	user.Password = sql.NullString{Valid: true, String: string(hashedPassword)}
	user.Utc = sql.NullString{Valid: true, String: sut.registerUserRequest.Utc}
	user.CreatedAt = sql.NullInt64{Valid: true, Int64: sut.nowUnixMili}
	sut.userRepositoryMock.Mock.On("Create", sut.db, sut.ctx, user).Return(int64(1), nil)
	response, err := sut.registerService.Register(sut.ctx, sut.requestId, sut.nowUnixMili, sut.registerUserRequest)
	registerUserResponse := models.RegisterUserResponse{}
	registerUserResponse.Username = "username17"
	registerUserResponse.Email = "email17@email.com"
	registerUserResponse.Utc = "+0800"
	sut.Equal(response, registerUserResponse)
	sut.Equal(err, nil)
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
