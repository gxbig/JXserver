package loginServer

import (
	"context"
	"github.com/name5566/leaf/log"
	"net/http"
	"server/msg"
	"server/redisClient"
	"server/util"
	"time"
)

var Mux *http.ServeMux

func init() {
	Mux = http.NewServeMux()
	login := util.HandlerFunc(loginHandler)
	Mux.Handle("/login", login)

	register := util.HandlerFunc(registerHandler)
	Mux.Handle("/register", register)
	code := util.HandlerFunc(getCodeHandler)
	Mux.Handle("/getCode", code)
	resetPassword := util.HandlerFunc(resetPasswordHandler)
	Mux.Handle("/resetPassword", resetPassword)
}

// 获取注册验证码
func getCodeHandler(w http.ResponseWriter, req *http.Request) {

	code := util.RandStringBytesMaskSrcUnsafe(4)
	util.SendMail(code, "18810994068@163.com")
	//保存redis
	ctx := context.Background()
	redisClient.Rdb.Set(ctx, code, code, time.Second*300)

	//转义
	codeArr := []rune(code)
	log.Debug(string(codeArr))
	var recodeArr = make([]rune, 4)
	for index, val := range codeArr {
		recodeArr[index] = val + rune(index) + 1
	}

	res, err := util.GetSuccess(string(recodeArr))

	if err != nil {
		res = util.GetError(code)
	}

	w.Write(res)

}

// 注册
func registerHandler(w http.ResponseWriter, req *http.Request) {
	var data = msg.UserRegister{}

	if err := util.Unpack(req, &data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//fmt.Fprintf(w, "Search:%+v\n", data)
	res, err := util.GetSuccess(data)
	if err != nil {
		res = util.GetError("")
	}

	w.Write(res)

}

// 登录
func loginHandler(w http.ResponseWriter, r *http.Request) {

	_, err := w.Write([]byte("the loginHandler"))
	if err != nil {
		return
	}
}

// 重置密码
func resetPasswordHandler(w http.ResponseWriter, r *http.Request) {

	_, err := w.Write([]byte("the resetPasswordHandler"))
	if err != nil {
		return
	}
}
