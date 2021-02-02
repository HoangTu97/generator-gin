package Mail

import (
  "<%= appName %>/helpers/service/Mail/Message"
)

type Mailer interface {
	Send(message MailMessage.Message)
}