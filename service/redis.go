package service

import (
	"bishe/dao/redis"
	"context"
	"time"
)

func ServiceInitRedis(addr, password string, db int) {
	redis.DaoInitRedis(addr, password, db)
}

// 加锁
func Lock(ctx context.Context, key string, expiration time.Duration) (string, bool, error) {
	return redis.Lock(ctx, key, expiration)
}

// Unlock 解锁
func Unlock(ctx context.Context, key string, value string) error {
	return redis.Unlock(ctx, key, value)
}
