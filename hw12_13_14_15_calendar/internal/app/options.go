package app

import (
	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/logger"
	memorystorage "github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/storage/memory"
)

type Option func(s *App)

func WithLogger(logger *logger.Logger) Option {
	return func(s *App) {
		s.logger = logger
	}
}

func WithStorage(storage *memorystorage.Storage) Option {
	return func(s *App) {
		s.storage = storage
	}
}
