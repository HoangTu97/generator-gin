package Mail

import (
  SmsMessage "<%= appName %>/pkg/service/Sms/Message"
)

type Sender interface {
  Send(message SmsMessage.Message)
}