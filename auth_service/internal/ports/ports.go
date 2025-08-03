package ports

import (
	"context"
	"github.com/NeF2le/url-shortener/auth_service/internal/models"
)

type StorageRepository interface {
	RegisterUser(ctx context.Context, email string, passHash []byte) (int64, error)
	LoginUser(ctx context.Context, email string) (models.User, error)
}

type AuthUseCase interface {
	Register(ctx context.Context, email string, pass string) (int64, error)
	Login(ctx context.Context, email string, pass string) (string, error)
}
