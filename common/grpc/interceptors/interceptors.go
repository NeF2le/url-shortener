package interceptors

import (
	"context"
	"github.com/NeF2le/url-shortener/common/logger"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"log/slog"
	"time"
)

func AddLogMiddleware(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	ctx = logger.New(ctx)
	ctx = context.WithValue(ctx, logger.KeyForRequestID, uuid.New().String())
	logger.GetLoggerFromCtx(ctx).Info(ctx, "gRPC request",
		slog.String("method", info.FullMethod),
		slog.String("path", info.FullMethod),
		slog.Time("request time", time.Now()),
	)
	reply, err := handler(ctx, req)
	if err != nil {
		logger.GetLoggerFromCtx(ctx).Warn(ctx, "gRPC handler returns error", logger.Err(err))
	}
	return reply, err
}
