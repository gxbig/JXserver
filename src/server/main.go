package main

import (
	_ "context"
	"net/http"
	"os"
	"os/signal"
	"server/conf"
	"server/httpServer"
	"server/redisClient"
	_ "server/sqlClient"
	"server/threadPool"
	"server/tool"
	"server/util"
)

func startServer() {
	tool.Debug("启动http登录服务！")
	err := http.ListenAndServe("0.0.0.0:9000", httpServer.Mux)
	if err != nil {
		tool.Error(err.Error())
	}
}

//// var logger *log.Logger
//func initLog() {
//	log.New("debug", "./log", 3)
//}

func main() {
	//util.InitLog()
	//启动http登录服务！
	go startServer()
	//defer util.Logger.Close()
	//初始化协程持
	threadPool.InitPool(conf.PoolMaxNum)
	//初始化redis
	redisClient.InitRedisClient()
	defer func() {
		_ = redisClient.Rdb.Close()
		tool.Close()
	}()

	util.InitWebsocket()

	//user := &msg.User{Id: 1}
	//err1 := user.RegisterInsetUser()
	//if err1 != nil {
	//	log.Error(err1.Error())
	//}
	//err := user.DeleteUser()
	//if err != nil {
	//	log.Error(err.Error())
	//}
	//queryUser, err := user.QueryUser()
	//if err != nil {
	//	log.Error(err.Error())
	//}
	//log.Debug(strconv.Itoa(queryUser.Id))
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
	//lconf.LogLevel = conf.Server.LogLevel
	//lconf.LogPath = conf.Server.LogPath
	//lconf.LogFlag = conf.LogFlag
	//lconf.ConsolePort = conf.Server.ConsolePort
	//lconf.ProfilePath = conf.Server.ProfilePath
	//
	//leaf.Run(
	//	game.Module,
	//	gate.Module,
	//	login.Module,
	//)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	sig := <-c
	tool.Release("closing down (signal: %v)", sig)
}
