package service_impl

import (
  "<%= appName %>/helpers/service"

  "fmt"
  "hash/fnv"
  "path/filepath"

  uuid "github.com/satori/go.uuid"
)

type file struct {
  location string
}

func NewFile() service.File {
  return &file{
    location: "./data/Files",
  }
}

func (s *file) GenBaseName(extension string) string {
  return uuid.NewV4().String() + extension
}

func (s *file) GetPath(fileName string) string {
  hash := s.hash(fileName)
  var mask uint32 = 255
  firstDir := hash & mask
  secondFir := (hash >> 8) & mask
  return filepath.Join(s.location, fmt.Sprintf("%02x", firstDir), fmt.Sprintf("%02x", secondFir), fileName)
}

func (s *file) GetPathDir(fileName string) string {
  return filepath.Dir(s.GetPath(fileName))
}

func (s *file) hash(str string) uint32 {
  h := fnv.New32a()
  _, _ = h.Write([]byte(str))
  return h.Sum32()
}
