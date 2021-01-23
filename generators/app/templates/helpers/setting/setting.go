package setting

import (
  "<%= appName %>/pkg/cache"
  "<%= appName %>/pkg/database"
  "<%= appName %>/pkg/logging"
  "time"
)

type App struct {
  JwtSecret       string
  PageSize        int
}

type Logger struct {
  logging.Config
}

type Server struct {
  RunMode      string
  HTTPPort     string
  SSL          bool
  ReadTimeout  time.Duration
  WriteTimeout time.Duration
}

type Database struct {
  *database.Config
}

type Cache struct {
  *cache.Config
}

type RabbitMQ struct {
  Host     string
  Port     string
  User     string
  Password string
}