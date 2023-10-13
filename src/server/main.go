package main

import (
	_ "context"
	"github.com/name5566/leaf"
	lconf "github.com/name5566/leaf/conf"
	"github.com/name5566/leaf/log"
	"net/http"
	"server/conf"
	"server/game"
	"server/gate"
	"server/login"
	"server/loginServer"
	"server/redisClient"
	_ "server/sqlClient"
	"server/threadPool"
)

func startServer() {
	log.Debug("启动http登录服务！")
	http.ListenAndServe(":9000", loginServer.Mux)
}
func main() {
	//启动http登录服务！
	go startServer()

	//初始化协程持
	threadPool.InitPool(conf.PoolMaxNum)
	//初始化redis
	redisClient.InitRedisClient()
	defer func() {
		_ = redisClient.Rdb.Close()

	}()

	//val, err := redisClient.RedisClient.Get(ctx, "8").Result()
	//ctx := context.Background()
	//
	//val, err := redisClient.Rdb.Get(ctx, "10").Result()
	//switch {
	//case err == redis.Nil:
	//	log.Error("key does not exist")
	//case err != nil:
	//	log.Error("Get failed", err)
	//case val == "":
	//
	//	log.Debug("value is empty")
	//}
	//log.Debug(val, err)
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.Server.ConsolePort
	lconf.ProfilePath = conf.Server.ProfilePath

	leaf.Run(
		game.Module,
		gate.Module,
		login.Module,
	)

}
