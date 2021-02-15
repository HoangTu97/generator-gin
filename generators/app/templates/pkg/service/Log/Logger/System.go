package Logger

import (
  "fmt"
  "os"
)

type system struct {
  callerDepth int
}

func NewSystem(callerDepth int) *system {
  return &system{callerDepth: callerDepth}
}

func (l *system) Debug(v ...interface{}) {
  fmt.Println(append([]interface{}{getPrefix(DEBUG, l.callerDepth)}, v...)...)
}

func (l *system) Info(v ...interface{}) {
  fmt.Println(append([]interface{}{getPrefix(INFO, l.callerDepth)}, v...)...)
}

func (l *system) Warn(v ...interface{}) {
  fmt.Println(append([]interface{}{getPrefix(WARNING, l.callerDepth)}, v...)...)
}

func (l *system) Error(v ...interface{}) {
  fmt.Println(append([]interface{}{getPrefix(ERROR, l.callerDepth)}, v...)...)
}

func (l *system) Fatal(v ...interface{}) {
  fmt.Println(append([]interface{}{getPrefix(FATAL, l.callerDepth)}, v...)...)
  os.Exit(1)
}

func (l *system) Close() {}
