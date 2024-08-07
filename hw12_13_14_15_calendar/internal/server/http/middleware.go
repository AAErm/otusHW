package internalhttp

import (
	"net/http"
	"time"
)

type middlewareLog struct {
	ClientIP    string
	ReqTime     time.Time
	MethodHTTP  string
	VersionHTTP string
	PathReq     string
	RespCode    int
	Latency     time.Duration
	UserAgent   string
}

func (s *server) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rw := NewResponseWriter(w)
		next.ServeHTTP(rw, r)
		mv := middlewareLog{
			ClientIP:    r.RemoteAddr,
			ReqTime:     start,
			MethodHTTP:  r.Method,
			VersionHTTP: r.Proto,
			PathReq:     r.RequestURI,
			UserAgent:   r.UserAgent(),
			Latency:     time.Since(start),
			RespCode:    rw.StatusCode,
		}
		s.logger.Info("%+v", mv)
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
