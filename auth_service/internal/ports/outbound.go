package ports

import (
	"context"
	"github.com/NeF2le/url-shortener/auth_service/internal/models"
)

type StorageRepository interface {
	RegisterUser(ctx context.Context, email string, passHash []byte) (string, error)
	LoginUser(ctx context.Context, email string) (models.User, error)
}

type CacheRepository interface {
	SaveToken(ctx context.Context, token, userID string, refresh bool) error
	GetToken(ctx context.Context, token string, refresh bool) (string, error)
	DeleteToken(ctx context.Context, token string, refresh bool) error
}
