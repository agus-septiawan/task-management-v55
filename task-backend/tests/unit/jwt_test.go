package unit

import (
	"testing"
	"time"

	"github.com/Mahathirrr/task-management-backend/pkg/jwt"
)

func TestJWTManager(t *testing.T) {
	// Setup
	accessSecret := "test-access-secret"
	refreshSecret := "test-refresh-secret"
	accessExpire := 15 * time.Minute
	refreshExpire := 24 * time.Hour

	jwtManager := jwt.NewJWTManager(accessSecret, refreshSecret, accessExpire, refreshExpire)

	userID := 1
	email := "test@example.com"
	role := "user"

	t.Run("GenerateTokenPair", func(t *testing.T) {
		tokenPair, err := jwtManager.GenerateTokenPair(userID, email, role)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if tokenPair.AccessToken == "" {
			t.Error("Expected access token to be generated")
		}

		if tokenPair.RefreshToken == "" {
			t.Error("Expected refresh token to be generated")
		}
	})

	t.Run("ValidateAccessToken", func(t *testing.T) {
		tokenPair, err := jwtManager.GenerateTokenPair(userID, email, role)
		if err != nil {
			t.Fatalf("Failed to generate token pair: %v", err)
		}

		claims, err := jwtManager.ValidateAccessToken(tokenPair.AccessToken)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if claims.UserID != userID {
			t.Errorf("Expected user ID %d, got %d", userID, claims.UserID)
		}

		if claims.Email != email {
			t.Errorf("Expected email %s, got %s", email, claims.Email)
		}

		if claims.Role != role {
			t.Errorf("Expected role %s, got %s", role, claims.Role)
		}
	})

	t.Run("ValidateRefreshToken", func(t *testing.T) {
		tokenPair, err := jwtManager.GenerateTokenPair(userID, email, role)
		if err != nil {
			t.Fatalf("Failed to generate token pair: %v", err)
		}

		claims, err := jwtManager.ValidateRefreshToken(tokenPair.RefreshToken)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if claims.UserID != userID {
			t.Errorf("Expected user ID %d, got %d", userID, claims.UserID)
		}
	})

	t.Run("ValidateInvalidToken", func(t *testing.T) {
		_, err := jwtManager.ValidateAccessToken("invalid-token")
		if err == nil {
			t.Error("Expected error for invalid token")
		}
	})

	t.Run("ValidateExpiredToken", func(t *testing.T) {
		// Create JWT manager with very short expiry
		shortJWT := jwt.NewJWTManager(accessSecret, refreshSecret, 1*time.Nanosecond, refreshExpire)
		
		tokenPair, err := shortJWT.GenerateTokenPair(userID, email, role)
		if err != nil {
			t.Fatalf("Failed to generate token pair: %v", err)
		}

		// Wait for token to expire
		time.Sleep(2 * time.Nanosecond)

		_, err = shortJWT.ValidateAccessToken(tokenPair.AccessToken)
		if err == nil {
			t.Error("Expected error for expired token")
		}
	})
}