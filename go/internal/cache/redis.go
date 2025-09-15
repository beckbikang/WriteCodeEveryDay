package cache

import (
	"context"
	"time"

	"github.com/beckbikang/WriteCodeEveryDay/global"
	"github.com/go-redis/redis/v8"
)

var (
	RedisClient *redis.Client
)

func SetupRedis() error {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:         global.RedisSetting.Addr,
		Password:     global.RedisSetting.Password,
		DB:           global.RedisSetting.DB,
		PoolSize:     global.RedisSetting.PoolSize,
		MinIdleConns: global.RedisSetting.MinIdleConns,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := RedisClient.Ping(ctx).Result()
	return err
}
