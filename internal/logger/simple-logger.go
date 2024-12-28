package logger

import (
	"fmt"
	"time"
)

type SimpleLogger struct{}

func (l *SimpleLogger) log(level string, message string, fields ...interface{}) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logMessage := fmt.Sprintf("[%s] [%s] %s", timestamp, level, message)
	if len(fields) > 0 {
		logMessage += fmt.Sprintf(" | Fields: %v", fields)
	}
	fmt.Println(logMessage)
}

func (l *SimpleLogger) Info(message string, fields ...interface{}) {
	l.log("INFO", message, fields...)
}

func (l *SimpleLogger) Warning(message string, fields ...interface{}) {
	l.log("WARNING", message, fields...)
}

func (l *SimpleLogger) Debug(message string, fields ...interface{}) {
	l.log("DEBUG", message, fields...)
}

func (l *SimpleLogger) Error(message string, fields ...interface{}) {
	l.log("ERROR", message, fields...)
}

func NewSimpleLogger() Logger {
	return &SimpleLogger{}
}
