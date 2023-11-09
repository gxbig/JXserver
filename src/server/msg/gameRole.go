/**游戏角色*/
package msg

import (
	"gorm.io/gorm"
	"time"
)

type GameRole struct {
	Id       int `json:"id" gorm:"column:id"`
	RoleName int `json:"roleName"` //角色名称
	RoleType int `json:"roleType"` //角色id

	Attack            int `json:"attack"`            //攻击力
	AttackGrowUp      int `json:"attackGrowUp"`      //攻击力每级增长
	Hp                int `json:"hp"`                //生命值
	Mp                int `json:"mp"`                //法术值
	Penetrate         int `json:"penetrate"`         //穿透
	AvoidInjury       int `json:"avoidInjury"`       //免伤
	HpGrowUp          int `json:"hpGrowUp"`          //hp每级增长
	MpGrowUp          int `json:"mpGrowUp"`          //mp每级增长
	SpellDamage       int `json:"spellDamage"`       //法术伤害
	SpellDamageGrowUp int `json:"spellDamageGrowUp"` //法术伤害每级增长

	SpellDefense       int `json:"spellDefense"`       //法术防御
	SpellDefenseGrowUp int `json:"spellDefenseGrowUp"` //法术防御每级增长

	PhysicalDefense int `json:"physicalDefense"` //物理防御
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

func (gameRole *GameRole) TableName() string {
	return "game_role"
}
