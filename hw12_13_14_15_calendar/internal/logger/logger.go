package logger

import "fmt"

type Logger struct {
	level string
}

func New(opts ...Option) *Logger {
	logger := &Logger{}

	for _, opt := range opts {
		opt(logger)
	}

	return logger
}

func (l Logger) Info(msg string) {
	fmt.Println(msg)
}

func (l Logger) Error(msg string) {
	// TODO
}

func (l Logger) Fatalf(msg string, args ...any) {
	// TODO
}
