package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GenerateJWT(userID int64, ttl time.Duration, jwtSecret string) (string, error) {
	claims := jwt.MapClaims{
		"exp": time.Now().Add(ttl).Unix(),
		"iat": time.Now().Unix(),
		"uid": userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}
