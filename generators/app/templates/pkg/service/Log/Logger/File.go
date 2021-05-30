package Logger

import (
  helperFile "<%= appName %>/helpers/file"

  "fmt"
  "log"
  "os"
  "time"
)

type file struct {
  logger      *log.Logger
  f           *os.File
  callerDepth int
}

func NewFile(runtimeRootPath, savePath, saveName, timeFormat, fileExt string, callerDepth int) *file {
  filePath := fmt.Sprintf("%s%s", runtimeRootPath, savePath)
  fileName := fmt.Sprintf("%s%s.%s", saveName, time.Now().Format(timeFormat), fileExt)
  f, err := helperFile.MustOpen(fileName, filePath)
  if err != nil {
    log.Fatalf("Logger.NewFile err: %v", err)
  }

  logger := log.New(f, "", log.LstdFlags)

  return &file{logger: logger, callerDepth: callerDepth, f: f}
}

func (l *file) Debug(v ...interface{}) {
  l.logger.Println(append([]interface{}{getPrefix(DEBUG, l.callerDepth)}, v...)...)
}

func (l *file) Info(v ...interface{}) {
  l.logger.Println(append([]interface{}{getPrefix(INFO, l.callerDepth)}, v...)...)
}

func (l *file) Warn(v ...interface{}) {
  l.logger.Println(append([]interface{}{getPrefix(WARNING, l.callerDepth)}, v...)...)
}

func (l *file) Error(v ...interface{}) {
  l.logger.Println(append([]interface{}{getPrefix(ERROR, l.callerDepth)}, v...)...)
}

func (l *file) Fatal(v ...interface{}) {
  l.logger.Fatalln(append([]interface{}{getPrefix(FATAL, l.callerDepth)}, v...)...)
}

func (l *file) Close() {
  l.f.Close()
}
