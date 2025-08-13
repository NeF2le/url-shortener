package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/NeF2le/url-shortener/auth_service/internal/models"
	errs "github.com/NeF2le/url-shortener/common/errors"
)

type AuthSQLiteAdapter struct {
	db *sql.DB
}

func NewAuthSQLiteAdapter(db *sql.DB) *AuthSQLiteAdapter {
	return &AuthSQLiteAdapter{db: db}
}

func (a AuthSQLiteAdapter) RegisterUser(ctx context.Context, email string, passHash []byte) (int64, error) {
	query := "INSERT INTO users (email, pass_hash) VALUES (?, ?) RETURNING id"
	var id int64

	err := a.db.QueryRowContext(ctx, query, email, passHash).Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, errs.ErrUserNotFound
		}
		return 0, fmt.Errorf("error to insert user: %w", err)
	}

	return id, nil
}

func (a AuthSQLiteAdapter) LoginUser(ctx context.Context, email string) (models.User, error) {
	query := "SELECT id, email, pass_hash FROM users WHERE email = ?"
	var user models.User

	err := a.db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Email, &user.PassHash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.User{}, errs.ErrUserNotFound
		}
		return models.User{}, fmt.Errorf("failed to fetch user: %w", err)
	}

	return user, nil
}
