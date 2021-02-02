package Mail

import (
  "<%= appName %>/pkg/service/Mail/Message"
)

type Service interface {
  NewMessage() MailMessage.Message

  // Send
  Send(message MailMessage.Message)
  // Queue(message MailMessage.Message)
  // Later(message MailMessage.Message, firetime int)
}

type service struct {
  mailer Mailer
}

func NewService(mailer Mailer) Service {
  return &service{mailer: mailer}
}

func (s *service) NewMessage() MailMessage.Message {
  return MailMessage.NewMessage()
}

func (s *service) Send(message MailMessage.Message) {
  s.mailer.Send(message)
}
