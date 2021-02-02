package Crypt

import (
  "log"
)

type Service interface {
  Encrypt(value interface{}) string
  Decrypt(payload string) interface{}
}

type service struct {
  key string
  cipher string
}

func NewService(key string, cipher string) Service {
  return &service{key:key, cipher: cipher}
}

func (s *service) Encrypt(value interface{}) string {
  log.Println("Crypt service : Encrypt not implement")
  return "abc"
}

func (s *service) Decrypt(payload string) interface{} {
  log.Println("Crypt service : Decrypt not implement")
  return nil
}
