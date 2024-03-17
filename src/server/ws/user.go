package ws

import (
	"github.com/gorilla/websocket"
	"net/http"
	"server/tool"
	"server/util"
)

func init() {
	util.WsRegister("getUser", func(data interface{}, conn *websocket.Conn, r *http.Request) {

		tool.Debug(data.(string))
	})
}
