package util

import (
	"github.com/gorilla/websocket"
	"net/http"
	"server/tool"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func InitWebsocket() {
	// 创建HTTP服务器
	http.HandleFunc("/ws", handleWebSocket)
	tool.Debug("Server started on :3564")
	tool.Error(http.ListenAndServe(":3564", nil).Error())
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// 升级HTTP连接为WebSocket连接
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		tool.Debug(err.Error())
		return
	}
	defer conn.Close()

	// 处理WebSocket连接
	for {
		// 读取消息
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			tool.Debug(err.Error())
			return
		}
		tool.Debug("Received message:", string(p))

		// 发送消息
		err = conn.WriteMessage(messageType, []byte("Hello, world!"))
		if err != nil {
			tool.Debug(err.Error())
			return
		}
	}
}
