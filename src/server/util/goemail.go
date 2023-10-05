package util

import (
	"fmt"
	"github.com/name5566/leaf/log"
	"gopkg.in/gomail.v2"
)

var UGoemail = uGoemail{}

type uGoemail struct{}

/**
to 主送 机构联系邮箱
cc 抄送  创建人邮箱
subject 标题
content 内容
result 1成功 2失败

host 邮件服务器
port 端口号
username 邮箱
password 授权码
*/

func SendMail(code string, to string) {

	// Atoi相当于ParseInt（s，10，0），转换为int类型
	m := gomail.NewMessage()
	m.SetHeader("From", "1101434570@qq.com")
	m.SetHeader("To", to)

	m.SetHeader("Subject", "金仙游戏")
	m.SetBody("text/html", fmt.Sprintf("注册码： <b>%s</b>（请不要向其他人透露，注册码有效期5分钟！）", code))

	d := gomail.NewDialer("smtp.qq.com", 587, "1101434570", "briqwxqjkppzffae")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		log.Debug(err.Error())
	}

}
