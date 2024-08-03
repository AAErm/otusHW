package internalhttp

import (
	"net/http"
	"time"

	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/logger"
)

type middleware struct {
	ClientIP    string
	ReqTime     time.Time
	MethodHTTP  string
	VersionHTTP string
	PathReq     string
	RespCode    int
	Latency     time.Duration
	UserAgent   string
}

func loggingMiddleware(next http.Handler, logg *logger.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rw := NewResponseWriter(w)
		next.ServeHTTP(rw, r)
		mv := middleware{
			ClientIP:    r.RemoteAddr,
			ReqTime:     start,
			MethodHTTP:  r.Method,
			VersionHTTP: r.Proto,
			PathReq:     r.RequestURI,
			UserAgent:   r.UserAgent(),
			Latency:     time.Since(start),
			RespCode:    rw.StatusCode,
		}
		logg.Info("%+v", mv)
	})
}

type CustomResponseWriter struct {
	http.ResponseWriter
	StatusCode int
}

func NewResponseWriter(writer http.ResponseWriter) *CustomResponseWriter {
	return &CustomResponseWriter{writer, 0}
}

func (lrw *CustomResponseWriter) WriteHeader(code int) {
	lrw.StatusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}
