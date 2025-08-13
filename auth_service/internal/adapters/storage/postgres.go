package storage

import (
	"context"
	"errors"
	"fmt"
	"github.com/NeF2le/url-shortener/auth_service/internal/models"
	errs "github.com/NeF2le/url-shortener/common/errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthPostgresAdapter struct {
	pool *pgxpool.Pool
}

func NewAuthPostgresAdapter(pool *pgxpool.Pool) *AuthPostgresAdapter {
	return &AuthPostgresAdapter{pool: pool}
}

func (a *AuthPostgresAdapter) RegisterUser(ctx context.Context, email string, passHash []byte) (string, error) {
	query := `INSERT INTO users.users (email, password_hash) VALUES ($1, $2) RETURNING id`

	var userID string
	err := a.pool.QueryRow(ctx, query, email, passHash).Scan(&userID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return "", errs.ErrUserAlreadyExists
		}
		return "", fmt.Errorf("error inserting user: %w", err)
	}

	return userID, nil
}

func (a *AuthPostgresAdapter) LoginUser(ctx context.Context, email string) (models.User, error) {
	query := `SELECT id, email, password_hash FROM users.users WHERE email = $1`

	var user models.User
	err := a.pool.QueryRow(ctx, query, email).Scan(&user.ID, &user.Email, &user.PassHash)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return user, errs.ErrUserNotFound
		}
		return user, fmt.Errorf("error getting user: %w", err)
	}

	return user, nil
}
