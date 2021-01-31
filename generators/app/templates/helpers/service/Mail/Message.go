package Mail

type Message interface {
  From(address string) Message
  // Sender(address string) Message
  To(address string) Message
  Cc(address string) Message
  Bcc(address string) Message
  ReplyTo(address string) Message
  Subject(subject string) Message
  Priority(level int) Message
  // Attach(file string, options []int) Message
  // AttachData(data interface, options []int) Message

  String() string
}

type message struct {
  from string
  to string
  cc string
  bcc string
  replyTo string
  subject string
  priority int
}

func NewMessage() Message {
  return &message{}
}

func (m *message) From(address string) Message {
  m.from = address
  return m
}

func (m *message) To(address string) Message {
  m.to = address
  return m
}

func (m *message) Cc(address string) Message {
  m.cc = address
  return m
}

func (m *message) Bcc(address string) Message {
  m.bcc = address
  return m
}

func (m *message) ReplyTo(address string) Message {
  m.replyTo = address
  return m
}

func (m *message) Subject(subject string) Message {
  m.subject = subject
  return m
}

func (m *message) Priority(level int) Message {
  m.priority = level
  return m
}

func (m *message) String() string {
  return "Message{"+"'to':"+m.to+"}"
}
