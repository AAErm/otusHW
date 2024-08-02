package app

import (
	"context"

	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/logger"
	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/storage"
)

type App struct {
	logger  *logger.Logger
	storage *storage.Storage
}

func New(opts ...Option) *App {
	app := &App{}

	for _, opt := range opts {
		opt(app)
	}

	return app
}

func (a *App) CreateEvent(ctx context.Context, id, title string) error {
	// TODO
	return nil
	// return a.storage.CreateEvent(storage.Event{ID: id, Title: title})
}

// TODO
