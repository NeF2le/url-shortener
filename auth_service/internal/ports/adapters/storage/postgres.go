package storage

import (
	"context"
	"github.com/NeF2le/url-shortener/auth_service/internal/models"
)

type AuthPostgresAdapter struct {
}

func NewAuthPostgresAdapter() *AuthPostgresAdapter {
	return &AuthPostgresAdapter{}
}

func (a *AuthPostgresAdapter) RegisterUser(ctx context.Context, email string, passHash []byte) (userId int64, err error) {
	panic("implement me")
}

func (a *AuthPostgresAdapter) LoginUser(ctx context.Context, email string) (user models.User, err error) {
	panic("implement me")
}

func (a *AuthPostgresAdapter) IsAdminUser(ctx context.Context, userId int64) (isAdmin bool, err error) {
	panic("implement me")
}
