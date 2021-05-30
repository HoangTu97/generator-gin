package Logger

type null struct {}

func NewNull() *null {
  return &null{}
}

func (l *null) Debug(v ...interface{}) {}
func (l *null) Info(v ...interface{}) {}
func (l *null) Warn(v ...interface{}) {}
func (l *null) Error(v ...interface{}) {}
func (l *null) Fatal(v ...interface{}) {}
func (l *null) Close() {}
