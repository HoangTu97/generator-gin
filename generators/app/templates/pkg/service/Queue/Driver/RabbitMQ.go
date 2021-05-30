package QueueDriver

import (
  "github.com/streadway/amqp"
)

type rabbitMq struct {
  conn *amqp.Connection
  ch   *amqp.Channel
}

func NewRabbitMQ(url string) *rabbitMq {
  conn, _ := amqp.Dial(url)
  ch, _ := conn.Channel()
  return &rabbitMq{conn: conn, ch: ch}
}

func (q *rabbitMq) Close() {
  q.ch.Close()
  q.conn.Close()
}
