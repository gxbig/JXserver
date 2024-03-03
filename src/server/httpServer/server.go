package httpServer

import (
	"net/http"
	"server/msg"
	"server/util"
)

func init() {

	getServers := util.HandlerFunc(getServersHandler)
	Mux.Handle("/getServers", getServers)
	getUsedServers := util.HandlerFunc(getUsedServersHandler)
	Mux.Handle("/getUsedServers", getUsedServers)

}

// 获取全部服务器地址
func getServersHandler(w http.ResponseWriter, req *http.Request) {
	server := &msg.Server{}
	servers := server.QueryServers()
	res := util.GetSuccess(req, servers)

	util.Write(w, res)
}

// 获取常用服务器地址
func getUsedServersHandler(w http.ResponseWriter, req *http.Request) {
	//获取用户
	sessionId := req.Header.Get("Token")
	user := util.GetSessionIdUser(sessionId)

	// 获取角色
	userRoles := &msg.UserRole{}
	roles := userRoles.GetUserRolesByUserId(user.Id)
	var codes []int
	for _, val := range roles {
		codes = append(codes, val.ServerCode)
	}

	server := &msg.Server{}
	servers := server.CodeQueryServers(codes)
	res := util.GetSuccess(req, servers)

	util.Write(w, res)
}
