package service

import (
	"bishe/internal/app/knowledge_sharing/dao/redis"
	"time"
)

func ServiceInitRedis(addr, password string, db int) {
	redis.DaoInitRedis(addr, password, db)
}

// 加锁
func Lock(key string, expiration time.Duration) (string, bool, error) {
	return redis.Lock(key, expiration)
}

// Unlock 解锁
func Unlock(key string, value string) error {
	return redis.Unlock(key, value)
}
