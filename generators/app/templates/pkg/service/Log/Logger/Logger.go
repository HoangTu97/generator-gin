package Logger

import (
  "path/filepath"
  "runtime"
  "time"
  "fmt"
)

type Logger interface {
  Debug(v ...interface{})
  Info(v ...interface{})
  Warn(v ...interface{})
  Error(v ...interface{})
  Fatal(v ...interface{})
}

type Closeable interface {
  Close()
}

type Level int
const (
  DEBUG Level = iota
  INFO
  WARNING
  ERROR
  FATAL
)

var (
  levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

func getPrefix(level Level, callerDepth int) string {
  _, f, line, ok := runtime.Caller(callerDepth)
  var logPrefix string
  t := time.Now()
  if ok {
    logPrefix = fmt.Sprintf("[%d] [%s][%s:%d]", t.Unix(), levelFlags[level], filepath.Base(f), line)
  } else {
    logPrefix = fmt.Sprintf("[%d] [%s]", t.Unix(), levelFlags[level])
  }
  return logPrefix
}
