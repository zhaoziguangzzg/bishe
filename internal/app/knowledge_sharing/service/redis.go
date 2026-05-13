package service

import (
	"context"
	"time"

	"bishe/internal/app/knowledge_sharing/dao/redis"
)

// 初始化 Redis 客户端

func ServiceInitRedis(addr, password string, db int) {
	redis.DaoInitRedis(addr, password, db)
}

func AcquireLock(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	return redis.SetNX(ctx, key, value, expiration)
}

func ReleaseLock(ctx context.Context, keys ...string) (int64, error) {
	return redis.Del(ctx, keys...)
}

func SetLockExpire(ctx context.Context, key string, expiration time.Duration) (bool, error) {
	return redis.Expire(ctx, key, expiration)
}
