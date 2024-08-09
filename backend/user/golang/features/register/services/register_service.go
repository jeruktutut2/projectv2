package services

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"project-user/commons/exceptions"
	"project-user/commons/helpers"
	"project-user/commons/utils"
	"project-user/features/register/models"
	"project-user/features/register/repositories"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type RegisterService interface {
	Register(ctx context.Context, requestId string, nowUnixMili int64, registerUserRequest models.RegisterUserRequest) (registerUserResponse models.RegisterUserResponse, err error)
}

type RegisterServiceImplementation struct {
	MysqlUtil      utils.MysqlUtil
	Validate       *validator.Validate
	BcryptHelper   helpers.BcryptHelper
	UserRepository repositories.UserRepository
}

func NewRegisterService(mysqlUtil utils.MysqlUtil, validate *validator.Validate, bcryptHelper helpers.BcryptHelper, userRepository repositories.UserRepository) RegisterService {
	return &RegisterServiceImplementation{
		MysqlUtil:      mysqlUtil,
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

	numberOfUser, err := service.UserRepository.CountByUsername(service.MysqlUtil.GetDb(), ctx, registerUserRequest.Username)
	if err != nil && err != sql.ErrNoRows {
		err = exceptions.CheckError(err, requestId)
		return
	}
	if numberOfUser > 0 {
		err = errors.New("username already exists")
		err = exceptions.ToResponseException(err, requestId, http.StatusBadRequest, "username already exists")
		return
	}

	numberOfUser, err = service.UserRepository.CountByEmail(service.MysqlUtil.GetDb(), ctx, registerUserRequest.Email)
	if err != nil && err != sql.ErrNoRows {
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
	user.Username = sql.NullString{Valid: true, String: registerUserRequest.Username}
	user.Email = sql.NullString{Valid: true, String: registerUserRequest.Email}
	user.Password = sql.NullString{Valid: true, String: string(hashedPassword)}
	user.Utc = sql.NullString{Valid: true, String: registerUserRequest.Utc}
	user.CreatedAt = sql.NullInt64{Valid: true, Int64: nowUnixMili}
	rowsAffected, err := service.UserRepository.Create(service.MysqlUtil.GetDb(), ctx, user)
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
