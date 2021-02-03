package Mailer

import (
  MailMessage "<%= appName %>/pkg/service/Mail/Message"

  "net/smtp"
  "fmt"
)

type smtpMailer struct {
  username string
  host string
  portNumber string
  auth smtp.Auth
}

func NewSmtp(username, password, host, portNumber string) *smtpMailer {
  auth := smtp.PlainAuth("", username, password, host)
  return &smtpMailer{auth: auth, username: username, host: host, portNumber: portNumber}
}

func (m *smtpMailer) Send(message MailMessage.Message) {
  smtp.SendMail(fmt.Sprintf("%s:%s", m.host, m.portNumber), m.auth, m.username, message.GetTo(), message.ToBytes())
}
