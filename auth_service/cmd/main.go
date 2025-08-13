package main

import (
	"context"
	"github.com/NeF2le/url-shortener/auth_service/internal/adapters/cache"
	"github.com/NeF2le/url-shortener/auth_service/internal/adapters/storage"
	"github.com/NeF2le/url-shortener/auth_service/internal/config"
	"github.com/NeF2le/url-shortener/auth_service/internal/server"
	"github.com/NeF2le/url-shortener/auth_service/internal/service"
	"github.com/NeF2le/url-shortener/common/logger"
	"github.com/NeF2le/url-shortener/common/redis"
	"github.com/NeF2le/url-shortener/common/storage/postgres"
	"log"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.WithValue(context.Background(), logger.KeyForLogLevel, cfg.LogLevel)
	ctx = logger.New(ctx)

	postgresClient, err := postgres.NewPostgresClient(ctx, &cfg.Postgres)
	if err != nil {
		log.Fatal(err)
	}
	redisClient, err := redis.NewRedisClient(ctx, &cfg.Redis, cfg.AuthService.RedisDB)
	if err != nil {
		log.Fatal(err)
	}

	err = postgres.Migrate(&cfg.Postgres, cfg.MigrationPath, "migrations")
	if err != nil {
		log.Fatal(err)
	}

	storageAdapter := storage.NewAuthPostgresAdapter(postgresClient)
	cacheAdapter := cache.NewAuthRedisAdapter(redisClient, cfg.RefreshExpiration, cfg.AccessExpiration)

	authService := service.NewAuthService(
		storageAdapter,
		cacheAdapter,
		cfg.JwtSecret,
		cfg.RefreshExpiration,
		cfg.AccessExpiration,
	)

	grpcAuthServer := server.NewGRPCAuthServer(authService)

	grpcServer, err := server.CreateGRPC(grpcAuthServer)
	if err != nil {
		log.Fatalf("failed to create gRPC server: %v", err)
	}

	go server.RunGRPC(ctx, grpcServer, cfg.AuthService.Port)

	<-ctx.Done()
	grpcServer.GracefulStop()
	logger.GetLoggerFromCtx(ctx).Info(ctx, "AUTH_SERVICE stopped")
}
