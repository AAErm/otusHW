package internalgrpc

import (
	"net"

	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/app"
	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/logger"
	"google.golang.org/grpc"
)

type Server interface {
	Start() error
}

type server struct {
	server *grpc.Server
	logger *logger.Logger
	app    app.App
	host   string
}

func New(opts ...Option) *server {
	s := &server{}

	for _, opt := range opts {
		opt(s)
	}
	s.server = grpc.NewServer(grpc.ChainUnaryInterceptor(s.loggingMiddleware))

	return s
}

func (s *server) Start() error {
	l, err := net.Listen("tcp", s.host)
	if err != nil {
		return err
	}

	return s.server.Serve(l)
}

func (s *server) Stop() error {
	s.server.GracefulStop()
	return nil
}
