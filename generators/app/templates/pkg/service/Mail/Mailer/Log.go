package Mailer

import (
	"<%= appName %>/pkg/service/Mail/Message"

	"log"
)

type logMailer struct {
}

func NewLog() *logMailer {
	return &logMailer{}
}

func (m *logMailer) Send(message MailMessage.Message) {
	log.Println("log mailer not implement yet")
}
