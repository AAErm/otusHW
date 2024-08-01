package sqlstorage

import (
	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/logger"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Option func(s *Storage)

func WithLogger(logger *logger.Logger) Option {
	return func(s *Storage) {
		s.logger = logger
	}
}

func WithConnect(conn *pgxpool.Pool) Option {
	return func(cl *Storage) {
		cl.conn = conn
	}
}
