package gate

import (
	"server/login"
	"server/msg"
)

func init() {
	msg.Processor.SetRouter(&msg.UserLogin{}, login.ChanRPC)
}
