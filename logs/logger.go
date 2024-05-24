package logs

type Logger interface {
	Info(msg string, args ...interface{})
	Error(msg string, args ...interface{})
}
