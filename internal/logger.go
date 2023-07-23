package internal

import (
	"io"
	"log"
)

type Logger interface {
	Info(string)
	Warning(string)
	Error(string)
}

type CustomLogger struct {
	logger *log.Logger
}

func NewLogger(writer io.Writer, flags int) *CustomLogger {
	logger := log.New(writer, "", flags)
	return &CustomLogger{logger}
}

func (l *CustomLogger) Info(msg string) {
	l.logger.Printf("INFO: %s", msg)
}

func (l *CustomLogger) Warning(msg string) {
	l.logger.Printf("WARNING: %s", msg)
}

func (l *CustomLogger) Error(msg string) {
	l.logger.Printf("ERROR: %s", msg)
}
