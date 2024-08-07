package sqlstorage

import (
	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/config"
	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/logger"
)

type Option func(s *Storage)

func WithLogger(logger *logger.Logger) Option {
	return func(s *Storage) {
		s.logger = logger
	}
}

func WithConfig(config config.DBConf) Option {
	return func(cl *Storage) {
		cl.conf = config
	}
}
