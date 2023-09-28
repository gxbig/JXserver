package redisClient

import (
	"github.com/go-redis/redis/v8"
	"github.com/name5566/leaf/log"
	"server/conf"
)

var RedisClient *redis.Client

// 创建redis客户端
func newClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     conf.Server.RedisAddr, // redis地址
		Password: conf.RedisPassword,    // 密码
		DB:       0,                     // 使用默认数据库
	})
	log.Release("初始化redis连接成功!")
	return client
}
func InitRedisClient() {
	log.Release("开始初始化redis连接!")
	RedisClient = newClient()
}
