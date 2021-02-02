package Mailer

import (
	MailMessage "<%= appName %>/pkg/service/Mail/Message"

	"log"

	"github.com/keighl/postmark"
)

type postmarkMailer struct {
	username string
	client *postmark.Client
}

func NewPostmark(username, serverToken, accountToken string) *postmarkMailer {
	client := postmark.NewClient(serverToken, accountToken)
	return &postmarkMailer{username: username, client: client}
}

func (m *postmarkMailer) Send(message MailMessage.Message) {
	email := postmark.Email{
		From: m.username,
		To: message.GetTo()[0],
		Subject: message.GetSubject(),
		TextBody: message.GetBody(),
		Tag: "pw-reset",
		TrackOpens: true,
	}
	_, err := m.client.SendEmail(email)
	if err != nil {
		log.Println(err)
	}
}
