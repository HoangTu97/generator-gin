package Mail

import (
  SmsMessage "<%= appName %>/pkg/service/Sms/Message"
)

type Service interface {
  NewMessage() SmsMessage.Message
  Send(message SmsMessage.Message)
}

type service struct {
  sender Sender
}

func NewService(sender Sender) Service {
  return &service{sender: sender}
}

func (s *service) NewMessage() SmsMessage.Message {
  return SmsMessage.NewMessage()
}

func (s *service) Send(message SmsMessage.Message) {
  s.sender.Send(message)
}
