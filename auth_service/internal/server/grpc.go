package server

import (
	"context"
	"fmt"
	"github.com/NeF2le/url-shortener/common/gen/auth_service"
	"github.com/NeF2le/url-shortener/common/grpc/interceptors"
	"github.com/NeF2le/url-shortener/common/logger"
	"google.golang.org/grpc"
	"net"
)

func CreateGRPC(grpcServer auth_service.AuthServiceServer) (*grpc.Server, error) {
	server := grpc.NewServer(grpc.UnaryInterceptor(interceptors.AddLogMiddleware))
	auth_service.RegisterAuthServiceServer(server, grpcServer)
	return server, nil
}

func RunGRPC(ctx context.Context, grpcServer *grpc.Server, port int) {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		logger.GetLoggerFromCtx(ctx).Error(ctx,
			"AUTH_SERVICE failed to create listener on port",
			"port", port,
			"error", err)
	}

	logger.GetLoggerFromCtx(ctx).Info(ctx, fmt.Sprintf("AUTH_SERVICE created listener on %d", port))
	if err = grpcServer.Serve(l); err != nil {
		logger.GetLoggerFromCtx(ctx).Error(ctx,
			"AUTH_SERVICE failed to start server",
			"error", err)
	}
}
