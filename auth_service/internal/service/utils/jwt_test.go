package utils

import (
	"strings"
	"testing"
	"time"
)

func TestGenerateJWT(t *testing.T) {
	tests := []struct {
		name   string
		userID string
		ttl    time.Duration
	}{
		{
			name:   "access token",
			userID: "user123",
			ttl:    time.Hour,
		},
	}

	secret := "test-secret"
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenStr, err := GenerateJWT(tt.userID, tt.ttl, secret)
			if err != nil {
				t.Fatalf("GenerateJWT() error = %v", err)
			}
			if tokenStr == "" {
				t.Fatalf("GenerateJWT() returned empty token")
			}
			parts := strings.Split(tokenStr, ".")
			if len(parts) != 3 {
				t.Fatalf("GenerateJWT() returned invalid JWT format: %s", tokenStr)
			}
		})
	}
}
