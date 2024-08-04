package logger

import (
	"github.com/sirupsen/logrus"
)

type Logger struct {
	level  string
	logger *logrus.Entry
}

func New(opts ...Option) *Logger {
	l := &Logger{}
	for _, option := range opts {
		option(l)
	}

	logger := logrus.New()
	logLevel, err := logrus.ParseLevel(l.level)
	if err != nil {
		logLevel = logrus.InfoLevel
	}

	logger.SetLevel(logLevel)
	logger.SetFormatter(&logrus.JSONFormatter{
		DisableTimestamp: true,
	})

	return &Logger{
		logger: logrus.NewEntry(logger).WithFields(logrus.Fields{}),
	}
}

func (l Logger) Debug(msg string) {
	l.logger.Debug(msg)
}

func (l Logger) Debugf(msg string, args ...any) {
	l.logger.Debugf(msg, args...)
}

func (l Logger) Info(msg string, args ...any) {
	l.logger.Infof(msg, args...)
}

func (l Logger) Infof(msg string, args ...any) {
	l.logger.Infof(msg, args...)
}

func (l Logger) Warn(msg string) {
	l.logger.Warn(msg)
}

func (l Logger) Error(msg string) {
	l.logger.Error(msg)
}

func (l Logger) Errorf(msg string, args ...any) {
	l.logger.Errorf(msg, args...)
}

func (l Logger) Fatalf(msg string, args ...any) {
	l.logger.Fatalf(msg, args...)
}
