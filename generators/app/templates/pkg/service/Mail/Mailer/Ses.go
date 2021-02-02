package Mailer

import (
	MailMessage "<%= appName %>/pkg/service/Mail/Message"

	"log"
)

type sesMailer struct {
}

func NewSes() *sesMailer {
	return &sesMailer{}
}

func (m *sesMailer) Send(message MailMessage.Message) {
	log.Println("Ses mailer not implement yet")
}
