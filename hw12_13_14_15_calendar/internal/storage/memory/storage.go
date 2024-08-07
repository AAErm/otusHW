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
	events map[storage.ID]storage.Event

	mu sync.RWMutex
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

func (s *Storage) Add(ctx context.Context, event storage.Event) (storage.ID, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	default:
	}

	event.ID = storage.ID(len(s.events) + 1)
	if s.isDateBusy(event.DateAt, storage.ID(0)) {
		return 0, storage.ErrDateBusy
	}

	s.events[event.ID] = event

	return event.ID, nil
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

func (s *Storage) ListEventByDay(ctx context.Context, t time.Time) ([]storage.Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	res := make([]storage.Event, 0, len(s.events))
	for _, event := range s.events {
		if isSameDay(event.DateTo, t) {
			res = append(res, event)
		}
	}

	return res, nil
}

func (s *Storage) ListEventsByWeek(ctx context.Context, t time.Time) ([]storage.Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	res := make([]storage.Event, 0, len(s.events))
	for _, event := range s.events {
		if isSameWeek(event.DateTo, t) {
			res = append(res, event)
		}
	}

	return res, nil
}

func (s *Storage) ListEventsByMonth(ctx context.Context, t time.Time) ([]storage.Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	res := make([]storage.Event, 0, len(s.events))
	for _, event := range s.events {
		if isSameMonth(event.DateTo, t) {
			res = append(res, event)
		}
	}

	return res, nil
}

func (s *Storage) GetEventsForNotification(ctx context.Context, start time.Time, end time.Time) ([]storage.Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	res := make([]storage.Event, 0, len(s.events))
	for _, event := range s.events {
		timeNotification := event.DateAt
		if event.NotificationTime != nil {
			timeNotification = *event.NotificationTime
		}

		if timeNotification.After(start) && timeNotification.Before(end) {
			res = append(res, event)
		}
	}

	return res, nil
}

func (s *Storage) DeleteExpiredEvents(ctx context.Context) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	for k, event := range s.events {
		if event.DateTo.Before(time.Now().AddDate(1, 0, 0)) {
			delete(s.events, k)
		}
	}

	return nil
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
		if eventID != event.ID && event.DateAt.Equal(t) {
			return true
		}
	}

	return false
}

func (s *Storage) Connect(context.Context) error {
	return nil
}

func (s *Storage) Close(context.Context) error {
	return nil
}

func (s *Storage) AddNotification(context.Context, storage.ID) error {
	return nil
}
