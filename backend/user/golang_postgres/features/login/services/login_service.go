package services

import (
	"context"
	"encoding/json"
	"errors"
	"golang-postgres/commons/exceptions"
	"golang-postgres/commons/helpers"
	"golang-postgres/commons/utils"
	"golang-postgres/features/login/models"
	"golang-postgres/features/login/repositories"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
)

type LoginService interface {
	Login(ctx context.Context, requestId string, sessionIdUser string, loginUserRequest models.LoginUserRequest) (sessionId string, err error)
}

type LoginServiceImplementation struct {
	PostgresUtil             utils.PostgresUtil
	RedisUtil                utils.RedisUtil
	Validate                 *validator.Validate
	UserRepository           repositories.UserRepository
	UserPermissionRepository repositories.UserPermissionRepository
	BcryptHelper             helpers.BcryptHelper
	RedisRepository          repositories.RedisRepository
	UuidHelper               helpers.UuidHelper
}

func NewLoginService(postgresUtil utils.PostgresUtil, redisUtil utils.RedisUtil, validate *validator.Validate, userRepository repositories.UserRepository, userPermissionRepository repositories.UserPermissionRepository, bcryptHelper helpers.BcryptHelper, redisRepository repositories.RedisRepository, uuidHelper helpers.UuidHelper) LoginService {
	return &LoginServiceImplementation{
		PostgresUtil:             postgresUtil,
		RedisUtil:                redisUtil,
		Validate:                 validate,
		UserRepository:           userRepository,
		UserPermissionRepository: userPermissionRepository,
		BcryptHelper:             bcryptHelper,
		RedisRepository:          redisRepository,
		UuidHelper:               uuidHelper,
	}
}

func (service *LoginServiceImplementation) Login(ctx context.Context, requestId string, sessionIdUser string, loginUserRequest models.LoginUserRequest) (sessionId string, err error) {
	if sessionIdUser != "" {
		var rowsAffected int64
		rowsAffected, err = service.RedisRepository.Del(service.RedisUtil.GetClient(), ctx, sessionIdUser)
		if err != nil {
			helpers.PrintLogToTerminal(err, requestId)
		} else if err == nil && rowsAffected != 1 {
			err = errors.New("rows affected not 1")
			helpers.PrintLogToTerminal(err, requestId)
		}
	}
	err = service.Validate.Struct(loginUserRequest)
	if err != nil {
		validationResult := helpers.GetValidatorError(err, loginUserRequest)
		if validationResult != nil {
			err = exceptions.ToResponseExceptionRequestValidation(requestId, validationResult)
			return
		}
	}

	user, err := service.UserRepository.FindByEmail(service.PostgresUtil.GetPool(), ctx, loginUserRequest.Email)
	if err != nil && err != pgx.ErrNoRows {
		err = exceptions.CheckError(err, requestId)
		return
	} else if err != nil && err == pgx.ErrNoRows {
		err = errors.New("wrong email or password")
		err = exceptions.ToResponseException(err, requestId, http.StatusBadRequest, "wrong email or password")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password.String), []byte(loginUserRequest.Password))
	if err != nil {
		err = errors.New("wrong email or password")
		err = exceptions.ToResponseException(err, requestId, http.StatusBadRequest, "wrong email or password")
		return
	}

	userPermissions, err := service.UserPermissionRepository.FindByUserId(service.PostgresUtil.GetPool(), ctx, user.Id.Int32)
	if err != nil {
		err = exceptions.CheckError(err, requestId)
		return
	}
	var idPermissions []int32
	for _, userPermission := range userPermissions {
		idPermissions = append(idPermissions, userPermission.PermissionId.Int32)
	}

	sessionId = service.UuidHelper.String()
	sessionValue := make(map[string]interface{})
	sessionValue["id"] = user.Id.Int32
	sessionValue["username"] = user.Username.String
	sessionValue["email"] = user.Email.String
	sessionValue["idPermissions"] = idPermissions
	sessionByte, err := json.Marshal(sessionValue)
	if err != nil {
		err = exceptions.CheckError(err, requestId)
		return
	}
	session := string(sessionByte)

	_, err = service.RedisRepository.Set(service.RedisUtil.GetClient(), ctx, sessionId, session, 0)
	if err != nil && err != redis.Nil {
		err = exceptions.CheckError(err, requestId)
		return
	}
	return
}
