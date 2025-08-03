package server

import (
	"context"
	"errors"
	"github.com/NeF2le/url-shortener/auth_service/internal/ports"
	errs "github.com/NeF2le/url-shortener/common/errors"
	"github.com/NeF2le/url-shortener/common/gen/auth_service"
	"github.com/NeF2le/url-shortener/common/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

type grpcAuthServer struct {
	auth ports.AuthUseCase
	auth_service.UnimplementedAuthServiceServer
}

func NewGRPCAuthServer(useCase ports.AuthUseCase) auth_service.AuthServiceServer {
	return &grpcAuthServer{auth: useCase}
}

func (s *grpcAuthServer) Register(ctx context.Context, req *auth_service.RegisterRequest) (*auth_service.RegisterResponse, error) {
	if req.GetEmail() == "" {
		return nil, errs.ErrEmptyEmail
	}
	if req.GetPassword() == "" {
		return nil, errs.ErrEmptyPassword
	}

	userId, err := s.auth.Register(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		if errors.Is(err, errs.ErrUserAlreadyExists) {
			logger.GetLoggerFromCtx(ctx).Warn(ctx,
				"user already exists",
				slog.String("email", req.GetEmail()),
				logger.Err(err),
			)
			return nil, errs.ErrUserAlreadyExists
		}
		logger.GetLoggerFromCtx(ctx).Error(ctx,
			"failed to register user",
			slog.String("email", req.GetEmail()),
			logger.Err(err),
		)
		return nil, status.Errorf(codes.Unauthenticated, "%v", err)
	}

	logger.GetLoggerFromCtx(ctx).Info(ctx,
		"user registered successfully",
		slog.String("email", req.GetEmail()),
		slog.Int64("userId", userId),
	)
	return &auth_service.RegisterResponse{UserId: userId}, nil
}

func (s *grpcAuthServer) Login(ctx context.Context, req *auth_service.LoginRequest) (*auth_service.LoginResponse, error) {
	if req.GetEmail() == "" {
		return nil, errs.ErrEmptyEmail
	}
	if req.GetPassword() == "" {
		return nil, errs.ErrEmptyPassword
	}

	token, err := s.auth.Login(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		if errors.Is(err, errs.ErrInvalidCredentials) {
			logger.GetLoggerFromCtx(ctx).Warn(ctx,
				"invalid credentials",
				slog.String("email", req.GetEmail()),
				logger.Err(err),
			)
			return nil, errs.ErrInvalidCredentials
		}
		logger.GetLoggerFromCtx(ctx).Error(ctx,
			"failed to login user",
			slog.String("email", req.GetEmail()),
			logger.Err(err),
		)
		return nil, status.Errorf(codes.Unauthenticated, "%v", err)
	}

	logger.GetLoggerFromCtx(ctx).Info(ctx,
		"user logged in successfully",
		slog.String("email", req.GetEmail()),
		slog.String("token", token),
	)
	return &auth_service.LoginResponse{Token: token}, nil
}
