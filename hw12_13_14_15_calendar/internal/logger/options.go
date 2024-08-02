package logger

type Option func(*Logger)

func WithLevel(level string) Option {
	return func(l *Logger) {
		l.level = level
	}
}
