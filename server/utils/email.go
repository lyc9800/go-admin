package utils

import (
	"fmt"
	"net/smtp"
	"server/define"

	"github.com/jordan-wright/email"
)

func SendEmail(to string, subject string, text string) {
	e := email.NewEmail()
	// 设置发送方邮箱
	e.From = define.EmailFrom
	// 设置接收方邮箱
	e.To = []string{to}
	// 邮件标题
	e.Subject = subject
	// 邮件内容
	e.Text = []byte(text)
	// 服务器相关设置
	emailAddr := define.EmailHost + ":" + define.EmailPort
	err := e.Send(emailAddr, smtp.PlainAuth("", define.EmailFrom, define.EmailPassword, define.EmailHost))
	if err != nil {
		fmt.Println(err)
	}
}
