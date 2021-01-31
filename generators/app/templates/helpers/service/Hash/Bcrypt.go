package Hash

import (
  "golang.org/x/crypto/bcrypt"
)

type bcryptService struct {
}

func NewBcrypt() Service {
  return &bcryptService{}
}

func (s *bcryptService) Make(value string) string {
  hashedValue, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
  if err != nil {
    return ""
  }
  return string(hashedValue)
}

func (s *bcryptService) Check(value string, hashedValue string) bool {
  err := bcrypt.CompareHashAndPassword([]byte(value), []byte(hashedValue))
  if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
    return false
  }
  return true
}
