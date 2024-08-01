package sqlstorage

import (
	"context"
	"time"

	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/logger"
	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/storage"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	logger *logger.Logger
	conn   *pgxpool.Pool
}

func New(opts ...Option) *Storage {
	s := &Storage{}

	for _, option := range opts {
		option(s)
	}

	return s
}

func (s *Storage) Connect(ctx context.Context) error {
	// TODO
	return nil
}

func (s *Storage) Close(ctx context.Context) error {
	// TODO
	return nil
}

func (s *Storage) Add(ctx context.Context, event storage.Event) error {
	// TODO
	return nil
}

func (s *Storage) Edit(ctx context.Context, event storage.Event) error {
	// TODO
	return nil
}

func (s *Storage) Delete(ctx context.Context, event storage.Event) error {
	// TODO
	return nil
}

func (s *Storage) ListEventByDay(context.Context, time.Time) ([]*storage.Event, error) {
	return nil, nil
}

func (s *Storage) ListEventsForWeek(context.Context, time.Time) ([]*storage.Event, error) {
	return nil, nil
}
func (s *Storage) ListEventsForMonth(context.Context, time.Time) ([]*storage.Event, error) {
	return nil, nil
}
