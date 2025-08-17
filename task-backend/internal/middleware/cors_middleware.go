package middleware

import (
	"net/http"

	"github.com/gorilla/handlers"
)

// CORSMiddleware sets up the CORS handling using gorilla/handlers.
func CORSMiddleware() func(http.Handler) http.Handler {
	// Define allowed origins, methods, and headers
	// Allow specific origins for better security and cookie handling
	allowedOrigins := handlers.AllowedOrigins([]string{"http://localhost:3000", "http://localhost:5173", "http://localhost:8080"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	allowedHeaders := handlers.AllowedHeaders([]string{"Content-Type", "Authorization", "X-Requested-With"})
	allowCredentials := handlers.AllowCredentials()
	maxAge := handlers.MaxAge(86400)

	// Return the CORS middleware handler
	return handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders, allowCredentials, maxAge)
}
