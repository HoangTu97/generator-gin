package Cache

import (
  "<%= appName %>/pkg/service/Cache/Store"

  "fmt"

  "github.com/spf13/viper"
)

type Manager interface {
  Store(name string) Service
  Driver(driver string) Service
  ForgetDriver(name string) Manager
}

type manager struct {
  stores map[string]Service
}

func NewManager() Manager {
  return &manager{
    stores: make(map[string]Service),
  }
}

func (m *manager) Store(name string) Service {
  if name == "" {
    name = m.getDefaultDriver()
  }
  m.stores[name] = m.get(name)
  return m.stores[name]
}

func (m *manager) Driver(name string) Service {
  return m.Store(name)
}

func (m *manager) get(name string) Service {
  if m.stores[name] == nil {
    return NewService(m.resolve(name))
  }
  return m.stores[name]
}

func (m *manager) resolve(name string) Repository {
  switch name {
  // case "Apc":
  //   return m.repository(CacheStore.NewApc())
  // case "Array":
  //   return m.repository(CacheStore.NewArray())
  // case "File":
  //   return m.repository(CacheStore.NewFile())
  case "memcached":
    return m.repository(CacheStore.NewMemcached(
      viper.GetDuration("cache.drivers.memcached.defaultExpiration"),
      viper.GetDuration("cache.drivers.memcached.purgeDuration"),
    ))
  case "null":
    return m.repository(CacheStore.NewNull())
  case "redis":
    return m.repository(CacheStore.NewRedis(
      fmt.Sprintf("%s:%s",viper.GetString("cache.drivers.redis.host"),viper.GetString("cache.drivers.redis.port")),
      viper.GetString("cache.drivers.redis.port"),
      viper.GetInt("cache.drivers.redis.maxIdle"),
      viper.GetInt("cache.drivers.redis.maxActive"),
      viper.GetDuration("cache.drivers.redis.idleTimeout"),
    ))
  // case "Database":
  //   return m.repository(CacheStore.NewDatabase())
  // case "Dynamodb":
  //   return m.repository(CacheStore.NewDynamodb())
  }
  return m.repository(CacheStore.NewNull())
}

func (m *manager) getDefaultDriver() string {
  return viper.GetString("cache.default")
}

func (m *manager) repository(store Store) Repository {
  return NewRepository(store)
}

func (m *manager) ForgetDriver(name string) Manager {
  for key := range m.stores {
    if (key == name) {
      delete(m.stores, key)
    }
  }
  return m
}
