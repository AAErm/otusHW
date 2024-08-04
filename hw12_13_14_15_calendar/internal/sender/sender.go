package sender

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/logger"
	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/storage"
	"github.com/imega/mt"
)

type Sender interface {
	Send()
}

type sender struct {
	logger  *logger.Logger
	storage storage.Storage
}

func New(opts ...Option) *sender {
	s := &sender{}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func (s *sender) Send() func(r *mt.Request) error {
	return func(r *mt.Request) error {
		s.logger.Info("event %s", string(r.Body))

		var event storage.Event
		if err := json.Unmarshal(r.Body, &event); err != nil {
			return fmt.Errorf("failed to unmarshal event body %s with err %v", string(r.Body), err)
		}

		if err := s.storage.AddNotification(context.Background(), event.ID); err != nil {
			return fmt.Errorf("failed to add notification to event %d with error %w", event.ID, err)
		}

		return nil
	}
}
