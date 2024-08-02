package memorystorage

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/logger"
	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/storage"
)

type Storage struct {
	logger *logger.Logger
	events map[storage.ID]storage.Event

	mu sync.RWMutex //nolint:unused
}

func New(opts ...Option) *Storage {
	s := &Storage{
		events: make(map[storage.ID]storage.Event),
	}

	for _, option := range opts {
		option(s)
	}

	return s
}

func (s *Storage) Add(ctx context.Context, event storage.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	if s.isDateBusy(event.DateAt, storage.ID("")) {
		return storage.ErrDateBusy
	}

	s.events[event.ID] = event

	return nil
}

func (s *Storage) Edit(ctx context.Context, event storage.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	_, ok := s.events[event.ID]
	if !ok {
		return storage.ErrNotFound
	}

	if s.isDateBusy(event.DateAt, event.ID) {
		return storage.ErrDateBusy
	}

	s.events[event.ID] = event

	return nil
}

func (s *Storage) Delete(ctx context.Context, eventID storage.ID) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	_, ok := s.events[eventID]
	if !ok {
		return storage.ErrNotFound
	}

	delete(s.events, eventID)

	return nil
}

func (s *Storage) ListEventByDay(ctx context.Context, t time.Time) ([]*storage.Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	res := make([]*storage.Event, 0, len(s.events))
	for _, event := range s.events {
		if isSameDay(event.DateTo, t) {
			res = append(res, &event)
		}
	}

	return res, nil
}

func (s *Storage) ListEventsForWeek(ctx context.Context, t time.Time) ([]*storage.Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	res := make([]*storage.Event, 0, len(s.events))
	for _, event := range s.events {
		if isSameWeek(event.DateTo, t) {
			res = append(res, &event)
		}
	}

	return res, nil
}

func (s *Storage) ListEventsForMonth(ctx context.Context, t time.Time) ([]*storage.Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	res := make([]*storage.Event, 0, len(s.events))
	for _, event := range s.events {
		if isSameMonth(event.DateTo, t) {
			res = append(res, &event)
		}
	}

	return res, nil
}

func isSameDay(t, day time.Time) bool {
	return t.Year() == day.Year() && t.YearDay() == day.YearDay()
}

func isSameWeek(t, day time.Time) bool {
	year1, week1 := t.ISOWeek()
	year2, week2 := day.ISOWeek()
	return year1 == year2 && week1 == week2
}

func isSameMonth(t, day time.Time) bool {
	return t.Year() == day.Year() && t.Month() == day.Month()
}

func (s *Storage) isDateBusy(t time.Time, eventID storage.ID) bool {
	for _, event := range s.events {
		fmt.Println(event.DateAt)
		fmt.Println(t)
		if eventID != event.ID && event.DateAt.Equal(t) {
			return true
		}
	}

	return false
}
