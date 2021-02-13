package Sms

import (
  SmsSender "<%= appName %>/pkg/service/Sms/Sender"

  "log"

  "github.com/spf13/viper"
)

type Manager interface {
  Driver(name string) Service
}

type manager struct {
  senders map[string]Service
}

func NewManager() Manager {
  return &manager{
    senders: make(map[string]Service),
  }
}

func (m *manager) Driver(name string) Service {
  if name == "" {
    name = m.getDefaultDriver()
  }
  m.senders[name] = m.get(name)
  return m.senders[name]
}

func (m *manager) get(name string) Service {
  if m.senders[name] == nil {
    log.Printf("Connecting Mailer %s", name)
    return NewService(m.resolve(name))
  }
  return m.senders[name]
}

func (m *manager) resolve(name string) Sender {
  switch name {
  case "vonage":
    return SmsSender.NewVonage(
      viper.GetString("sms.drivers.vonage.apiKey"),
      viper.GetString("sms.drivers.vonage.apiSecret"),
    )
  }
  return nil
}

func (m *manager) getDefaultDriver() string {
  return viper.GetString("sms.default")
}
