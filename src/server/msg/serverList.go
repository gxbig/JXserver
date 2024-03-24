package msg

import (
	"gorm.io/gorm"
	"server/sqlClient"
	"time"
)

// 服务表
type Server struct {
	Id         int    `json:"id" gorm:"column:id"`
	Address    string `json:"address"  gorm:"column:address"` //服务器地址
	Code       int    `json:"code"  gorm:"column:code"`       //代码
	Status     int    `json:"status"  gorm:"column:status"`   //状态 1、流畅 2、爆满 3、维护
	Max        int    `json:"max"  gorm:"column:max"`         //最大用户数
	ServerName string `json:"serverName"`                     //服务器名称
	Recommend  int    `json:"recommend"`                      //推荐服务器
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func (server *Server) TableName() string {
	return "servers"
}

// 获取用户服务器数据
func (server *Server) CodeQueryServers(codes []int) []Server {
	var serves []Server
	if len(codes) > 0 {
		sqlClient.DB.Where("code in (?)", codes).Order("code").Find(&serves)
	}
	return serves
}

// 获取所有服务器数据
func (server *Server) QueryServers() []Server {
	var serves []Server
	sqlClient.DB.Order("code").Find(&serves)
	return serves
}

// 更新密码数据
func (server *Server) UpdateUserPw() error {
	result := sqlClient.DB.Model(server).Where("CODE=?", server.Code).Updates(server)
	return result.Error
}
