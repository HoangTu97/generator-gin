package Logger

type stack struct {
  channels []Logger
}

func NewStack(channels []Logger) *stack {
  return &stack{channels: channels}
}

func (l *stack) Debug(v ...interface{}) {
  for _, channel := range l.channels {
    channel.Debug(v)
  }
}

func (l *stack) Info(v ...interface{}) {
  for _, channel := range l.channels {
    channel.Debug(v)
  }
}

func (l *stack) Warn(v ...interface{}) {
  for _, channel := range l.channels {
    channel.Warn(v)
  }
}

func (l *stack) Error(v ...interface{}) {
  for _, channel := range l.channels {
    channel.Error(v)
  }
}

func (l *stack) Fatal(v ...interface{}) {
  for _, channel := range l.channels {
    channel.Fatal(v)
  }
}

func (l *stack) Close() {}
