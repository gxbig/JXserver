package httpServer

import (
	"github.com/name5566/leaf/log"
	"math/rand"
	"net/http"
	"regexp"
	"server/msg"
	"server/redisClient"
	"server/util"
	"strconv"
	"strings"
	"time"
)

func init() {

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
	//获取邮箱
	var data = msg.UserEmail{}
	if err := util.Unpack(req, &data); err != nil {
		res := util.GetError(err.Error())
		util.Write(w, res)
		return
	}
	regex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+.[a-zA-Z]{2,}$`)
	match := regex.MatchString(data.Email)

	//格式错误
	if !match {
		res := util.GetError("邮箱地址错误！")
		util.Write(w, res)
		return
	}

	//重复校验
	user := msg.User{Email: data.Email}
	queryUser := user.QueryUser()

	if queryUser.Id != 0 {
		res := util.GetError("邮箱已注册！")
		util.Write(w, res)
		return
	}

	//发送验证码
	code := strconv.Itoa(int(rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(10000000)))
	err := util.SendMailCode(code, "18810994068@163.com")
	if err != nil {
		log.Error(err.Error())
		return
	}

	//保存redis

	redisClient.Rdb.Set(ctx, data.Email, code, time.Second*300)

	//转义
	codeArr := []rune(code)
	log.Debug(string(codeArr))

	res := util.GetSuccess("验证码发送成功！")

	util.Write(w, res)

}

// 注册
func registerHandler(w http.ResponseWriter, req *http.Request) {
	var data = &msg.User{}

	//获取数据
	if err := util.Unpack(req, data); err != nil {
		res := util.GetError(err.Error())
		util.Write(w, res)
		return
	}

	code := data.Code

	//校验邮箱
	regex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+.[a-zA-Z]{2,}$`)
	match := regex.MatchString(data.Email)

	//格式错误
	if !match {
		res := util.GetError("邮箱地址错误！")
		util.Write(w, res)
		return
	}
	//校验8位密码
	LoginPW := strings.TrimSpace(data.Pw)
	if strings.Count(LoginPW, "") < 8 {
		res := util.GetError("请输入最少八位数密码！")
		util.Write(w, res)
		return
	}
	data.Pw = LoginPW
	//校验验证码
	result := redisClient.Rdb.Get(ctx, data.Email).Val()

	if result != code {
		res := util.GetError("请输入正确的验证码")
		util.Write(w, res)
		return
	}
	data.Account = util.RandStringBytesMaskSrcUnsafe(12)

	//创建sessionsId
	sessionId := util.RandStringBytesMaskSrcUnsafe(24)
	//用户信息插入数据库
	err1 := data.RegisterInsetUser()
	if err1 != nil {
		res := util.GetError(err1.Error())
		util.Write(w, res)
		return
	}
	//

	//用户信息放到内存或redis
	util.SetSessionIdUser(sessionId, data)

	//fmt.Fprintf(w, "Search:%+v\n", data)
	res := util.GetSuccess(sessionId)
	util.Write(w, res)

}

// 登录
func loginHandler(w http.ResponseWriter, r *http.Request) {

	var data = &msg.User{}

	//获取数据
	if err := util.Unpack(r, data); err != nil {
		res := util.GetError(err.Error())
		util.Write(w, res)
		return
	}

	//校验邮箱
	regex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+.[a-zA-Z]{2,}$`)
	match := regex.MatchString(data.Email)

	//格式错误
	if !match {
		res := util.GetError("邮箱地址错误！")
		util.Write(w, res)
		return
	}
	//校验8位密码
	LoginPW := strings.TrimSpace(data.Pw)
	if strings.Count(LoginPW, "") < 8 {
		res := util.GetError("请输入最少八位数密码！")
		util.Write(w, res)
		return
	}

	//创建sessionsId
	sessionId := util.RandStringBytesMaskSrcUnsafe(12)
	//数据库查询用户信息
	//users, err := sqlClient.QueryUser(data)

	//if err != nil {
	//	res := util.GetError(err.Error())
	//	util.Write(w,res)
	//	return
	//}

	////用户信息放到内存或redis
	//util.SetSessionIdUser(sessionId, users[0])

	//fmt.Fprintf(w, "Search:%+v\n", data)
	res := util.GetSuccess(sessionId)

	util.Write(w, res)
}

// 重置密码
func resetPasswordHandler(w http.ResponseWriter, r *http.Request) {
	var data = &msg.User{}

	//获取数据
	if err := util.Unpack(r, data); err != nil {
		res := util.GetError(err.Error())
		util.Write(w, res)
		return
	}

	code := data.Code

	//校验邮箱
	regex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+.[a-zA-Z]{2,}$`)
	match := regex.MatchString(data.Email)

	//格式错误
	if !match {
		res := util.GetError("邮箱地址错误！")
		util.Write(w, res)
		return
	}
	//校验8位密码
	LoginPW := strings.TrimSpace(data.Pw)
	if strings.Count(LoginPW, "") < 8 {
		res := util.GetError("请输入最少八位数密码！")
		util.Write(w, res)
		return
	}

	//校验验证码
	result, err := redisClient.Rdb.Get(ctx, data.Email).Result()
	if err != nil {
		res := util.GetError(err.Error())
		util.Write(w, res)
		return
	}
	if result != code {
		res := util.GetError("请输入正确的验证码")
		util.Write(w, res)
		return
	}

	//用户信息插入数据库
	//id, err := sqlClient.RegisterInsetUser(data)
	//data.Id = id
	//if err != nil {
	//	res := util.GetError(err.Error())
	//	util.Write(w,res)
	//	return
	//}

	//fmt.Fprintf(w, "Search:%+v\n", data)
	res := util.GetSuccess("重置密码成功！")

	util.Write(w, res)
}
