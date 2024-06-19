package util

import (
	"fmt"
	"net/smtp"

	"aioc/pkg/setting"
)

var EmailMaster, EmailMasterPwd string

func init() {

}

func SendEmail(target, title, body string) error {

	msg := []byte("From: fan " + setting.EmailSetting.Email + "\r\n" +
		"To: " + target + "\r\n" +
		"Subject: " + title + " \r\n" +
		"\r\n" +
		body)

	auth := smtp.PlainAuth("", setting.EmailSetting.Email, setting.EmailSetting.Password, setting.EmailSetting.Smtp)                                          // 邮箱认证信息
	return smtp.SendMail(fmt.Sprintf("%s:%d", setting.EmailSetting.Smtp, setting.EmailSetting.Port), auth, setting.EmailSetting.Email, []string{target}, msg) // 发送邮件
}
