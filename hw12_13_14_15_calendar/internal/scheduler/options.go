package scheduler

import (
	"time"

	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/logger"
	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/storage"
	"github.com/imega/mt"
)

type Option func(s *scheduler)

func WithLogger(logger *logger.Logger) Option {
	return func(s *scheduler) {
		s.logger = logger
	}
}

func WithStorage(storage storage.Storage) Option {
	return func(s *scheduler) {
		s.storage = storage
	}
}

func WithMt(mt mt.MT) Option {
	return func(s *scheduler) {
		s.mt = mt
	}
}

func WithInterval(interval time.Duration) Option {
	return func(s *scheduler) {
		s.interval = interval
	}
}

func WithServiceName(serviceName string) Option {
	return func(s *scheduler) {
		s.serviceName = serviceName
	}
}
