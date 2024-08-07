package scheduler

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/logger"
	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/storage"
	"github.com/imega/mt"
)

type scheduler struct {
	logger      *logger.Logger
	mt          mt.MT
	storage     storage.Storage
	interval    time.Duration
	serviceName string
}

type Scheduler interface {
	Run()
}

func New(opts ...Option) *scheduler {
	s := &scheduler{}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func (s *scheduler) Run() func() {
	return func() {
		ctx := context.Background()
		records, err := s.GetEvents(ctx)
		if err != nil {
			s.logger.Errorf("failed to get events to notification with error %v", err)
			return
		}
		for _, record := range records {
			err = s.Cast(ctx, record)
			if err != nil {
				s.logger.Errorf("failed to cast events with error %v", err)
				return
			}
		}

		err = s.DeleteExpiredEvents(ctx)
		if err != nil {
			s.logger.Errorf("failed to delete expored events with error %v", err)
		}
	}
}

func (s *scheduler) GetEvents(ctx context.Context) ([]storage.Event, error) {
	return s.storage.GetEventsForNotification(ctx, time.Now(), time.Now().Add(s.interval))
}

func (s *scheduler) Cast(ctx context.Context, event storage.Event) error {
	b, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal event %w", err)
	}

	return s.mt.Cast(s.serviceName, mt.Request{
		Body: b,
	})
}

func (s *scheduler) DeleteExpiredEvents(ctx context.Context) error {
	return s.storage.DeleteExpiredEvents(ctx)
}
