package internalhttp

import (
	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/app"
	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/logger"
	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/storage"
)

type Option func(s *server)

func WithLogger(logger *logger.Logger) Option {
	return func(s *server) {
		s.logger = logger
	}
}

func WithApplication(app app.App) Option {
	return func(s *server) {
		s.app = app
	}
}

func WithStorage(storage storage.Storage) Option {
	return func(s *server) {
		s.storage = storage
	}
}

func WithHost(host string) Option {
	return func(s *server) {
		s.host = host
	}
}

func WithPort(port int64) Option {
	return func(s *server) {
		s.port = port
	}
}
