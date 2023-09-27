package redisClient

import (
	"github.com/go-redis/redis/v8"
	"github.com/name5566/leaf/log"
)

var RedisClient *redis.Client

// 创建redis客户端
func newClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // redis地址
		Password: "",               // 密码
		DB:       0,                // 使用默认数据库
	})
	return client
}
func InitRedisClient() {
	log.Release("开始初始化redis连接!")
	RedisClient = newClient()
}
