package ws

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"server/tool"
	"server/util"
)

func init() {

	util.WsRegister("/getUser", func(data *json.RawMessage, userCon *util.UserConnection, userConnections util.UserConnections) {
		//var user = &msg.User{}
		//err := json.Unmarshal(*data, user)
		//if err != nil {
		//	tool.Error("参数解析异常：type=%s 地址=%s id=%d email=%s  姓名=%s", "/getUser", userCon.R.RemoteAddr, userCon.User.Id, userCon.User.Email, userCon.User.Name)
		//	return
		//}

		res, _ := json.Marshal(&util.WebsocketWData{Type: "/setUser", Data: userCon.User})
		if err := userCon.Conn.WriteMessage(websocket.TextMessage, res); err != nil {
			tool.Error("/getUser发送:id=%s 消息=%s 错误=%s", userCon.User.Id, res, err.Error())
		} else {
			tool.Debug("/getUser发送:id=%s 消息=%s", userCon.User.Id, res)
		}
		defer func() {
			if r := recover(); r != nil {
				tool.Error("/getUser 异常：", r)
			}
		}()

	})

}
