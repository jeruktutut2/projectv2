package repositories

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisRepository interface {
	Del(client *redis.Client, ctx context.Context, key string) (result int64, err error)
}

type RedisRepositoryImplementation struct {
}

func NewRedisRepository() RedisRepository {
	return &RedisRepositoryImplementation{}
}

func (repository *RedisRepositoryImplementation) Del(client *redis.Client, ctx context.Context, key string) (result int64, err error) {
	return client.Del(ctx, key).Result()
}
