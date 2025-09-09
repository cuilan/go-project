package rdb

import (
	redis "github.com/redis/go-redis/v9"
)

// redisClient 将在 redis 配置存在时被初始化
var redisClient *redis.Client

// GetRedis 是一个便捷的辅助函数，用于获取已初始化的客户端
func GetRedis() *redis.Client {
	if redisClient == nil {
		// 这种情况理论上不应该发生，因为模块化系统会确保它在使用前被初始化
		// 如果发生，说明在没有redis配置的情况下调用了它，或者初始化失败
		panic("redis client is not initialized")
	}
	return redisClient
}
