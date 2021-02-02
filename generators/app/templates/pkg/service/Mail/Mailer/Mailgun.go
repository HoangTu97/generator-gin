package Mailer

import (
  "<%= appName %>/pkg/service/Mail/Message"

  "context"
  "log"
  "time"

  "github.com/mailgun/mailgun-go/v4"
)

type mailgunMailer struct {
  username string
  domain string
  apiKey string
  mg *mailgun.MailgunImpl
}

func NewMailgun(username, domain, apiKey string) *mailgunMailer {
  mg := mailgun.NewMailgun(domain, apiKey)
  return &mailgunMailer{username: username, mg:mg}
}

func (m *mailgunMailer) Send(message MailMessage.Message) {
  mgMsg := m.mg.NewMessage(m.username, message.GetSubject(), message.GetBody(), message.GetTo()[0])
  ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
  defer cancel()

  if len(message.GetTo()) > 1 {
    for i := 1; i < len(message.GetTo()); i++ {
      mgMsg.AddRecipient(message.GetTo()[i])
    }
  }

  // Send the message with a 10 second timeout
  resp, id, err := m.mg.Send(ctx, mgMsg)
  if err != nil {
    log.Println(err)
  }
  log.Printf("ID: %s Resp: %s\n", id, resp)
}
