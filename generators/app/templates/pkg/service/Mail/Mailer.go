package Mail

import (
  "<%= appName %>/pkg/service/Mail/Message"
)

type Mailer interface {
	Send(message MailMessage.Message)
}