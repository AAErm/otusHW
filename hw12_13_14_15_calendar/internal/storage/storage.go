package storage

import (
	"context"
	"time"
)

// * добавление события в хранилище;
// * изменение события в хранилище;
// * удаление события из хранилища;
// * СписокСобытийНаДень (дата);
// * СписокСобытийНаНеделю (дата начала недели);
// * СписокСобытийНaМесяц (дата начала месяца).
type Storage interface {
	Add(context.Context, Event) error
	Edit(context.Context, Event) error
	Delete(context.Context, Event) error
	ListEventByDay(context.Context, time.Time) ([]*Event, error)
	ListEventsForWeek(context.Context, time.Time) ([]*Event, error)
	ListEventsForMonth(context.Context, time.Time) ([]*Event, error)
}
