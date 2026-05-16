package redis

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"
)

// 加锁
func Lock(ctx context.Context, key string, expiration time.Duration) (string, bool, error) {
	// 生成唯一ID
	lockValue := fmt.Sprintf("%d-%d", os.Getpid(), time.Now().UnixNano())

	ok, err := RedisClient.SetNX(ctx, key, lockValue, expiration).Result()
	return lockValue, ok, err
}

// Unlock 解锁
func Unlock(ctx context.Context, key string, value string) error {
	// Lua 脚本保证：判断 + 删除 是原子操作
	script := `
		if redis.call("get", KEYS[1]) == ARGV[1] then
			return redis.call("del", KEYS[1])
		else
			return 0
		end
	`
	// 执行脚本
	res, err := RedisClient.Eval(ctx, script, []string{key}, value).Result()
	if err != nil {
		return err
	}

	// res == 1 表示删除成功
	if res.(int64) != 1 {
		return errors.New("解锁失败：锁不属于当前请求或已过期")
	}
	return nil
}
