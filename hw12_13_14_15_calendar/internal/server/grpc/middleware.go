package internalgrpc

import (
	"context"
	"time"

	"google.golang.org/grpc"
)

type middlewareLog struct {
	ReqTime    time.Time
	MethodGRPC string
	Latency    time.Duration
}

func (s *server) loggingMiddleware(
	ctx context.Context,
	req any,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (any, error) {
	res, err := handler(ctx, req)
	start := time.Now()
	mv := middlewareLog{
		ReqTime:    start,
		MethodGRPC: info.FullMethod,
		Latency:    time.Since(start),
	}
	s.logger.Info("%+v", mv)

	return res, err
}
