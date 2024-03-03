package httpServer

import (
	"math/rand"
	"net/http"
	"regexp"
	"server/msg"
	"server/redisClient"
	"server/tool"
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
		res := util.GetError(req, err.Error())
		util.Write(w, res)
		return
	}
	regex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+.[a-zA-Z]{2,}$`)
	match := regex.MatchString(data.Email)

	//格式错误
	if !match {
		res := util.GetError(req, "邮箱地址错误！")
		util.Write(w, res)
		return
	}

	//重复校验
	//user := msg.User{Email: data.Email}
	//queryUser := user.QueryUser()
	//
	//if queryUser.Id != 0 {
	//	res := util.GetError("邮箱已注册！")
	//	util.Write(w, res)
	//	return
	//}

	//发送验证码
	code := strconv.Itoa(int(rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(10000000)))
	//转义
	codeArr := []rune(code)
	tool.Debug(string(codeArr))
	//保存redis
	redisClient.Rdb.Set(ctx, data.Email, code, time.Second*300)
	err := util.SendMailCode(code, "18810994068@163.com")
	if err != nil {
		util.Write(w, util.GetError(req, "验证码发送失败！"))
		return
	}

	res := util.GetSuccess(req, "验证码发送成功！")

	util.Write(w, res)

}

// 注册
func registerHandler(w http.ResponseWriter, req *http.Request) {
	var data = &msg.User{}

	//获取数据
	if err := util.Unpack(req, data); err != nil {
		res := util.GetError(req, err.Error())
		util.Write(w, res)
		return
	}

	code := data.Code

	//校验邮箱
	regex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+.[a-zA-Z]{2,}$`)
	match := regex.MatchString(data.Email)

	//格式错误
	if !match {
		res := util.GetError(req, "邮箱地址错误！")
		util.Write(w, res)
		return
	}
	//重复校验
	user := msg.User{Email: data.Email}
	queryUser := user.QueryUser()

	if queryUser.Id != 0 {
		res := util.GetError(req, "邮箱已注册！")
		util.Write(w, res)
		return
	}

	//校验8位密码
	LoginPW := strings.TrimSpace(data.Pw)
	if strings.Count(LoginPW, "") < 8 {
		res := util.GetError(req, "请输入最少八位数密码！")
		util.Write(w, res)
		return
	}
	data.Pw = LoginPW
	//校验验证码
	result := redisClient.Rdb.Get(ctx, data.Email).Val()

	if result != code {
		res := util.GetError(req, "请输入正确的验证码")
		util.Write(w, res)
		return
	}
	data.Account = util.RandStringBytesMaskSrcUnsafe(12)

	//创建sessionsId
	sessionId := util.RandStringBytesMaskSrcUnsafe(24)
	//用户信息插入数据库
	err1 := data.RegisterInsetUser()
	if err1 != nil {
		res := util.GetError(req, err1.Error())
		util.Write(w, res)
		return
	}
	//
	//删除redis
	redisClient.Rdb.Del(ctx, data.Email, data.Code)
	//用户信息放到内存或redis
	util.SetSessionIdUser(sessionId, data)

	res := util.GetSuccess(req, sessionId)
	util.Write(w, res)

}

// 登录
func loginHandler(w http.ResponseWriter, r *http.Request) {

	var data = &msg.User{}

	//获取数据
	if err := util.Unpack(r, data); err != nil {
		res := util.GetError(r, err.Error())
		util.Write(w, res)
		return
	}

	//校验邮箱
	regex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+.[a-zA-Z]{2,}$`)
	match := regex.MatchString(data.Email)

	//格式错误
	if !match {
		res := util.GetError(r, "邮箱地址错误！")
		util.Write(w, res)
		return
	}
	//邮箱是否存在
	user := msg.User{Email: data.Email}
	queryUser := user.QueryUser()

	if queryUser.Id == 0 {
		res := util.GetError(r, "邮箱未注册！")
		util.Write(w, res)
		return
	}

	//校验8位密码
	LoginPW := strings.TrimSpace(data.Pw)
	if strings.Count(LoginPW, "") < 8 {
		res := util.GetError(r, "请输入最少八位数密码！")
		util.Write(w, res)
		return
	}

	if queryUser.Pw != LoginPW {
		res := util.GetError(r, "登录密码错误！")
		util.Write(w, res)
		return
	}

	//创建sessionsId
	sessionId := util.RandStringBytesMaskSrcUnsafe(24)

	//用户信息放到内存或redis
	util.SetSessionIdUser(sessionId, queryUser)

	res := util.GetSuccess(r, sessionId)

	util.Write(w, res)
}

// 重置密码
func resetPasswordHandler(w http.ResponseWriter, r *http.Request) {
	var data = &msg.User{}

	//获取数据
	if err := util.Unpack(r, data); err != nil {
		res := util.GetError(r, err.Error())
		util.Write(w, res)
		return
	}

	code := data.Code

	//校验邮箱
	regex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+.[a-zA-Z]{2,}$`)
	match := regex.MatchString(data.Email)

	//格式错误
	if !match {
		res := util.GetError(r, "邮箱地址错误！")
		util.Write(w, res)
		return
	}
	//校验8位密码
	LoginPW := strings.TrimSpace(data.Pw)
	if strings.Count(LoginPW, "") < 8 {
		res := util.GetError(r, "请输入最少八位数密码！")
		util.Write(w, res)
		return
	}

	//校验验证码
	result := redisClient.Rdb.Get(ctx, data.Email).Val()

	if result != code {
		res := util.GetError(r, "请输入正确的验证码")
		util.Write(w, res)
		return
	}

	//用户信息插入数据库
	err1 := data.UpdateUserPw()

	if err1 != nil {
		res := util.GetError(r, err1.Error())
		util.Write(w, res)
		return
	}
	//删除redis
	redisClient.Rdb.Del(ctx, data.Email, data.Code)
	//fmt.Fprintf(w, "Search:%+v\n", data)
	res := util.GetSuccess(r, "重置密码成功！")

	util.Write(w, res)
}
