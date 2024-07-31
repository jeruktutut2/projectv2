package initialize

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

func SetDataRedis(client *redis.Client, ctx context.Context, key string, value interface{}, expiration time.Duration) {
	_, err := client.Set(ctx, key, value, expiration).Result()
	if err != nil {
		log.Fatalln("error when setting data redis:", err.Error())
	}
	log.Println("set data redis succedded")
}

func DelDataRedis(client *redis.Client, ctx context.Context, key string) {
	_, err := client.Del(ctx, key).Result()
	if err != nil {
		log.Fatalln("error when deleting data redis:", err.Error())
	}
	log.Println("delete data redis succeded")
}

func GetDataRedis(client *redis.Client, ctx context.Context, key string) (result string, err error) {
	result, err = client.Get(ctx, key).Result()
	if err != nil && err != redis.Nil {
		log.Fatalln("error when getting data redis:", err)
	}
	log.Println("get data redis succedded")
	return
}
