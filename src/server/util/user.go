package util

import (
	"context"
	"encoding/json"
	"server/msg"
	"server/redisClient"
	"strconv"
)

var ctx = context.Background()
var SessionIdUser = make(map[string]*msg.User, 100)
var IdSessionId = make(map[int]string, 100)

// 保存session对应的user
func SetSessionIdUser(SessionId string, user *msg.User) {
	id := strconv.Itoa(user.Id)
	//先去掉保存的数据
	DelSessionUserById(id)
	//再增加
	SessionIdUser[SessionId] = user
	IdSessionId[user.Id] = SessionId
	jsonUser, _ := json.Marshal(*user)
	redisClient.Rdb.HSet(ctx, "SessionIdToUser", SessionId, string(jsonUser))
	redisClient.Rdb.HSet(ctx, "userIdToSessionId", id, SessionId)
}

// 获取session对应的user
func GetSessionIdUser(SessionId string) *msg.User {
	//先从内存获取
	user := SessionIdUser[SessionId]
	if user != nil {
		return user
	}

	//从redis获取
	rUserr := redisClient.Rdb.HGet(ctx, "SessionIdToUser", SessionId).Val()
	rUser := &msg.User{}
	_ = json.Unmarshal([]byte(rUserr), rUser)
	if rUser != nil {
		return rUser
	}
	return nil
}

// 删除 user 通过id
func DelSessionUserById(id string) {
	SessionId := redisClient.Rdb.HGet(ctx, "sessionId", id).Val()

	//删除sessionId对应的userId
	redisClient.Rdb.HDel(ctx, "userIdToSessionId", id)
	//删除sessionId对应的user
	redisClient.Rdb.HDel(ctx, "SessionIdToUser", SessionId)

}
