package internalhttp

import (
	"context"
)

type Server struct {
	logger Logger
	app    Application
}

type Logger interface { // TODO
}

type Application interface { // TODO
}

func NewServer(opts ...Option) *Server {
	s := &Server{}

	for _, option := range opts {
		option(s)
	}

	return s
}

func (s *Server) Start(ctx context.Context) error {
	// TODO
	<-ctx.Done()
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	// TODO
	return nil
}

// TODO
