package middleware

import (
	"net/http"

	"github.com/Mahathirrr/task-management-backend/internal/config"
	"github.com/gorilla/handlers"
)

// CORSMiddleware sets up the CORS handling using gorilla/handlers with configurable options.
func CORSMiddleware(corsConfig *config.CORSConfig) func(http.Handler) http.Handler {
	// Use configuration values for CORS settings
	allowedOrigins := handlers.AllowedOrigins(corsConfig.AllowedOrigins)
	allowedMethods := handlers.AllowedMethods(corsConfig.AllowedMethods)
	allowedHeaders := handlers.AllowedHeaders(corsConfig.AllowedHeaders)
	maxAge := handlers.MaxAge(corsConfig.MaxAge)

	// Conditionally add credentials support
	if corsConfig.AllowCredentials {
		allowCredentials := handlers.AllowCredentials()
		return handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders, allowCredentials, maxAge)
	}

	return handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders, maxAge)
}
