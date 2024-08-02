package sqlstorage

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/logger"
	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/storage"
	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/storage/sql/query"
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
	query, args := query.BuildAddEventQuery(event)
	_, err := s.conn.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to add event %s", err)
	}
	return nil
}

func (s *Storage) Edit(ctx context.Context, event storage.Event) error {
	query, args := query.BuildEditEventQuery(event)
	_, err := s.conn.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to edit event %s", err)
	}
	return nil
}

func (s *Storage) Delete(ctx context.Context, ID storage.ID) error {
	query, args := query.BuildDeleteEventQuery(ID)
	_, err := s.conn.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to delete event %s", err)
	}
	return nil
}

func (s *Storage) ListEventByDay(ctx context.Context, t time.Time) ([]*storage.Event, error) {
	query, args := query.BuildListEventByDay(t)
	return s.getList(ctx, query, args)
}

func (s *Storage) ListEventsForWeek(ctx context.Context, t time.Time) ([]*storage.Event, error) {
	query, args := query.BuildListEventByWeak(t)
	return s.getList(ctx, query, args)
}
func (s *Storage) ListEventsForMonth(ctx context.Context, t time.Time) ([]*storage.Event, error) {
	query, args := query.BuildListEventByWeak(t)
	return s.getList(ctx, query, args)
}

func (s *Storage) getList(ctx context.Context, query string, args []any) ([]*storage.Event, error) {
	rows, err := s.conn.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query %s with error %s", query, err.Error())
	}

	var rr []*storage.Event
	for rows.Next() {
		var r storage.Event
		if err := rows.Scan(
			&r.ID, &r.UserID, &r.Title, &r.DateAt,
			&r.DateTo, &r.Description, &r.NotificationAdvance,
		); err != nil && err != sql.ErrNoRows {
			return nil, fmt.Errorf("failed to get list event with error %w", err)
		}

		rr = append(rr, &r)
	}

	return rr, nil
}
