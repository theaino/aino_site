package misc

import (
	"aino-spring.com/aino_site/config"
	"gopkg.in/gomail.v2"
)

type Emailer struct {
	Dialer *gomail.Dialer
	Email  string
}

func NewEmailer(conf *config.Config) *Emailer {
	dialer := gomail.NewDialer(conf.SMTPHost, conf.SMTPPort, conf.Email, conf.EmailPassword)
	emailer := Emailer{Dialer: dialer, Email: conf.Email}
	return &emailer
}

func (emailer *Emailer) SendMail(email string, subject string, content string) error {
	s, err := emailer.Dialer.Dial()
	if err != nil {
		return err
	}

	message := gomail.NewMessage()

	message.SetHeader("From", emailer.Email)
	message.SetAddressHeader("To", email, email)
	message.SetHeader("Subject", subject)
	message.SetBody("text/html", content)

	return gomail.Send(s, message)
}
