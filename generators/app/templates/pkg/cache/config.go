package cache

import "time"

type Config struct {
  Type        string
  Host        string
  Port        string
  Password    string
  SSL         bool
  MaxIdle     int
  MaxActive   int
  IdleTimeout time.Duration
}
