package memorystorage

import (
	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/logger"
)

type Option func(s *Storage)

func WithLogger(logger *logger.Logger) Option {
	return func(s *Storage) {
		s.logger = logger
	}
}
