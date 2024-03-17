package util

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"net/http"
	"net/url"
	"server/conf"
	"server/msg"
	"server/tool"
)

type WebsocketData struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}
type sessionId struct {
	SessionId string `json:"sessionId"`
}

// //消息处理
type handel func(interface{}, *websocket.Conn, *http.Request)

var messageRouter = map[string]handel{}

// 注册ws路由处理事件
func WsRegister(router string, handle func(interface{}, *websocket.Conn, *http.Request)) {
	messageRouter[router] = handle
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func InitWebsocket() {
	// 创建HTTP服务器
	http.HandleFunc("/", handleWebSocket)
	tool.Debug("websocket启动:" + conf.Server.WSAddr)
	tool.Error(http.ListenAndServe(conf.Server.WSAddr, nil).Error())
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// 升级HTTP连接为WebSocket连接
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		tool.Debug(err.Error())
		return
	}
	var data = sessionId{}
	if unpackErr := Unpack(r, &data); unpackErr != nil {
		tool.Error("ws客户端连接失败：%s:%s", r.RemoteAddr, unpackErr.Error())
		return
	}
	//获取当前登陆人
	user := GetSessionIdUser(data.SessionId)
	// 目标URL
	targetURL, err := url.Parse(conf.Server.HttpClientAddr + "/getUser")
	if err != nil {
		tool.Error("http获取用户信息失败:" + err.Error())
		return
	}
	// 创建请求
	request := &http.Request{
		Method: "GET",
		URL:    targetURL,
		Header: http.Header{
			"User-Agent": []string{"MyClient/0.1"},     // 设置User-Agent请求头
			"Accept":     []string{"application/json"}, // 设置Accept请求头
			"Token":      []string{data.SessionId},     // 设置Accept请求头
		},
	}

	// 发送请求
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		tool.Error("http获取用户-获取请求:" + err.Error())
		return
	}

	// 读取响应体
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		tool.Error("http获取用户-读取请求:" + err.Error())
		return
	}

	result := &Results{}
	err = json.Unmarshal(body, result)

	if result.Code == "200" {
		tool.Debug("ws客户端连接成功：地址=%s id=%d email=%s  姓名=%s", r.RemoteAddr, user.Id, user.Email, user.Name)
	} else if result.Code == "502" {
		res, _ := json.Marshal(&WebsocketData{Type: "unLogin"})
		Writeerr := conn.WriteMessage(websocket.BinaryMessage, res)
		if Writeerr != nil {
			tool.Error("未登录:%s:%s", res, Writeerr.Error())
		} else {
			tool.Debug("未登录:%s:%s", r.RemoteAddr, res)
		}
		return
	} else {
		res, _ := json.Marshal(&WebsocketData{Type: "error"})
		Writeerr := conn.WriteMessage(websocket.BinaryMessage, res)
		if Writeerr != nil {
			tool.Error("获取用户异常:%s %s", result.Message, Writeerr.Error())
		} else {
			tool.Error("获取用户异常:%s", result.Message)
		}
		return
	}

	users := &msg.User{}
	err = json.Unmarshal(body, users)
	SetSessionIdUser(data.SessionId, users)

	defer conn.Close()

	// 处理WebSocket连接
	for {
		// 读取消息
		_, p, err := conn.ReadMessage()
		if err != nil {
			tool.Error("读取消息：%s:%s", r.RemoteAddr, err.Error())
			return
		}

		var websocketData = &WebsocketData{}
		unmarshalError := json.Unmarshal(p, websocketData)
		if unmarshalError != nil {
			tool.Error("接收消息:%s:%s:%s", r.RemoteAddr, p, unmarshalError.Error())
		} else {
			tool.Debug("接收消息:%s:%s", r.RemoteAddr, p)
			//转到对应的路由
			if messageRouter[websocketData.Type] == nil {
				tool.Error("ws消息路由未注册: type=%s : %s: %s", websocketData.Type, r.RemoteAddr, p)
			} else {
				messageRouter[websocketData.Type](websocketData.Data, conn, r)
			}
		}

		//res, _ := json.Marshal(&msg.User{
		//	Id:  1,
		//	Age: 5,
		//})
		//// 发送消息
		//Writeerr := conn.WriteMessage(messageType, res)
		//if Writeerr != nil {
		//	tool.Error("发送消息:%s:%s", res, Writeerr.Error())
		//} else {
		//	tool.Debug("发送消息:%s:%s", r.RemoteAddr, res)
		//}
	}
}
