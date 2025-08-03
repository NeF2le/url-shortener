package main

import (
	"context"
	"github.com/NeF2le/url-shortener/auth_service/internal/config"
	"github.com/NeF2le/url-shortener/auth_service/internal/ports/adapters/storage"
	"github.com/NeF2le/url-shortener/auth_service/internal/server"
	"github.com/NeF2le/url-shortener/auth_service/internal/service"
	"github.com/NeF2le/url-shortener/common/logger"
	"github.com/NeF2le/url-shortener/common/storage/sqlite"
	"log"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.WithValue(context.Background(), "log_level", cfg.LogLevel)
	ctx = logger.New(ctx)

	sqlitePath := "/Users/nef1le/Desktop/url-shortener/storage/sqlite/url-shortener.db"
	sqliteClient, err := sqlite.NewStorageSQLite(sqlitePath)
	if err != nil {
		log.Fatal(err)
	}

	err = sqlite.Migrate(sqlitePath, cfg.MigrationPath, "migrations")
	if err != nil {
		log.Fatal(err)
	}

	storageAdapter := storage.NewAuthSQLiteAdapter(sqliteClient)

	authService := service.NewAuthService(
		storageAdapter,
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
