package models

import (
	"fmt"
	"net/smtp"
	"strings"
)

type Mail struct {
	From    string
	To      []string
	Subject string
	Body    string
}

func (m *Mail) Build() []byte {
	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: %s\r\n", m.From)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(m.To, ","))
	msg += fmt.Sprintf("Subject: %s\r\n", m.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", m.Body)

	return []byte(msg)
}

func (m *Mail) Send(from string, password string) error {
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, m.From, m.To, m.Build())
	if err != nil {
		return err
	}
	return nil
}
