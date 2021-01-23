package cache

import (
  "<%= appName %>/pkg/converter"
  "strings"
)

type Cache interface {
  GenKey(data ...interface{}) string
  Set(key string, data interface{}, time int) error
  Exists(key string) bool
  Get(key string) ([]byte, error)
  Delete(key string) (bool, error)
  LikeDeletes(key string) error
}

func NewCache(config Config) Cache {
  if config.Type == "redis" {
    return NewRedis(config)
  }
  if config.Type == "memory" {
    return NewMem(config)
  }
  return nil
}

func genKey(data ...interface{}) string {
  values := make([]string, len(data))

  for i, dt := range data {
    values[i] = converter.ToStr(dt)
  }

  return strings.Join(values, "_")
}
