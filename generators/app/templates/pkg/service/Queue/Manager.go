package Queue

import (
  "<%= appName %>/pkg/service/Queue/Driver"

  "fmt"
  "log"

  "github.com/spf13/viper"
)

type Manager interface {
  Driver(driver string) Service
  Shutdown()
}

type manager struct {
  drivers map[string]Service
}

func NewManager() Manager {
  return &manager{
    drivers: make(map[string]Service),
  }
}

func (m *manager) Driver(name string) Service {
  if name == "" {
    name = m.getDefaultDriver()
  }
  m.drivers[name] = m.get(name)
  return m.drivers[name]
}

func (m *manager) get(name string) Service {
  if m.drivers[name] == nil {
    log.Printf("Connecting Queue %s", name)
    return m.resolve(name)
  }
  return m.drivers[name]
}

func (m *manager) resolve(name string) Service {
  switch name {
  case "RabbitMQ":
    return QueueDriver.NewRabbitMQ(
      fmt.Sprintf("amqp://%s:%s@%s:%d/",
        viper.GetString("queue.drivers.RabbitMQ.username"),
        viper.GetString("cache.drivers.RabbitMQ.password"),
        viper.GetString("cache.drivers.RabbitMQ.host"),
        viper.GetInt("cache.drivers.RabbitMQ.port"),
      ),
    )
  }
  return nil
}

func (m *manager) getDefaultDriver() string {
  return viper.GetString("queue.default")
}

func (m *manager) Shutdown() {
  for key, driver := range m.drivers {
    log.Printf("Disconnecting Queue connection : %s \n", key)
    driver.Close()
  }
}
