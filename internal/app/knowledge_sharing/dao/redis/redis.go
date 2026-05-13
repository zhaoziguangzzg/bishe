package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func DaoInitRedis(addr, password string, db int) {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     addr,     // Redis 地址
		Password: password, // 无密码
		DB:       db,       // 默认 DB
	})

}

func SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	return RedisClient.SetNX(ctx, key, value, expiration).Result()
}

func Del(ctx context.Context, keys ...string) (int64, error) {
	return RedisClient.Del(ctx, keys...).Result()
}

func Expire(ctx context.Context, key string, expiration time.Duration) (bool, error) {
	return RedisClient.Expire(ctx, key, expiration).Result()
}
