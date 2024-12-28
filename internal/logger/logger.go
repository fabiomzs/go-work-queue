package logger

type Logger interface {
	Info(message string, fields ...interface{})
	Warning(message string, fields ...interface{})
	Debug(message string, fields ...interface{})
	Error(message string, fields ...interface{})
}
