package smtp

import (
	"github.com/junminhong/member-center-service/pkg/logger"
	"github.com/spf13/viper"
	_smtp "net/smtp"
)

var zapLogger = logger.Setup()

type EmailMsg struct {
	message []byte
}

func (emailMsg *EmailMsg) SetUpEmailMsg(title string, body string) {
	emailMsg.message = []byte(
		"Subject: " + title + "\r\n" +
			"From: membercentersmtp@gmail.com\r\n" +
			`Content-Type: text/plain; boundary="qwertyuio"` + "\r\n" +
			"\r\n" +
			body + "\r\n" +
			"\r\n",
	)
}

func (emailMsg *EmailMsg) SendEmail(email string) {
	account := viper.GetString("SMTP.EMAIL_ACCOUNT")
	password := viper.GetString("SMTP.EMAIL_PASSWORD")
	targetEmail := []string{
		email,
	}
	message := emailMsg.message
	auth := _smtp.PlainAuth("", account, password, "smtp.gmail.com")
	err := _smtp.SendMail(viper.GetString("SMTP.EMAIL_HOST")+":"+viper.GetString("SMTP.EMAIL_PORT"), auth, account, targetEmail, message)
	if err != nil {
		zapLogger.Info(err.Error())
	}
}
