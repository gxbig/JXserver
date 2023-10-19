package util

import (
	"context"
	"encoding/json"
	"server/msg"
	"server/redisClient"
)

var ctx = context.Background()
var SessionIdUser = make(map[string]*msg.User, 100)

// 保存session对应的user
func SetSessionIdUser(SessionId string, user *msg.User) {
	SessionIdUser[SessionId] = user
	jsonUser, _ := json.Marshal(*user)
	redisClient.Rdb.HSet(ctx, "SessionId", SessionId, string(jsonUser))
}

// 获取session对应的user
func GetSessionIdUser(SessionId string) *msg.User {
	//先从内存获取
	user := SessionIdUser[SessionId]
	if user != nil {
		return user
	}

	//从redis获取
	rUserr := redisClient.Rdb.HGet(ctx, "SessionId", SessionId).Val()
	rUser := &msg.User{}
	_ = json.Unmarshal([]byte(rUserr), rUser)
	if rUser != nil {
		return rUser
	}
	return nil
}
