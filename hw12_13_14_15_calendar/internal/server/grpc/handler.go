package internalgrpc

import (
	"context"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/api/proto"
	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/storage"
)

func (s *server) ListEventsByDay(ctx context.Context, req *proto.ListEventsRequest) (*proto.ListEventsResponse, error) {
	events, err := s.app.ListEventsByDay(ctx, time.Unix(req.GetTime(), 0))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	pbevs := prepareListEventsResponse(events)

	return &proto.ListEventsResponse{Events: pbevs}, nil
}

func (s *server) ListEventsByWeak(ctx context.Context, req *proto.ListEventsRequest) (*proto.ListEventsResponse, error) {
	events, err := s.app.ListEventsByWeek(ctx, time.Unix(req.GetTime(), 0))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	pbevs := prepareListEventsResponse(events)

	return &proto.ListEventsResponse{Events: pbevs}, nil
}

func (s *server) ListEventsByMonth(ctx context.Context, req *proto.ListEventsRequest) (*proto.ListEventsResponse, error) {
	events, err := s.app.ListEventsByMonth(ctx, time.Unix(req.GetTime(), 0))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	pbevs := prepareListEventsResponse(events)

	return &proto.ListEventsResponse{Events: pbevs}, nil
}

func prepareListEventsResponse(events []storage.Event) []*proto.EventResponse {
	pbevs := make([]*proto.EventResponse, 0, len(events))
	for _, e := range events {
		pbevs = append(pbevs, &proto.EventResponse{
			Id:                  int64(e.ID),
			Title:               e.Title,
			Description:         *e.Description,
			UserId:              int64(e.UserID),
			DateAt:              e.DateAt.Unix(),
			DateTo:              e.DateTo.Unix(),
			NotificationAdvance: e.NotificationAdvance.Unix(),
		})
	}

	return pbevs
}

func (s *server) CreateEvent(ctx context.Context, req *proto.CreateEventRequest) (*proto.EditEventResponse, error) {
	desc := req.GetDescription()
	notificationAdvance := time.Unix(req.GetNotificationAdvance(), 0)
	e := storage.Event{
		Title:               req.GetTitle(),
		Description:         &desc,
		UserID:              storage.ID(req.GetUserId()),
		DateAt:              time.Unix(req.GetDateAt(), 0),
		DateTo:              time.Unix(req.GetDateTo(), 0),
		NotificationAdvance: &notificationAdvance,
	}

	id, err := s.app.CreateEvent(ctx, e)
	if err != nil {
		sc := codes.Internal

		return nil, status.Error(sc, err.Error())
	}

	return &proto.EditEventResponse{Id: int64(id)}, nil
}

func (s *server) UpdateEvent(ctx context.Context, req *proto.UpdateEventRequest) (*proto.EditEventResponse, error) {
	desc := req.GetDescription()
	notificationAdvance := time.Unix(req.GetNotificationAdvance(), 0)
	e := storage.Event{
		ID:                  storage.ID(req.GetId()),
		Title:               req.GetTitle(),
		Description:         &desc,
		UserID:              storage.ID(req.GetUserId()),
		DateAt:              time.Unix(req.GetDateAt(), 0),
		DateTo:              time.Unix(req.GetDateTo(), 0),
		NotificationAdvance: &notificationAdvance,
	}

	err := s.app.UpdateEvent(ctx, e)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.EditEventResponse{Id: int64(e.ID)}, nil
}

func (s *server) DeleteEvent(ctx context.Context, req *proto.DeleteEventRequest) (*proto.DeleteEventResponse, error) {
	err := s.app.DeleteEvent(ctx, storage.ID(req.GetId()))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.DeleteEventResponse{Result: "deleted"}, nil
}
