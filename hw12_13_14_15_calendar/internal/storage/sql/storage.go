package sqlstorage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/config"
	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/logger"
	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/storage"
	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/storage/sql/query"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	logger *logger.Logger
	conf   config.DBConf
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
	conn, err := conn(ctx, s.conf)
	if err != nil {
		return err
	}

	s.conn = conn

	return nil
}

func (s *Storage) Close(context.Context) error {
	s.conn.Close()
	return nil
}

func (s *Storage) Add(ctx context.Context, event storage.Event) error {
	query, args := query.BuildAddEventQuery(event)
	_, err := s.conn.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to add event %w", err)
	}
	return nil
}

func (s *Storage) Edit(ctx context.Context, event storage.Event) error {
	query, args := query.BuildEditEventQuery(event)
	_, err := s.conn.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to edit event %w", err)
	}
	return nil
}

func (s *Storage) Delete(ctx context.Context, eventID storage.ID) error {
	query, args := query.BuildDeleteEventQuery(eventID)
	_, err := s.conn.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to delete event %w", err)
	}
	return nil
}

func (s *Storage) ListEventByDay(ctx context.Context, t time.Time) ([]storage.Event, error) {
	query, args := query.BuildListEventByDay(t)
	return s.getList(ctx, query, args)
}

func (s *Storage) ListEventsForWeek(ctx context.Context, t time.Time) ([]storage.Event, error) {
	query, args := query.BuildListEventByWeak(t)
	return s.getList(ctx, query, args)
}

func (s *Storage) ListEventsForMonth(ctx context.Context, t time.Time) ([]storage.Event, error) {
	query, args := query.BuildListEventByWeak(t)
	return s.getList(ctx, query, args)
}

func (s *Storage) getList(ctx context.Context, query string, args []any) ([]storage.Event, error) {
	rows, err := s.conn.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query %s with error %w", query, err)
	}

	var rr []storage.Event
	for rows.Next() {
		var r storage.Event
		if err := rows.Scan(
			&r.ID, &r.UserID, &r.Title, &r.DateAt,
			&r.DateTo, &r.Description, &r.NotificationAdvance,
		); err != nil && !errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("failed to get list event with error %w", err)
		}

		rr = append(rr, r)
	}

	return rr, nil
}
