package SmsMessage

type Message interface {
  From(phone string) Message
  To(phone string) Message
  Text(text string) Message

  GetFrom() string
  GetTo() string
  GetText() string

  String() string
}

type message struct {
  from string
  to []string
  text string
}

func NewMessage() Message {
  return &message{}
}

func (m *message) From(phone string) Message {
  m.from = phone
  return m
}

func (m *message) To(phone string) Message {
  m.to = phone
  return m
}

func (m *message) Text(text string) Message {
  m.text = text
  return m
}

func (m *message) GetFrom() string {
  return m.from
}

func (m *message) GetTo() string {
  return m.to
}

func (m *message) GetText() string {
  return m.text
}

func (m *message) String() string {
  return "Message{"+"'to':"+m.to[0]+"}"
}
