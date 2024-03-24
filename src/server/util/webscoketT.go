package util

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"io"
	"net/http"
	"net/url"
	"server/conf"
	"server/msg"
	"server/tool"
	"sync"
)

// WebsocketWData ws发送的数据
type WebsocketWData struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

// WebsocketRData ws接收到的数据
type WebsocketRData struct {
	Type string           `json:"type"`
	Data *json.RawMessage `json:"data"`
}

type sessionId struct {
	SessionId string `json:"sessionId"`
}

// WsRequestResults ws 请求结果
type WsRequestResults struct {
	Code    string           `json:"code"`
	Message string           `json:"message"`
	Data    *json.RawMessage `json:"data"`
}

// UserConnection ws 链接数据
type UserConnection struct {
	Conn      *websocket.Conn `json:"conn"`
	User      *msg.User       `json:"user"`
	UuId      string          `json:"uuId"`
	R         *http.Request   `json:"r"`
	SessionId string          `json:"sessionId"`
}

// UserConnections 全部用户对应消息类型
type UserConnections map[int]*UserConnection

// RegisterHandel //消息处理
type RegisterHandel func(*json.RawMessage, *UserConnection, UserConnections)

// 路由数据
var messageRouter = map[string]RegisterHandel{}

// ConnectionsMutex 连接数据互斥锁
var ConnectionsMutex sync.Mutex

// 链接数据
var userConnections = make(UserConnections, 100)

// WsRegister 注册ws路由处理事件
func WsRegister(router string, handle RegisterHandel) {
	messageRouter[router] = handle
}

var upgrade = websocket.Upgrader{
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
func init() {
	//发送心跳
	WsRegister("heartbeat", func(data *json.RawMessage, userCon *UserConnection, userConnections UserConnections) {
		res, _ := json.Marshal(&WebsocketWData{Type: "heartbeat"})
		if err := userCon.Conn.WriteMessage(websocket.BinaryMessage, res); err != nil {
			tool.Error("heartbeat 发送:id=%s 消息=%s 错误=%s", userCon.User.Id, res, err.Error())
		}
		defer func() {
			if r := recover(); r != nil {
				tool.Error("heartbeat 异常：", r)
			}
		}()
	})
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// 升级HTTP连接为WebSocket连接
	conn, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		tool.Debug(err.Error())
		return
	}
	uuId := RandStringBytesMaskSrcUnsafe(24)
	var data = sessionId{}
	if unpackErr := Unpack(r, &data); unpackErr != nil {
		tool.Error("ws客户端连接失败：%s:%s", r.RemoteAddr, unpackErr.Error())
		return
	}
	//获取当前登陆人
	user := GetSessionIdUser(data.SessionId)
	//ws链接存保存
	var userConnection = &UserConnection{
		Conn:      conn,
		UuId:      uuId,
		R:         r,
		SessionId: data.SessionId,
	}
	if user == nil {
		if user = userConnection.HttpGetUser(); user == nil {
			return
		}
	}
	userConnection.User = user

	//单点登录
	//存在当前用户的上一个连接，发送单点断开消息，再断开
	if userConnections[user.Id] != nil {
		res, _ := json.Marshal(&WebsocketWData{Type: "/singClose"})
		userCon := userConnections[user.Id]
		if err := userCon.Conn.WriteMessage(websocket.TextMessage, res); err != nil {
			tool.Error("新uuid:%s与旧uuid:%s  重复登录断开连接：id=%s 消息=%s 错误=%s", uuId, userConnections[user.Id].UuId, userCon.User.Id, res, err.Error())
		} else {
			tool.Debug("新uuid:%s与旧uuid:%s  重复登录断开连接：:id=%s 消息=%s", uuId, userConnections[user.Id].UuId, userCon.User.Id, res)
		}
		_ = userCon.Conn.Close()
	}
	ConnectionsMutex.Lock()
	userConnections[user.Id] = userConnection
	ConnectionsMutex.Unlock()
	tool.Debug("uuid:%s ws客户端连接成功：地址=%s id=%d email=%s  姓名=%s", uuId, r.RemoteAddr, user.Id, user.Email, user.Name)

	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			tool.Error("uuid:%s ws客户端连接关闭失败：地址=%s id=%d email=%s  姓名=%s", uuId, r.RemoteAddr, user.Id, user.Email, user.Name)
		}
	}(conn)

	// 处理WebSocket连接
	for {
		// 读取消息
		messageType, p, err := conn.ReadMessage()
		//异常关闭
		if err != nil {
			ConnectionsMutex.Lock()
			//删除属于当前保存的连接
			if userConnections[user.Id] != nil && uuId == userConnections[user.Id].UuId {
				delete(userConnections, user.Id)
			}
			ConnectionsMutex.Unlock()
			//正常关闭
			if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				tool.Debug("uuid:%s ws客户端连接关闭：地址=%s id=%d email=%s  姓名=%s", uuId, r.RemoteAddr, user.Id, user.Email, user.Name)
				return
			}
			//其他关闭
			tool.Error("uid:%s 其他错误，ws客户端连接关闭：地址=%s id=%d email=%s  错误=%s", uuId, r.RemoteAddr, user.Id, user.Email, err.Error())
			return
		}

		var WebsocketRData = &WebsocketRData{}
		unmarshalError := json.Unmarshal(p, WebsocketRData)
		if unmarshalError != nil {
			tool.Error("%s 接收消息:id=%d %s:%s:%s", WebsocketRData.Type, user.Id, r.RemoteAddr, p, unmarshalError.Error())
		} else {
			if WebsocketRData.Type != "heartbeat" { //心跳不加日志
				tool.Debug("%s 接收消息:id=%d %d %s %s", WebsocketRData.Type, user.Id, messageType, r.RemoteAddr, p)
			}
			//转到对应的路由 //注册的路由解析为struct
			if messageRouter[WebsocketRData.Type] == nil {
				tool.Error("ws消息路由未注册: type=%s  id=%d  %s %s", WebsocketRData.Type, user.Id, r.RemoteAddr, p)
			} else {
				//增加用户信息和连接
				messageRouter[WebsocketRData.Type](WebsocketRData.Data, userConnection, userConnections)
			}
		}
	}
}

// HttpGetUser 通过sessionId获取用户信息
func (userCon *UserConnection) HttpGetUser() *msg.User {
	// 目标URL
	targetURL, err := url.Parse(conf.Server.HttpClientAddr + "/getUser")
	if err != nil {
		tool.Error("http获取用户信息失败:" + err.Error())
		return nil
	}
	// 创建请求
	request := &http.Request{
		Method: "GET",
		URL:    targetURL,
		Header: http.Header{
			"User-Agent": []string{"MyClient/0.1"},     // 设置User-Agent请求头
			"Accept":     []string{"application/json"}, // 设置Accept请求头
			"Token":      []string{userCon.SessionId},  // 设置Accept请求头
		},
	}

	// 发送请求
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		tool.Error("http获取用户-获取请求:" + err.Error())
		return nil
	}

	// 读取响应体
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			tool.Error("http获取用户-关闭读取请求异常:" + err.Error())
		}
	}(response.Body)
	body, err := io.ReadAll(response.Body)
	if err != nil {
		tool.Error("http获取用户-读取请求:" + err.Error())
		return nil
	}

	result := &WsRequestResults{}
	err = json.Unmarshal(body, result)
	//
	userData := &msg.User{}
	if result.Code == "200" {

		//取内部数据
		err = json.Unmarshal(*result.Data, userData)
		SetSessionIdUser(userCon.SessionId, userData)
		return userData
	} else if result.Code == "502" {
		res, _ := json.Marshal(&WebsocketWData{Type: "/unLogin"})
		err = userCon.Conn.WriteMessage(websocket.TextMessage, res)
		if err != nil {
			tool.Error("未登录:%s:%s", res, err.Error())
		} else {
			tool.Debug("未登录:%s:%s", userCon.R.RemoteAddr, res)
		}
		return nil
	} else {
		res, _ := json.Marshal(&WebsocketWData{Type: "error"})
		err = userCon.Conn.WriteMessage(websocket.TextMessage, res)
		if err != nil {
			tool.Error("获取用户异常:%s %s", result.Message, err.Error())
		} else {
			tool.Error("获取用户异常:%s", result.Message)
		}
		return nil
	}
}
