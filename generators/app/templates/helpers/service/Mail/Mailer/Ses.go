package Mailer

import (
	"<%= appName %>/helpers/service/Mail/Message"

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
