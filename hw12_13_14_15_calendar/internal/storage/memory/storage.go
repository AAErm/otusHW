package memorystorage

import (
	"context"
	"sync"
	"time"

	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/logger"
	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/storage"
)

type Storage struct {
	logger *logger.Logger

	mu sync.RWMutex //nolint:unused
}

func New(opts ...Option) *Storage {
	s := &Storage{}

	for _, option := range opts {
		option(s)
	}

	return s
}

// TODO

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
