package Logger

import (
  "log"
)

type system struct {
  callerDepth int
}

func NewSystem(callerDepth int) *system {
  return &system{callerDepth: callerDepth}
}

func (l *system) Debug(v ...interface{}) {
  log.Println(append([]interface{}{getPrefix(DEBUG, l.callerDepth)}, v...)...)
}

func (l *system) Info(v ...interface{}) {
  log.Println(append([]interface{}{getPrefix(INFO, l.callerDepth)}, v...)...)
}

func (l *system) Warn(v ...interface{}) {
  log.Println(append([]interface{}{getPrefix(WARNING, l.callerDepth)}, v...)...)
}

func (l *system) Error(v ...interface{}) {
  log.Println(append([]interface{}{getPrefix(ERROR, l.callerDepth)}, v...)...)
}

func (l *system) Fatal(v ...interface{}) {
  log.Fatalln(append([]interface{}{getPrefix(FATAL, l.callerDepth)}, v...)...)
}

func (l *system) Close() {}
