package memorystorage

import (
	"context"
	"testing"
	"time"

	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/logger"
	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/storage"
)

func TestStorage_Add(t *testing.T) {
	type fields struct {
		logger *logger.Logger
		events map[storage.ID]storage.Event
	}
	type args struct {
		ctx   context.Context
		event storage.Event
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		wantErr          bool
		wantAmountEvents int
	}{
		{
			name: "add event is dateBusy",
			fields: fields{
				logger: nil,
				events: map[storage.ID]storage.Event{
					storage.ID(1): {
						ID:     storage.ID(1),
						Title:  "some title",
						DateAt: time.Now().Truncate(time.Second),
						DateTo: time.Now().Add(time.Hour),
						UserID: storage.ID(1),
					},
				},
			},
			args: args{
				ctx: context.TODO(),
				event: storage.Event{
					ID:     storage.ID(2),
					Title:  "some title",
					DateAt: time.Now().Truncate(time.Second),
					DateTo: time.Now().Add(time.Hour),
					UserID: storage.ID(1),
				},
			},
			wantErr:          true,
			wantAmountEvents: 1,
		},
		{
			name: "correct add event",
			fields: fields{
				logger: nil,
				events: map[storage.ID]storage.Event{
					storage.ID(1): {
						ID:     storage.ID(1),
						Title:  "some title",
						DateAt: time.Now().Truncate(time.Second),
						DateTo: time.Now().Add(time.Hour),
						UserID: storage.ID(1),
					},
				},
			},
			args: args{
				ctx: context.TODO(),
				event: storage.Event{
					ID:     storage.ID(2),
					Title:  "some title",
					DateAt: time.Now().Add(time.Hour).Truncate(time.Second),
					DateTo: time.Now().Add(time.Hour),
					UserID: storage.ID(1),
				},
			},
			wantErr:          false,
			wantAmountEvents: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				logger: tt.fields.logger,
				events: tt.fields.events,
			}
			if err := s.Add(tt.args.ctx, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("Storage.Add() error = %v, wantErr %v", err, tt.wantErr)
			} else if len(tt.fields.events) != tt.wantAmountEvents {
				t.Errorf("Storage.Add() want %d events, have %d", tt.wantAmountEvents, len(tt.fields.events))
			}
		})
	}
}

func TestStorage_Edit(t *testing.T) {
	type fields struct {
		logger *logger.Logger
		events map[storage.ID]storage.Event
	}
	type args struct {
		ctx   context.Context
		event storage.Event
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "edit event. event not found",
			fields: fields{
				logger: nil,
				events: map[storage.ID]storage.Event{
					storage.ID(1): {
						ID:     storage.ID(1),
						Title:  "some title",
						DateAt: time.Now().Truncate(time.Second),
						DateTo: time.Now().Add(time.Hour),
						UserID: storage.ID(1),
					},
				},
			},
			args: args{
				ctx: context.TODO(),
				event: storage.Event{
					ID:     storage.ID(2),
					Title:  "some title",
					DateAt: time.Now().Truncate(time.Second),
					DateTo: time.Now().Add(time.Hour),
					UserID: storage.ID(1),
				},
			},
			wantErr: true,
		},
		{
			name: "edit event. correct edit title",
			fields: fields{
				logger: nil,
				events: map[storage.ID]storage.Event{
					storage.ID(1): {
						ID:     storage.ID(1),
						Title:  "some title",
						DateAt: time.Now().Truncate(time.Second),
						DateTo: time.Now().Add(time.Hour),
						UserID: storage.ID(1),
					},
				},
			},
			args: args{
				ctx: context.TODO(),
				event: storage.Event{
					ID:     storage.ID(1),
					Title:  "some title2",
					DateAt: time.Now().Truncate(time.Second),
					DateTo: time.Now().Add(time.Hour),
					UserID: storage.ID(1),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				logger: tt.fields.logger,
				events: tt.fields.events,
			}
			if err := s.Edit(tt.args.ctx, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("Storage.Edit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
