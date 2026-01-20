package service

import "bishe/internal/app/knowledge_sharing/dao/redis"

// 初始化 Redis 客户端
func ServiceInitRedis(addr, password string, db int) {
	redis.DaoInitRedis(addr, password, db)
}
