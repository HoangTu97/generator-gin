package config

import (
  "<%= appName %>/helpers/setting"
  "<%= appName %>/pkg/cache"
  "<%= appName %>/pkg/database"
  "log"
  "time"

  "github.com/go-ini/ini"
)

var AppSetting = &setting.App{}
var LoggerSetting = &setting.Logger{}
var ServerSetting = &setting.Server{}
var DatabaseSetting = &setting.Database{Config: &database.Config{}}
var CacheSetting = &setting.Cache{Config: &cache.Config{}}
var RabbitMQSetting = &setting.RabbitMQ{}

// Setup initialize the configuration instance
func Setup() {
  var err error
  cfg, err := ini.Load("app.ini")
  if err != nil {
    log.Fatalf("setting.Setup, fail to parse 'app.ini': %v", err)
  }

  mapTo(cfg, "app", AppSetting)
  mapTo(cfg, "logger", LoggerSetting)
  mapTo(cfg, "server", ServerSetting)
  mapTo(cfg, "database", DatabaseSetting.Config)
  mapTo(cfg, "cache", CacheSetting.Config)

  ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
  ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
  CacheSetting.IdleTimeout = CacheSetting.IdleTimeout * time.Second
}

// mapTo map section
func mapTo(cfg *ini.File, section string, v interface{}) {
  err := cfg.Section(section).MapTo(v)
  if err != nil {
    log.Fatalf("Cfg.MapTo %s err: %v", section, err)
  }
}
