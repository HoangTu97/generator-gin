package File

import (
  "fmt"
  "hash/fnv"
  "path/filepath"

  uuid "github.com/satori/go.uuid"
)

type Service interface {
  GenBaseName(extension string) string
  GetPath(fileName string) string
  GetPathDir(fileName string) string
}

type service struct {
  location string
}

func NewService() Service {
  return &service{
    location: "./data/Files",
  }
}

func (s *service) GenBaseName(extension string) string {
  return uuid.NewV4().String() + extension
}

func (s *service) GetPath(fileName string) string {
  hash := s.hash(fileName)
  var mask uint32 = 255
  firstDir := hash & mask
  secondFir := (hash >> 8) & mask
  return filepath.Join(s.location, fmt.Sprintf("%02x", firstDir), fmt.Sprintf("%02x", secondFir), fileName)
}

func (s *service) GetPathDir(fileName string) string {
  return filepath.Dir(s.GetPath(fileName))
}

func (s *service) hash(str string) uint32 {
  h := fnv.New32a()
  _, _ = h.Write([]byte(str))
  return h.Sum32()
}
