package msg

// 用户登陆协议
type UserLogin struct {
	LoginName string `json:"loginName" ` // 用户名
	LoginPW   string `json:"loginPW"`    // 密码
	Code      string `json:"code"`       //验证码
}

// 注册协议
type UserRegister struct {
	LoginName string `json:"loginName"` // 用户名
	LoginPW   string `json:"loginPW"`   // 密码
	Code      string `json:"code"`      //验证码
	Phone     string `json:"phone"`     // 手机号
	Email     string `json:"email"`     // 邮箱
}

// 玩家的临时结构
// 玩家有角色的情况
type UserST struct {
	UID      string // 账号ID
	ServerID string // 服务器ID
	RoleUID  string // 角色UID
	RoleName string // 角色名字
	RoleLev  string // 角色等级
	Coin     string // 金币
	// 其他的暂时不做定义
}
