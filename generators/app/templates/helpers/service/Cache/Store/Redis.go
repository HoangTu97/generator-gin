package CacheStore

import (
  "encoding/json"
  "time"

  "github.com/gomodule/redigo/redis"
)

type redisStore struct {
  conn *redis.Pool
}

func NewRedis() *redisStore {
  redisConn := &redis.Pool{
    MaxIdle:     30,
    MaxActive:   30,
    IdleTimeout: -1,
    Dial: func() (redis.Conn, error) {
      c, err := redis.Dial("tcp", "localhost")
      if err != nil {
        return nil, err
      }
      if "" != "" {
        if _, err := c.Do("AUTH", ""); err != nil {
          c.Close()
          return nil, err
        }
      }
      return c, err
    },
    TestOnBorrow: func(c redis.Conn, t time.Time) error {
      _, err := c.Do("PING")
      return err
    },
  }
  return &redisStore{conn: redisConn}
}

func (s *redisStore) Get(key string) interface{} {
  c := s.connection()
  defer c.Close()

  bytes, err := redis.Bytes(c.Do("GET", key))
  if err != nil {
    return nil
  }

  var value interface{}
  err = json.Unmarshal(bytes, &value)
  if err != nil {
    return nil
  }
  return value
}

func (s *redisStore) Many(keys []string) []interface{} {
  return make([]interface{}, 0)
}

func (s *redisStore) Put(key string, value interface{}, ttl int) bool {
  c := s.connection()
  defer c.Close()
  data, err := json.Marshal(value)
  if err != nil {
    return false
  }
  _, err = c.Do("SETEX", key, ttl, data)
  if err != nil {
    return false
  }
  return true
}

func (s *redisStore) Add(key string, value interface{}, ttl int) bool {
  c := s.connection()
  defer c.Close()
  var err error
  exists, err := redis.Bool(c.Do("EXISTS", key))
  if err != nil || exists {
    return false
  }
  data, err := json.Marshal(value)
  if err != nil {
    return false
  }
  _, err = c.Do("SETEX", key, ttl, data)
  if err != nil {
    return false
  }
  return true
}

func (s *redisStore) Increment(key string, value int) bool {
  return true
}

func (s *redisStore) Decrement(key string, value int) bool {
  return true
}

func (s *redisStore) Forever(key string, value interface{}) bool {
  return true
}

func (s *redisStore) Forget(key string) bool {
  c := s.connection()
  defer c.Close()
  success, err := redis.Bool(c.Do("DEL", key))
  if !success || err != nil {
    return false
  }
  return true
}

func (s *redisStore) Flush() bool {
  c := s.connection()
  defer c.Close()
  err := c.Flush()
  if err != nil {
    return false
  }
  return true
}

func (s *redisStore) connection() redis.Conn {
  return s.conn.Get()
}
