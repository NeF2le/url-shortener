package service

import (
	"context"
	"errors"
	"github.com/NeF2le/url-shortener/auth_service/internal/ports"
	"github.com/NeF2le/url-shortener/auth_service/internal/service/utils"
	errs "github.com/NeF2le/url-shortener/common/errors"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthService struct {
	storage           ports.StorageRepository
	jwtSecret         string
	RefreshExpiration time.Duration
	AccessExpiration  time.Duration
}

func NewAuthService(
	storage ports.StorageRepository,
	jwtSecret string,
	RefreshExpiration time.Duration,
	AccessExpiration time.Duration,
) *AuthService {
	return &AuthService{
		storage:           storage,
		jwtSecret:         jwtSecret,
		RefreshExpiration: RefreshExpiration,
		AccessExpiration:  AccessExpiration,
	}
}

func (s *AuthService) Register(ctx context.Context, email string, pass string) (int64, error) {
	passHash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	userId, err := s.storage.RegisterUser(ctx, email, passHash)
	if err != nil {
		if errors.Is(err, errs.ErrUserAlreadyExists) {
			return 0, errs.ErrUserAlreadyExists
		}
		return 0, err
	}

	return userId, nil
}

func (s *AuthService) Login(ctx context.Context, email string, pass string) (string, error) {
	user, err := s.storage.LoginUser(ctx, email)
	if err != nil {
		if errors.Is(err, errs.ErrUserNotFound) {
			return "", errs.ErrInvalidCredentials
		}
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword(user.PassHash, []byte(pass)); err != nil {
		return "", errs.ErrInvalidCredentials
	}

	token, err := utils.GenerateJWT(user.ID, s.AccessExpiration, s.jwtSecret)
	if err != nil {
		return "", err
	}

	return token, nil
}
