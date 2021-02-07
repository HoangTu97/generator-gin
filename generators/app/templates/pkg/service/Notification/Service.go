package Notification

import (
  "context"
  "log"

  "github.com/spf13/viper"
  "google.golang.org/api/fcm/v1"
  "google.golang.org/api/option"
)

type Service interface {
  NewMessage() fcm.Message
  Send(message fcm.Message)
}

type service struct {
  parentId string
  fcmService *fcm.Service
}

func NewService() Service {
  ctx := context.Background()
  fcmService, err := fcm.NewService(ctx, option.WithAPIKey(viper.GetString("fcm.apiKey")))
  if err != nil {
    log.Fatal("New fcm service error ", err)
  }
  return &service{fcmService: fcmService, parentId: viper.GetString("fcm.parentId")}
}

func (s *service) NewMessage() fcm.Message {
  return fcm.Message{}
}

func (s *service) Send(message fcm.Message) {
  projectsMessagesSendCall := s.fcmService.Projects.Messages.Send(
    s.parentId, &fcm.SendMessageRequest{Message: &message},
  )
  response, err := projectsMessagesSendCall.Do()
  if err != nil {
    log.Println(err)
  }
  log.Println(response)
}
