package Mailer

import (
	"<%= appName %>/helpers/service/Mail/Message"

	"log"
)

type postmarkMailer struct {
}

func NewPostmark() *postmarkMailer {
	return &postmarkMailer{}
}

func (m *postmarkMailer) Send(message MailMessage.Message) {
	log.Println("postmark mailer not implement yet")
}
