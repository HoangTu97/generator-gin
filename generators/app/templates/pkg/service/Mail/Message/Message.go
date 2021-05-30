package MailMessage

import (
  "bytes"
  "encoding/base64"
  "fmt"
  "mime/multipart"
)

type Message interface {
  // From(address string) Message
  // Sender(address string) Message
  To(addresses []string) Message
  Cc(address string) Message
  Bcc(address string) Message
  ReplyTo(address string) Message
  Subject(subject string) Message
  Priority(level int) Message
  Body(body string) Message
  // Attach(file string, options []int) Message
  // AttachData(data interface, options []int) Message

  GetTo() []string
  GetCc() string
  GetBcc() string
  GetSubject() string
  GetBody() string

  String() string
  ToBytes() []byte
}

type message struct {
  // from string
  to []string
  cc string
  bcc string
  replyTo string
  subject string
  priority int
  body string
  attachments map[string][]byte
}

func NewMessage() Message {
  return &message{}
}

// func (m *message) From(address string) Message {
//   m.from = address
//   return m
// }

func (m *message) To(addresses []string) Message {
  m.to = addresses
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

func (m *message) Body(body string) Message {
  m.body = body
  return m
}

func (m *message) GetTo() []string {
  return m.to
}

func (m *message) GetCc() string {
  return m.cc
}

func (m *message) GetBcc() string {
  return m.bcc
}

func (m *message) GetSubject() string {
  return m.subject
}

func (m *message) GetBody() string {
  return m.body
}

func (m *message) String() string {
  return "Message{"+"'to':"+m.to[0]+"}"
}

func (m *message) ToBytes() []byte {
  buf := bytes.NewBuffer(nil)
  withAttachments := len(m.attachments) > 0

  buf.WriteString(fmt.Sprintf("Subject: %s\n", m.subject))
  buf.WriteString("MIME-Version: 1.0\n")
  writer := multipart.NewWriter(buf)
  boundary := writer.Boundary()

  if withAttachments {
    buf.WriteString(fmt.Sprintf("Content-Type: multipart/mixed; boundary=%s\n", boundary))
    buf.WriteString(fmt.Sprintf("--%s\n", boundary))
  }

  buf.WriteString("Content-Type: text/plain; charset=utf-8\n")
  buf.WriteString(m.body)

  if withAttachments {
    for k, v := range m.attachments {
      buf.WriteString(fmt.Sprintf("\n\n--%s\n", boundary))
      buf.WriteString("Content-Type: application/octet-stream\n")
      buf.WriteString("Content-Transfer-Encoding: base64\n")
      buf.WriteString(fmt.Sprintf("Content-Disposition: attachment; filename=%s\n", k))

      b := make([]byte, base64.StdEncoding.EncodedLen(len(v)))
      base64.StdEncoding.Encode(b, v)
      buf.Write(b)
      buf.WriteString(fmt.Sprintf("\n--%s", boundary))
    }

    buf.WriteString("--")
  }

  return buf.Bytes()
}
