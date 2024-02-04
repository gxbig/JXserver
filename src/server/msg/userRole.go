/*用户角色*/
package msg

import (
	"gorm.io/gorm"
	"server/sqlClient"
	"time"
)

type UserRole struct {
	Id              int    `json:"id" gorm:"column:id"`
	UserName        string `json:"userName" gorm:"column:userName"` //用户角色名称
	UserId          int    `json:"userId"`                          //用户id
	GameRoleId      int    `json:"gameRoleId"`                      //角色id
	Grade           int    `json:"grade"`                           //等级
	Attack          int    `json:"attack"`                          //攻击力
	Hp              int    `json:"hp"`                              //生命值
	Mp              int    `json:"mp"`                              //法术值
	Penetrate       int    `json:"penetrate"`                       //穿透
	AvoidInjury     int    `json:"avoidInjury"`                     //免伤
	SpellDamage     int    `json:"spellDamage"`                     //法术伤害
	SpellDefense    int    `json:"spellDefense"`                    //法术防御
	PhysicalDefense int    `json:"physicalDefense"`                 //物理防御
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

func (userRole *UserRole) TableName() string {
	return "user_role"
}

// 获取第一个数据
func (userRole *UserRole) GetUserRole() *UserRole {
	queryUser := &UserRole{}
	sqlClient.DB.Where(userRole).First(queryUser)
	return queryUser
}

// 获取所有用户角色数据
func (userRole *UserRole) GetUserRoles() []UserRole {
	var queryUser []UserRole
	sqlClient.DB.Where(userRole).Find(&queryUser)
	return queryUser
}
