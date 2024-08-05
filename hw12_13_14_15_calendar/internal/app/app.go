package app

import (
	"context"
	"time"

	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/logger"
	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/storage"
)

type App interface {
	CreateEvent(ctx context.Context, event storage.Event) (storage.ID, error)
	UpdateEvent(ctx context.Context, event storage.Event) error
	DeleteEvent(ctx context.Context, eventID storage.ID) error
	ListEventsByDay(ctx context.Context, day time.Time) ([]storage.Event, error)
	ListEventsByWeek(ctx context.Context, weak time.Time) ([]storage.Event, error)
	ListEventsByMonth(ctx context.Context, month time.Time) ([]storage.Event, error)
}

type app struct {
	logger  *logger.Logger
	storage storage.Storage
}

func New(opts ...Option) *app {
	app := &app{}

	for _, opt := range opts {
		opt(app)
	}

	return app
}

func (a *app) CreateEvent(ctx context.Context, event storage.Event) (storage.ID, error) {
	return a.storage.Add(ctx, event)
}

func (a *app) ListEventsByDay(ctx context.Context, day time.Time) ([]storage.Event, error) {
	return a.storage.ListEventByDay(ctx, day)
}

func (a *app) ListEventsByWeek(ctx context.Context, weak time.Time) ([]storage.Event, error) {
	return a.storage.ListEventsByWeek(ctx, weak)
}

func (a *app) ListEventsByMonth(ctx context.Context, month time.Time) ([]storage.Event, error) {
	return a.storage.ListEventsByMonth(ctx, month)
}

func (a *app) UpdateEvent(ctx context.Context, event storage.Event) error {
	return a.storage.Edit(ctx, event)
}

func (a *app) DeleteEvent(ctx context.Context, eventID storage.ID) error {
	return a.storage.Delete(ctx, eventID)
}
