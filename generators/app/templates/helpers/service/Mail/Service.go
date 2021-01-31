package Mail

import (
  "log"
)

type Service interface {
  NewMessage() Message

  // Send
  Send(message Message)
  // Queue(message Message)
  // Later(message Message, firetime int)
}

type service struct {
}

func NewService() Service {
  return &service{}
}

func (s *service) NewMessage() Message {
  return NewMessage()
}

func (s *service) Send(message Message) {
  log.Println("Mail service : Send not implement : " + message.String())
}
