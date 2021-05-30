package Hash

import (
  "crypto/md5"
  "encoding/hex"
)

type md5Service struct {
}

func NewMd5() Service {
  return &md5Service{}
}

func (s *md5Service) Make(value string) string {
  m := md5.New()
  m.Write([]byte(value))
  return hex.EncodeToString(m.Sum(nil))
}

func (s *md5Service) Check(hashedValue, value string) bool {
  hashedValue2 := s.Make(value)
  return hashedValue == hashedValue2
}
