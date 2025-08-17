package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/Mahathirrr/task-management-backend/internal/model"
	"github.com/Mahathirrr/task-management-backend/pkg/jwt"
	"github.com/Mahathirrr/task-management-backend/pkg/response"
)

// contextKey menggunakan struct untuk menghindari collision
type contextKey struct{}

var UserContextKey = contextKey{}

// AuthMiddleware memvalidasi JWT token
func AuthMiddleware(jwtManager *jwt.JWTManager) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Ambil token dari header Authorization
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				response.Error(w, http.StatusUnauthorized, model.ErrUnauthorized)
				return
			}

			// Cek format Bearer token
			tokenParts := strings.Split(authHeader, " ")
			if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
				response.Error(w, http.StatusUnauthorized, "Invalid token format")
				return
			}

			// Validasi token
			claims, err := jwtManager.ValidateAccessToken(tokenParts[1])
			if err != nil {
				response.Error(w, http.StatusUnauthorized, "Invalid token")
				return
			}

			// Simpan user info ke context
			ctx := context.WithValue(r.Context(), UserContextKey, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// AdminMiddleware memvalidasi role admin
func AdminMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Ambil user dari context
			claims, ok := r.Context().Value(UserContextKey).(*jwt.Claims)
			if !ok {
				response.Error(w, http.StatusUnauthorized, model.ErrUnauthorized)
				return
			}

			// Cek role admin
			if claims.Role != string(model.UserRoleAdmin) {
				response.Error(w, http.StatusForbidden, model.ErrForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// GetUserFromContext mengambil user claims dari context
func GetUserFromContext(r *http.Request) (*jwt.Claims, bool) {
	claims, ok := r.Context().Value(UserContextKey).(*jwt.Claims)
	return claims, ok
}