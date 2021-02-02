package Mail

import (
  MailMailer "<%= appName %>/pkg/service/Mail/Mailer"
)

type Manager interface {
  Mailer(name string) Service
  Driver(name string) Service
}

type manager struct {
  mailers map[string]Service
}

func NewManager() Manager {
  return &manager{
    mailers: make(map[string]Service),
  }
}

func (m *manager) Mailer(name string) Service {
  if name == "" {
    name = m.getDefaultDriver()
  }
  m.mailers[name] = m.get(name)
  return m.mailers[name]
}

func (m *manager) Driver(name string) Service {
  return m.Mailer(name)
}

func (m *manager) get(name string) Service {
  if m.mailers[name] == nil {
    return NewService(m.resolve(name))
  }
  return m.mailers[name]
}

func (m *manager) resolve(name string) Mailer {
  switch name {
  case "smtp":
    return MailMailer.NewSmtp("abc", "", "localhost", "8080")
  case "ses":
    return MailMailer.NewSes()
  case "mailgun":
    return MailMailer.NewMailgun()
  case "postmark":
    return MailMailer.NewPostmark()
  }
  return nil
}

func (m *manager) getDefaultDriver() string {
  return "mailgun"
}
