package storage

import (
	"context"
	"errors"
	"time"
)

// * добавление события в хранилище;
// * изменение события в хранилище;
// * удаление события из хранилища;
// * СписокСобытийНаДень (дата);
// * СписокСобытийНаНеделю (дата начала недели);
// * СписокСобытийНaМесяц (дата начала месяца).
type Storage interface {
	Connect(context.Context) error
	Close(context.Context) error
	Add(context.Context, Event) (ID, error)
	Edit(context.Context, Event) error
	Delete(context.Context, ID) error
	ListEventByDay(context.Context, time.Time) ([]Event, error)
	ListEventsByWeek(context.Context, time.Time) ([]Event, error)
	ListEventsByMonth(context.Context, time.Time) ([]Event, error)
	GetEventsForNotification(context.Context, time.Time, time.Time) ([]Event, error)
	DeleteExpiredEvents(context.Context) error
}

var (
	ErrDateBusy = errors.New("DateBusy")
	ErrNotFound = errors.New("NotFound")
)
