package services

import (
	"context"
	"errors"
	"project-user/commons/exceptions"
	"project-user/commons/utils"
	"project-user/features/login/repositories"

	"github.com/redis/go-redis/v9"
)

type LogoutService interface {
	Logout(ctx context.Context, requestId string, sessionId string) (err error)
}

type LogoutServiceImplementation struct {
	RedisUtil       utils.RedisUtil
	RedisRepository repositories.RedisRepository
}

func NewLogoutService(redisUtil utils.RedisUtil, redisRepository repositories.RedisRepository) LogoutService {
	return &LogoutServiceImplementation{
		RedisUtil:       redisUtil,
		RedisRepository: redisRepository,
	}
}

func (service *LogoutServiceImplementation) Logout(ctx context.Context, requestId string, sessionId string) (err error) {
	rowsAffected, err := service.RedisRepository.Del(service.RedisUtil.GetClient(), ctx, sessionId)
	if err != nil && err != redis.Nil {
		err = exceptions.CheckError(err, requestId)
		return
	}
	if rowsAffected != 1 {
		err = errors.New("rows affected not 1 when delete data to redis")
		err = exceptions.CheckError(err, requestId)
		return
	}
	return
}
