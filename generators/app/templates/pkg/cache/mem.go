package cache

import (
  "encoding/json"
  "time"

  "github.com/patrickmn/go-cache"
)

type memCache struct {
  conn *cache.Cache
}

func NewMem(config Config) Cache {
  conn := cache.New(5*time.Minute, 10*time.Minute)

  return &memCache{conn: conn}
}

func (r *memCache) GenKey(data ...interface{}) string {
  return genKey(data)
}

// Set a key/value
func (c *memCache) Set(key string, data interface{}, time int) error {
  c.conn.Set(key, data, cache.DefaultExpiration)
  return nil
}

// Exists check a key
func (c *memCache) Exists(key string) bool {
  _, exists := c.conn.Get(key)
  return exists
}

// Get get a key
func (c *memCache) Get(key string) ([]byte, error) {
  reply, exists := c.conn.Get(key)
  if !exists {
    return nil, nil
  }

  value, err := json.Marshal(reply)
  if err != nil {
    return nil, err
  }

  return value, nil
}

// Delete delete a kye
func (c *memCache) Delete(key string) (bool, error) {
  c.conn.Delete(key)
  return true, nil
}

// LikeDeletes batch delete
func (c *memCache) LikeDeletes(key string) error {
  return nil
}
