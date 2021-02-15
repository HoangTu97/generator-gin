package Mail

import (
  "<%= appName %>/pkg/service/Log/Logger"
  MailMailer "<%= appName %>/pkg/service/Mail/Mailer"


  "github.com/spf13/viper"
)

type Manager interface {
  Mailer(name string) Service
  Driver(name string) Service
}

type manager struct {
  mailers map[string]Service
  logger Logger.Logger
}

func NewManager(logger Logger.Logger) Manager {
  return &manager{
    mailers: make(map[string]Service),
    logger: logger,
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
    m.logger.Info("Connecting Mailer ", name)
    return NewService(m.resolve(name))
  }
  return m.mailers[name]
}

func (m *manager) resolve(name string) Mailer {
  switch name {
  case "smtp":
    return MailMailer.NewSmtp(
      viper.GetString("mail.mailers.username"),
      viper.GetString("mail.mailers.password"),
      viper.GetString("mail.mailers.host"),
      viper.GetString("mail.mailers.port"),
    )
  case "ses":
    return MailMailer.NewSes("abc", "us-west-2")
  case "mailgun":
    return MailMailer.NewMailgun("abc", "localhost", "")
  case "postmark":
    return MailMailer.NewPostmark("abc", "[SERVER-TOKEN]", "[ACCOUNT-TOKEN]")
  case "sendgrid":
    return MailMailer.NewSendgrid("abc", "SENDGRID_API_KEY")
  }
  return nil
}

func (m *manager) getDefaultDriver() string {
  return viper.GetString("mail.default")
}
