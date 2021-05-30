package CacheStore

import (
  "time"

  "github.com/patrickmn/go-cache"
)

type memcachedStore struct {
  conn *cache.Cache
}

func NewMemcached(defaultExpiration, purgeDuration time.Duration) *memcachedStore {
  conn := cache.New(defaultExpiration, purgeDuration)
  return &memcachedStore{conn: conn}
}

func (s *memcachedStore) Get(key string) interface{} {
  c := s.connection()
  value, exists := c.Get(key)
  if !exists {
    return nil
  }
  return value
}

func (s *memcachedStore) Many(keys []string) []interface{} {
  return make([]interface{}, 0)
}

func (s *memcachedStore) Add(key string, value interface{}, ttl int) bool {
  c := s.connection()
  _, exists := c.Get(key)
  if exists {
    return false
  }
  c.Set(key, value, time.Duration(ttl))
  return true
}

func (s *memcachedStore) Put(key string, value interface{}, ttl int) bool {
  c := s.connection()
  c.Set(key, value, time.Duration(ttl))
  return true
}

func (s *memcachedStore) Increment(key string, value int) bool {
  return true
}

func (s *memcachedStore) Decrement(key string, value int) bool {
  return true
}

func (s *memcachedStore) Forever(key string, value interface{}) bool {
  return true
}

func (s *memcachedStore) Forget(key string) bool {
  c := s.connection()
  c.Delete(key)
  return true
}

func (s *memcachedStore) Flush() bool {
  return true
}

func (s *memcachedStore) connection() *cache.Cache {
  return s.conn
}
