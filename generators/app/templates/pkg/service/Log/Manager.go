package Log

import (
  "<%= appName %>/pkg/service/Log/Logger"

  "log"

  "github.com/spf13/viper"
)

type Manager interface {
  Channel(name string) Logger.Logger
  Driver(name string) Logger.Logger
  Shutdown()
}

type manager struct {
  drivers     map[string]Logger.Logger
  callerDepth int
}

func NewManager() *manager {
  return &manager{
    drivers: make(map[string]Logger.Logger),
    callerDepth: viper.GetInt("log.callerDepth"),
  }
}

func (m *manager) Channel(name string) Logger.Logger {
  return m.Driver(name)
}

func (m *manager) Driver(name string) Logger.Logger {
  if name == "" {
    name = m.getDefaultDriver()
  }
  m.drivers[name] = m.get(name)
  return m.drivers[name]
}

func (m *manager) getDefaultDriver() string {
  return viper.GetString("log.default")
}

func (m *manager) get(name string) Logger.Logger {
  if m.drivers[name] == nil {
    log.Printf("Initializing Logger %s", name)
    return m.resolve(name, m.callerDepth)
  }
  return m.drivers[name]
}

func (m *manager) resolve(name string, calldepth int) Logger.Logger {
  switch name {
  case "null":
    return Logger.NewNull()
  case "stack":
    calldepth = calldepth + 1
    strChannels := viper.GetStringSlice("log.drivers.stack.channels")
    channels := make([]Logger.Logger, len(strChannels))
    for i, chName := range strChannels {
      channels[i] = m.resolve(chName, calldepth + 1)
    }
    return Logger.NewStack(channels)
  case "system":
    return Logger.NewSystem(calldepth)
  case "file":
    return Logger.NewFile(
      viper.GetString("log.drivers.file.runtimeRootPath"),
      viper.GetString("log.drivers.file.savePath"),
      viper.GetString("log.drivers.file.saveName"),
      viper.GetString("log.drivers.file.timeFormat"),
      viper.GetString("log.drivers.file.ext"),
      calldepth,
    )
  }
  return Logger.NewNull()
}

func (m *manager) Shutdown() {
  for name, driver := range m.drivers {
    log.Printf("Closing Logger %s\n", name)
    (driver.(Logger.Closeable)).Close()
  }
}
