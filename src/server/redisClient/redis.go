package redisClient

import (
	"github.com/go-redis/redis/v8"
	"server/conf"
	"server/tool"
)

var Rdb *redis.Client

// 创建redis客户端
func newClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     conf.Server.RedisAddr, // redis地址
		Password: conf.RedisPassword,    // 密码
		DB:       0,                     // 使用默认数据库
	})
	tool.Release("初始化redis连接成功!")
	return client
}
func InitRedisClient() {
	tool.Release("开始初始化redis连接!")
	Rdb = newClient()
}
