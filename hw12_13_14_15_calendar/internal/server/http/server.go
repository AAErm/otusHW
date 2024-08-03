package internalhttp

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/app"
	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/logger"
	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/storage"
	"github.com/gorilla/mux"
)

type server struct {
	logger  *logger.Logger
	app     app.App
	storage storage.Storage

	host   string
	port   int64
	server *http.Server
}

type Server interface {
	Start(context.Context) error
	Stop(context.Context) error
}

func NewServer(opts ...Option) *server {
	s := &server{}

	for _, option := range opts {
		option(s)
	}

	return s
}

func (s *server) Start(ctx context.Context) error {
	router := mux.NewRouter()
	router.HandleFunc("/", s.hello)

	handlerWithMiddleware := loggingMiddleware(router, s.logger)

	s.server = &http.Server{
		Addr:              fmt.Sprintf("%s:%d", s.host, s.port),
		Handler:           handlerWithMiddleware,
		ReadHeaderTimeout: 10 * time.Second,
	}

	go func() {
		<-ctx.Done()
		s.Stop(ctx)
	}()

	s.logger.Info(fmt.Sprintf("Server started at %s:%d", s.host, s.port))
	return s.server.ListenAndServe()
}

func (s *server) Stop(ctx context.Context) error {
	if s.server != nil {
		s.logger.Info("Shutting down the server...")
		return s.server.Shutdown(ctx)
	}
	return nil
}

func (s *server) hello(w http.ResponseWriter, r *http.Request) {
	_ = r
	fmt.Fprintln(w, "Hello, World!")
}
