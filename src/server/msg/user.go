package msg

import (
	"gorm.io/gorm"
	"server/sqlClient"
)

// 用户表
type User struct {
	gorm.Model
	Id          int    `json:"id" gorm:"column:id"`
	Age         int    `json:"loginName" gorm:"column:age"`                    // 年龄
	Gender      string `json:"gender" gorm:"column:gender;type:enum('0','1')"` // 性别
	Code        string `json:"code" gorm:"-"`                                  //验证码
	Name        string `json:"name" gorm:"column:name"`                        // 姓名
	Identity    string `json:"identity" gorm:"column:identity"`                // 身份证号码
	DateOfBirth string `json:"dateOfBirth" `                                   // 出生日期
	Address     string `json:"address" `                                       // 地址
	Server      string `json:"server" `                                        // 服务区
	PostalCode  string `json:"postalCode" `                                    // 邮政编码
	Qq          string `json:"qq" `                                            // qq
	Wx          string `json:"wx" `                                            // 微信
	Zfb         string `json:"zfb"`                                            // 支付宝
	Account     string `json:"account"`                                        // 账户
	Pw          string `json:"pw" gorm:"column:pw"`                            // 密码
	Phone       string `json:"phone" gorm:"column:phone"`                      // 手机号
	Email       string `json:"email"  gorm:"column:email"`                     // 邮箱

}

func (user *User) TableName() string {
	return "user"
}

// 插入数据
func (user *User) RegisterInsetUser() error {

	result := sqlClient.DB.Create(user)
	return result.Error
}

// 数据
func (user *User) QueryUser() (*User, error) {
	queryUser := &User{}
	result := sqlClient.DB.Where(user).First(queryUser)
	return queryUser, result.Error
}

// 更新密码数据
func (user *User) UpdateUserPw() error {
	result := sqlClient.DB.Model(user).Updates(map[string]interface{}{"pw": user.Pw})
	return result.Error
}
func (user *User) DeleteUser() error {
	result := sqlClient.DB.Where(user).Delete(user)
	return result.Error
}
