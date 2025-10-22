package logger

import "log"
type Logger interface {
	Info(args ...interface{})
	Error(args ...interface{})
}

type stdLogger struct{}

func New(_ interface{}) Logger { return &stdLogger{} }

func (*stdLogger) Info(args ...interface{}) { log.Println(args...) }

func (*stdLogger) Error(args ...interface{}) { log.Println(args...) }
