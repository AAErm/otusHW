package app

import (
	"context"

	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/logger"
	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/storage"
)

type App interface {
	CreateEvent(ctx context.Context, id storage.ID, title string) error
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

func (a *app) CreateEvent(ctx context.Context, id storage.ID, title string) error {
	return a.storage.Add(ctx, storage.Event{ID: id, Title: title})
}
