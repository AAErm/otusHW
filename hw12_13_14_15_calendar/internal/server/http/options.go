package internalhttp

import (
	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/app"
	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/logger"
)

type Option func(s *Server)

func WithLogger(logger *logger.Logger) Option {
	return func(s *Server) {
		s.logger = logger
	}
}

func WithApplication(app *app.App) Option {
	return func(s *Server) {
		s.app = app
	}
}
