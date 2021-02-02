package Mailer

import (
	"<%= appName %>/helpers/service/Mail/Message"

	"log"
)

type arrayMailer struct {
}

func NewArray() *arrayMailer {
	return &arrayMailer{}
}

func (m *arrayMailer) Send(message MailMessage.Message) {
	log.Println("array mailer not implement yet")
}
