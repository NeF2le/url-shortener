package ports

import "context"

type AuthUseCase interface {
	Register(ctx context.Context, email string, pass string) (string, error)
	Login(ctx context.Context, email string, pass string) (string, string, string, error)
	Refresh(ctx context.Context, refreshToken string) (string, string, error)
}
