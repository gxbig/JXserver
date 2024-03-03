package httpServer

import (
	"net/http"
	"server/msg"
	"server/util"
)

func init() {
	getUser := util.HandlerFunc(getUserHandler)
	Mux.Handle("/getUser", getUser)
}

// 获取人员
func getUserHandler(w http.ResponseWriter, r *http.Request) {
	//获取邮箱
	//用户信息放到内存或redis
	sessionId := r.Header.Get("Token")
	user := util.GetSessionIdUser(sessionId)
	user.Pw = ""
	user.Identity = ""
	res := util.GetSuccess(r, user)
	util.Write(w, res)

}

// 获取当前用户所有角色
func getUserRolesHandler(w http.ResponseWriter, r *http.Request) {

	//用户信息放到内存或redis
	sessionId := r.Header.Get("Token")
	user := util.GetSessionIdUser(sessionId)

	//获取所有角色
	role := &msg.UserRole{UserId: user.Id}
	userRoles := role.GetUserRoles()
	res := util.GetSuccess(r, userRoles)

	util.Write(w, res)

}
