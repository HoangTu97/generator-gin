package Mailer

import (
	"<%= appName %>/helpers/service/Mail/Message"

	"log"
)

type sendmailMailer struct {
}

func NewSendmail() *sendmailMailer {
	return &sendmailMailer{}
}

func (m *sendmailMailer) Send(message MailMessage.Message) {
	log.Println("sendmail mailer not implement yet")
}
