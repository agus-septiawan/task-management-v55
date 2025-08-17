package middleware

import (
	"net/http"

	"github.com/gorilla/handlers"
)

// CORSMiddleware sets up the CORS handling using gorilla/handlers.
func CORSMiddleware() func(http.Handler) http.Handler {
	// Define allowed origins, methods, and headers
	// Anda bisa mengganti "*" dengan domain frontend Anda (misal: "http://localhost:5173") untuk keamanan lebih
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	allowedHeaders := handlers.AllowedHeaders([]string{"Content-Type", "Authorization", "X-Requested-With"})
	allowCredentials := handlers.AllowCredentials()
	maxAge := handlers.MaxAge(86400)

	// Return the CORS middleware handler
	return handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders, allowCredentials, maxAge)
}