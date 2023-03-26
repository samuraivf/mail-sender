package mail

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"os"
)

type Sender struct {
	fromEmail    string
	fromAppPassword string
	smtpHost     string
	smtpPort     string
}

func NewSender(config SenderConfig) *Sender {
	return &Sender{
		fromEmail:    config.FromEmail,
		fromAppPassword: config.FromAppPassword,
		smtpHost:     config.SMTPHost,
		smtpPort:     config.SMTPPort,
	}
}

func (s *Sender) Send(msg string) error {
	parsedMsg, err := parseMessage(msg)
	if err != nil {
		return err
	}

	auth := smtp.PlainAuth("", s.fromEmail, s.fromAppPassword, s.smtpHost)

	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	t, err := template.ParseFiles(wd + "/templates/template.html")
	if err != nil {
		return err
	}

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: Email Verification \n%s\n\n", mimeHeaders)))

	t.Execute(&body, struct{
		Link string
	}{Link: os.Getenv("BUG_TRACKER_ADDRESS") + parsedMsg.link})

	err = smtp.SendMail(s.smtpHost + ":" + s.smtpPort, auth, s.fromEmail, []string{parsedMsg.email}, body.Bytes())
	if err != nil {
		return err
	}
	
	return nil
}
