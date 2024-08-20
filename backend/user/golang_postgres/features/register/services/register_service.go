package services

import (
	"context"
	"errors"
	"golang-postgres/commons/exceptions"
	"golang-postgres/commons/helpers"
	"golang-postgres/commons/utils"
	"golang-postgres/features/register/models"
	"golang-postgres/features/register/repositories"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/crypto/bcrypt"
)

type RegisterService interface {
	Register(ctx context.Context, requestId string, nowUnixMili int64, registerUserRequest models.RegisterUserRequest) (registerUserResponse models.RegisterUserResponse, err error)
}

type RegisterServiceImplementation struct {
	PostgresUtil   utils.PostgresUtil
	Validate       *validator.Validate
	BcryptHelper   helpers.BcryptHelper
	UserRepository repositories.UserRepository
}

func NewRegisterService(postgressUtil utils.PostgresUtil, validate *validator.Validate, bcryptHelper helpers.BcryptHelper, userRepository repositories.UserRepository) RegisterService {
	return &RegisterServiceImplementation{
		PostgresUtil:   postgressUtil,
		Validate:       validate,
		BcryptHelper:   bcryptHelper,
		UserRepository: userRepository,
	}
}

func (service *RegisterServiceImplementation) Register(ctx context.Context, requestId string, nowUnixMili int64, registerUserRequest models.RegisterUserRequest) (registerUserResponse models.RegisterUserResponse, err error) {
	err = service.Validate.Struct(registerUserRequest)
	if err != nil {
		validationResult := helpers.GetValidatorError(err, registerUserRequest)
		if validationResult != nil {
			err = exceptions.ToResponseExceptionRequestValidation(requestId, validationResult)
			return
		}
	}
	if registerUserRequest.Password != registerUserRequest.Confirmpassword {
		err = errors.New("password and confirm password is different")
		err = exceptions.ToResponseException(err, requestId, http.StatusBadRequest, "password and confirm password is different")
		return
	}

	numberOfUser, err := service.UserRepository.CountByEmail(service.PostgresUtil.GetPool(), ctx, registerUserRequest.Email)
	if err != nil && err != pgx.ErrNoRows {
		err = exceptions.CheckError(err, requestId)
		return
	}
	if numberOfUser > 0 {
		err = errors.New("email already exists")
		err = exceptions.ToResponseException(err, requestId, http.StatusBadRequest, "email already exists")
		return
	}

	hashedPassword, err := service.BcryptHelper.GenerateFromPassword([]byte(registerUserRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		err = exceptions.CheckError(err, requestId)
		return
	}

	var user models.User
	user.Username = pgtype.Text{Valid: true, String: registerUserRequest.Username}
	user.Email = pgtype.Text{Valid: true, String: registerUserRequest.Email}
	user.Password = pgtype.Text{Valid: true, String: string(hashedPassword)}
	user.Utc = pgtype.Text{Valid: true, String: registerUserRequest.Utc}
	user.CreatedAt = pgtype.Int8{Valid: true, Int64: nowUnixMili}
	rowsAffected, err := service.UserRepository.Create(service.PostgresUtil.GetPool(), ctx, user)
	if err != nil {
		err = exceptions.CheckError(err, requestId)
		return
	}
	if rowsAffected != 1 {
		err = errors.New("rows affected not 1 when create user")
		err = exceptions.CheckError(err, requestId)
		return
	}

	registerUserResponse = models.ToRegisterUserResponse(user)
	return
}
