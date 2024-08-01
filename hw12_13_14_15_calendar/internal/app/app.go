package app

import (
	"context"

	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/storage"
)

type App struct {
	logger  Logger
	storage *storage.Storage
}

type Logger interface { // TODO
}

type Storage interface { // TODO
}

func New(logger Logger, storage Storage) *App {
	return &App{}
}

func (a *App) CreateEvent(ctx context.Context, id, title string) error {
	// TODO
	return nil
	// return a.storage.CreateEvent(storage.Event{ID: id, Title: title})
}

// TODO
