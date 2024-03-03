package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"reflect"
	"server/msg"
)

var ag []gate.Agent

type MsgHandler func([]interface{})

func handleMsg(m interface{}, handler MsgHandler) {

	skeleton.RegisterChanRPC(reflect.TypeOf(m), func(args []interface{}) {
		m := args[0].(*msg.UserLogin)
		log.Debug("Test login %v", m.LoginName)
		// 消息的发送者
		a := args[1].(gate.Agent)
		ag = append(ag, a)

		a.WriteMsg(&msg.Exceptional{
			Status: "1",
		})
		ag[0].WriteMsg(&msg.Exceptional{
			Status: string(rune(len(ag))),
		})

		if a.UserData() == nil {

			return
		} else {
			log.Debug(a.UserData().(string))
		}

		// 1 查询数据库--判断用户是不是合法
		// 2 如果数据库返回查询正确--保存到缓存或者内存
		// 输出收到的消息的内容
		//log.Debug("Test login %v", m.LoginName)
		// 给发送者回应一个 Test 消息

		handler(args)
	})
}

func init() {
	handleMsg(&msg.UserLogin{}, handleTest)
}

// 消息处理
func handleTest(args []interface{}) {
	// 收到的 Test 消息
	m := args[0].(*msg.UserLogin)
	// 消息的发送者
	a := args[1].(gate.Agent)

	a.SetUserData(m.LoginName)
	// 1 查询数据库--判断用户是不是合法
	// 2 如果数据库返回查询正确--保存到缓存或者内存
	// 输出收到的消息的内容
	log.Debug("Test login %v", m.LoginName)
	// 给发送者回应一个 Test 消息
	a.WriteMsg(&msg.UserLogin{
		LoginName: "client",
	})
}
