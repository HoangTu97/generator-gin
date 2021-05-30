package Cache

import (
  "<%= appName %>/pkg/converter"
  "strings"
)

type Service interface {
  GenKey(data ...interface{}) string

  Has(key string) bool
  Missing(key string) bool
  Get(key string) interface{}
  Set(key string, value interface{}, ttl int) bool
  Delete(key string) bool
  Clear() bool
  Many(keys []string) map[string]interface{}
  PutMany(values map[string]interface{}, ttl int) bool
  DeleteMultiple(keys []string) bool

  Pull(key string) interface{}
  Put(key string, value interface{}, ttl int) bool
  Add(key string, value interface{}, ttl int) bool
  Increment(key string, value int) bool
  Increment1(key string) bool
  Decrement(key string, value int) bool
  Decrement1(key string) bool
  Forever(key string, value interface{}) bool
  Forget(key string) bool
}

type service struct {
  repository Repository
}

func NewService(repository Repository) Service {
  return &service{repository: repository}
}

func (s *service) GenKey(data ...interface{}) string {
  values := make([]string, len(data))

  for i, dt := range data {
    values[i] = converter.ToStr(dt)
  }

  return strings.Join(values, "_")
}

func (s *service) Has(key string) bool {
  return s.repository.Has(key)
}

func (s *service) Missing(key string) bool {
  return s.repository.Missing(key)
}

func (s *service) Get(key string) interface{} {
  return s.repository.Get(key)
}

func (s *service) Set(key string, value interface{}, ttl int) bool {
  return s.repository.Set(key, value, ttl)
}

func (s *service) Delete(key string) bool {
  return s.repository.Delete(key)
}

func (s *service) Clear() bool {
  return s.repository.Clear()
}

func (s *service) Many(keys []string) map[string]interface{} {
  return s.repository.Many(keys)
}

func (s *service) PutMany(values map[string]interface{}, ttl int) bool {
  return s.repository.PutMany(values, ttl)
}

func (s *service) DeleteMultiple(keys []string) bool {
  return s.repository.DeleteMultiple(keys)
}

func (s *service) Pull(key string) interface{} {
  return s.repository.Pull(key)
}

func (s *service) Put(key string, value interface{}, ttl int) bool {
  return s.repository.Put(key, value, ttl)
}

func (s *service) Add(key string, value interface{}, ttl int) bool {
  return s.repository.Add(key, value, ttl)
}

func (s *service) Increment(key string, value int) bool {
  return s.repository.Increment(key, value)
}

func (s *service) Increment1(key string) bool {
  return s.repository.Increment1(key)
}

func (s *service) Decrement(key string, value int) bool {
  return s.repository.Decrement(key, value)
}

func (s *service) Decrement1(key string) bool {
  return s.repository.Decrement1(key)
}

func (s *service) Forever(key string, value interface{}) bool {
  return s.repository.Forever(key, value)
}

func (s *service) Forget(key string) bool {
  return s.repository.Forget(key)
}
