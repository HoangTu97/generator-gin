package Mailer

import (
	"<%= appName %>/helpers/service/Mail/Message"

	"log"
)

type mailgunMailer struct {
}

func NewMailgun() *mailgunMailer {
	return &mailgunMailer{}
}

func (m *mailgunMailer) Send(message MailMessage.Message) {
	log.Println("mailgun mailer not implement yet")
}
