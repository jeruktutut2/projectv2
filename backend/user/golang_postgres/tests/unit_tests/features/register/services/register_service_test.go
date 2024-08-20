package services_test

import (
	"context"
	"errors"
	"fmt"
	"golang-postgres/commons/helpers"
	"golang-postgres/features/register/models"
	"golang-postgres/features/register/services"
	mockhelpers "golang-postgres/tests/unit_tests/commons/helpers/mocks"
	mockutils "golang-postgres/tests/unit_tests/commons/utils/mocks"
	mockrepositories "golang-postgres/tests/unit_tests/features/register/mocks/repositories"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"
)

type RegisterServiceTestSuite struct {
	suite.Suite
	ctx                 context.Context
	requestId           string
	pool                *pgxpool.Pool
	errTimeout          error
	errInternalServer   error
	nowUnixMili         int64
	registerUserRequest models.RegisterUserRequest
	validate            *validator.Validate
	bcryptHelperMock    *mockhelpers.BcryptHelperMock
	postgresUtilMock    *mockutils.PostgresUtilMock
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
	sut.pool = &pgxpool.Pool{}
	sut.errTimeout = context.Canceled
	sut.errInternalServer = errors.New("internal server error")
	sut.nowUnixMili = 1720886992508
	sut.validate = validator.New()
	helpers.UsernameValidator(sut.validate)
	helpers.PasswordValidator(sut.validate)
	helpers.TelephoneValidator(sut.validate)
}

func (sut *RegisterServiceTestSuite) SetupTest() {
	fmt.Println()
	sut.T().Log("SetupTest")
	sut.registerUserRequest = models.RegisterUserRequest{
		Username:        "username17",
		Email:           "email17@email.com",
		Password:        "password@A1",
		Confirmpassword: "password@A1",
		Utc:             "+0800",
	}
	sut.bcryptHelperMock = new(mockhelpers.BcryptHelperMock)
	sut.userRepositoryMock = new(mockrepositories.UserRepositoryMock)
	sut.postgresUtilMock = new(mockutils.PostgresUtilMock)
	sut.registerService = services.NewRegisterService(sut.postgresUtilMock, sut.validate, sut.bcryptHelperMock, sut.userRepositoryMock)
}

func (sut *RegisterServiceTestSuite) BeforeTest(suiteName, testName string) {
	sut.T().Log("BeforeTest: " + suiteName + " " + testName)
}

func (sut *RegisterServiceTestSuite) Test01RegisterRegisterUserRequestValidationError() {
	sut.T().Log("Test01RegisterRegisterUserRequestValidationError")
	registerUserRequest := models.RegisterUserRequest{}
	response, err := sut.registerService.Register(sut.ctx, sut.requestId, sut.nowUnixMili, registerUserRequest)
	sut.Equal(response, models.RegisterUserResponse{})
	sut.Equal(err.Error(), "validation error")
}

func (sut *RegisterServiceTestSuite) Test02RegisterPasswordAndConfirmpasswordIsDifferent() {
	sut.T().Log("Test02RegisterPasswordAndConfirmpasswordIsDifferent")
	sut.registerUserRequest.Confirmpassword = "password@A1-"
	response, err := sut.registerService.Register(sut.ctx, sut.requestId, sut.nowUnixMili, sut.registerUserRequest)
	sut.Equal(response, models.RegisterUserResponse{})
	sut.Equal(err.Error(), "password and confirm password is different")
}

func (sut *RegisterServiceTestSuite) Test03RegisterUserRepositoryCountByEmailTimeoutError() {
	sut.T().Log("Test03RegisterUserRepositoryCountByEmailTimeoutError")
	sut.postgresUtilMock.Mock.On("GetPool").Return(sut.pool)
	sut.userRepositoryMock.Mock.On("CountByEmail", sut.pool, sut.ctx, sut.registerUserRequest.Email).Return(0, sut.errTimeout)
	response, err := sut.registerService.Register(sut.ctx, sut.requestId, sut.nowUnixMili, sut.registerUserRequest)
	sut.Equal(response, models.RegisterUserResponse{})
	sut.Equal(err.Error(), "time out or user cancel the request")
}

func (sut *RegisterServiceTestSuite) Test04RegisterUserRepositoryCountByEmailInternalServerError() {
	sut.T().Log("Test04RegisterUserRepositoryCountByEmailInternalServerError")
	sut.postgresUtilMock.Mock.On("GetPool").Return(sut.pool)
	sut.userRepositoryMock.Mock.On("CountByEmail", sut.pool, sut.ctx, sut.registerUserRequest.Email).Return(0, sut.errInternalServer)
	response, err := sut.registerService.Register(sut.ctx, sut.requestId, sut.nowUnixMili, sut.registerUserRequest)
	sut.Equal(response, models.RegisterUserResponse{})
	sut.Equal(err.Error(), "internal server error")
}

func (sut *RegisterServiceTestSuite) Test05RegisterUserRepositoryCountByEmailEmailAlreadyExists() {
	sut.T().Log("Test05RegisterUserRepositoryCountByEmailEmailAlreadyExists")
	sut.postgresUtilMock.Mock.On("GetPool").Return(sut.pool)
	sut.userRepositoryMock.Mock.On("CountByEmail", sut.pool, sut.ctx, sut.registerUserRequest.Email).Return(1, nil)
	response, err := sut.registerService.Register(sut.ctx, sut.requestId, sut.nowUnixMili, sut.registerUserRequest)
	sut.Equal(response, models.RegisterUserResponse{})
	sut.Equal(err.Error(), "email already exists")
}

func (sut *RegisterServiceTestSuite) Test06RegisterBcryptGenerateFromPasswordTimeoutError() {
	sut.T().Log("Test06RegisterBcryptGenerateFromPasswordTimeoutError")
	sut.postgresUtilMock.Mock.On("GetPool").Return(sut.pool)
	sut.userRepositoryMock.Mock.On("CountByEmail", sut.pool, sut.ctx, sut.registerUserRequest.Email).Return(0, nil)
	sut.bcryptHelperMock.Mock.On("GenerateFromPassword", []byte(sut.registerUserRequest.Password), bcrypt.DefaultCost).Return([]uint8{}, sut.errTimeout)
	response, err := sut.registerService.Register(sut.ctx, sut.requestId, sut.nowUnixMili, sut.registerUserRequest)
	sut.Equal(response, models.RegisterUserResponse{})
	sut.Equal(err.Error(), "time out or user cancel the request")
}

func (sut *RegisterServiceTestSuite) Test07RegisterBcryptGenerateFromPasswordInternalServerError() {
	sut.T().Log("Test07RegisterBcryptGenerateFromPasswordInternalServerError")
	sut.postgresUtilMock.Mock.On("GetPool").Return(sut.pool)
	sut.userRepositoryMock.Mock.On("CountByEmail", sut.pool, sut.ctx, sut.registerUserRequest.Email).Return(0, nil)
	sut.bcryptHelperMock.Mock.On("GenerateFromPassword", []byte(sut.registerUserRequest.Password), bcrypt.DefaultCost).Return([]uint8{}, sut.errInternalServer)
	response, err := sut.registerService.Register(sut.ctx, sut.requestId, sut.nowUnixMili, sut.registerUserRequest)
	sut.Equal(response, models.RegisterUserResponse{})
	sut.Equal(err.Error(), "internal server error")
}

func (sut *RegisterServiceTestSuite) Test08RegisterUserRepositoryCreateTimeoutError() {
	sut.T().Log("Test08RegisterUserRepositoryCreateTimeoutError")
	sut.postgresUtilMock.Mock.On("GetPool").Return(sut.pool)
	sut.userRepositoryMock.Mock.On("CountByEmail", sut.pool, sut.ctx, sut.registerUserRequest.Email).Return(0, nil)
	hashedPassword := []uint8{36, 50, 97, 36, 49, 48, 36, 121, 50, 52, 100, 109, 102, 86, 117, 108, 116, 68, 115, 118, 56, 55, 97, 77, 52, 105, 52, 89, 101, 76, 73, 74, 70, 118, 57, 104, 85, 54, 104, 119, 112, 97, 112, 65, 82, 102, 104, 84, 103, 121, 48, 116, 111, 75, 72, 78, 108, 114, 76, 83}
	sut.bcryptHelperMock.Mock.On("GenerateFromPassword", []byte(sut.registerUserRequest.Password), bcrypt.DefaultCost).Return(hashedPassword, nil)
	var user models.User
	user.Username = pgtype.Text{Valid: true, String: sut.registerUserRequest.Username}
	user.Email = pgtype.Text{Valid: true, String: sut.registerUserRequest.Email}
	user.Password = pgtype.Text{Valid: true, String: string(hashedPassword)}
	user.Utc = pgtype.Text{Valid: true, String: sut.registerUserRequest.Utc}
	user.CreatedAt = pgtype.Int8{Valid: true, Int64: sut.nowUnixMili}
	sut.userRepositoryMock.Mock.On("Create", sut.pool, sut.ctx, user).Return(int64(0), sut.errTimeout)
	response, err := sut.registerService.Register(sut.ctx, sut.requestId, sut.nowUnixMili, sut.registerUserRequest)
	sut.Equal(response, models.RegisterUserResponse{})
	sut.Equal(err.Error(), "time out or user cancel the request")
}

func (sut *RegisterServiceTestSuite) Test09RegisterUserRepositoryCreateInternalServerError() {
	sut.T().Log("Test09RegisterUserRepositoryCreateInternalServerError")
	sut.postgresUtilMock.Mock.On("GetPool").Return(sut.pool)
	sut.userRepositoryMock.Mock.On("CountByEmail", sut.pool, sut.ctx, sut.registerUserRequest.Email).Return(0, nil)
	hashedPassword := []uint8{36, 50, 97, 36, 49, 48, 36, 121, 50, 52, 100, 109, 102, 86, 117, 108, 116, 68, 115, 118, 56, 55, 97, 77, 52, 105, 52, 89, 101, 76, 73, 74, 70, 118, 57, 104, 85, 54, 104, 119, 112, 97, 112, 65, 82, 102, 104, 84, 103, 121, 48, 116, 111, 75, 72, 78, 108, 114, 76, 83}
	sut.bcryptHelperMock.Mock.On("GenerateFromPassword", []byte(sut.registerUserRequest.Password), bcrypt.DefaultCost).Return(hashedPassword, nil)
	var user models.User
	user.Username = pgtype.Text{Valid: true, String: sut.registerUserRequest.Username}
	user.Email = pgtype.Text{Valid: true, String: sut.registerUserRequest.Email}
	user.Password = pgtype.Text{Valid: true, String: string(hashedPassword)}
	user.Utc = pgtype.Text{Valid: true, String: sut.registerUserRequest.Utc}
	user.CreatedAt = pgtype.Int8{Valid: true, Int64: sut.nowUnixMili}
	sut.userRepositoryMock.Mock.On("Create", sut.pool, sut.ctx, user).Return(int64(0), sut.errInternalServer)
	response, err := sut.registerService.Register(sut.ctx, sut.requestId, sut.nowUnixMili, sut.registerUserRequest)
	sut.Equal(response, models.RegisterUserResponse{})
	sut.Equal(err.Error(), "internal server error")
}

func (sut *RegisterServiceTestSuite) Test10RegisterUserRepositoryCreateRowsAffectedNotOne() {
	sut.T().Log("Test10RegisterUserRepositoryCreateRowsAffectedNotOne")
	sut.postgresUtilMock.Mock.On("GetPool").Return(sut.pool)
	sut.userRepositoryMock.Mock.On("CountByEmail", sut.pool, sut.ctx, sut.registerUserRequest.Email).Return(0, nil)
	hashedPassword := []uint8{36, 50, 97, 36, 49, 48, 36, 121, 50, 52, 100, 109, 102, 86, 117, 108, 116, 68, 115, 118, 56, 55, 97, 77, 52, 105, 52, 89, 101, 76, 73, 74, 70, 118, 57, 104, 85, 54, 104, 119, 112, 97, 112, 65, 82, 102, 104, 84, 103, 121, 48, 116, 111, 75, 72, 78, 108, 114, 76, 83}
	sut.bcryptHelperMock.Mock.On("GenerateFromPassword", []byte(sut.registerUserRequest.Password), bcrypt.DefaultCost).Return(hashedPassword, nil)
	var user models.User
	user.Username = pgtype.Text{Valid: true, String: sut.registerUserRequest.Username}
	user.Email = pgtype.Text{Valid: true, String: sut.registerUserRequest.Email}
	user.Password = pgtype.Text{Valid: true, String: string(hashedPassword)}
	user.Utc = pgtype.Text{Valid: true, String: sut.registerUserRequest.Utc}
	user.CreatedAt = pgtype.Int8{Valid: true, Int64: sut.nowUnixMili}
	sut.userRepositoryMock.Mock.On("Create", sut.pool, sut.ctx, user).Return(int64(0), nil)
	response, err := sut.registerService.Register(sut.ctx, sut.requestId, sut.nowUnixMili, sut.registerUserRequest)
	sut.Equal(response, models.RegisterUserResponse{})
	sut.Equal(err.Error(), "internal server error")
}

func (sut *RegisterServiceTestSuite) Test11RegisterSuccess() {
	sut.T().Log("Test11RegisterSuccess")
	sut.postgresUtilMock.Mock.On("GetPool").Return(sut.pool)
	sut.userRepositoryMock.Mock.On("CountByEmail", sut.pool, sut.ctx, sut.registerUserRequest.Email).Return(0, nil)
	hashedPassword := []uint8{36, 50, 97, 36, 49, 48, 36, 121, 50, 52, 100, 109, 102, 86, 117, 108, 116, 68, 115, 118, 56, 55, 97, 77, 52, 105, 52, 89, 101, 76, 73, 74, 70, 118, 57, 104, 85, 54, 104, 119, 112, 97, 112, 65, 82, 102, 104, 84, 103, 121, 48, 116, 111, 75, 72, 78, 108, 114, 76, 83}
	sut.bcryptHelperMock.Mock.On("GenerateFromPassword", []byte(sut.registerUserRequest.Password), bcrypt.DefaultCost).Return(hashedPassword, nil)
	var user models.User
	user.Username = pgtype.Text{Valid: true, String: sut.registerUserRequest.Username}
	user.Email = pgtype.Text{Valid: true, String: sut.registerUserRequest.Email}
	user.Password = pgtype.Text{Valid: true, String: string(hashedPassword)}
	user.Utc = pgtype.Text{Valid: true, String: sut.registerUserRequest.Utc}
	user.CreatedAt = pgtype.Int8{Valid: true, Int64: sut.nowUnixMili}
	sut.userRepositoryMock.Mock.On("Create", sut.pool, sut.ctx, user).Return(int64(1), nil)
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
	fmt.Println()
}

func (sut *RegisterServiceTestSuite) TearDownSuite() {
	sut.T().Log("TearDownSuite")
}
