package redis

import (
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
