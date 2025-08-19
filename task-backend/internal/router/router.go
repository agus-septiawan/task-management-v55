package router

import (
	"net/http"

	"github.com/Mahathirrr/task-management-backend/internal/config"
	"github.com/Mahathirrr/task-management-backend/internal/handler"
	"github.com/Mahathirrr/task-management-backend/internal/middleware"
	"github.com/Mahathirrr/task-management-backend/pkg/jwt"
	"github.com/gorilla/mux"
)

func SetupRoutes(authHandler *handler.AuthHandler, oauthHandler *handler.OAuthHandler, taskHandler *handler.TaskHandler, adminHandler *handler.AdminHandler, jwtManager *jwt.JWTManager, corsConfig *config.CORSConfig) http.Handler {
	r := mux.NewRouter()

	// Apply global middleware - CORS must be first
	r.Use(middleware.CORSMiddleware(corsConfig))
	r.Use(middleware.LoggingMiddleware())

	// Health check endpoint (before API routes)
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "ok", "message": "Server is running"}`))
	}).Methods("GET", "OPTIONS")

	// API v1 routes
	api := r.PathPrefix("/api/v1").Subrouter()

	// Public routes (tidak perlu authentication)
	auth := api.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/register", authHandler.Register).Methods("POST", "OPTIONS")
	auth.HandleFunc("/login", authHandler.Login).Methods("POST", "OPTIONS")
	auth.HandleFunc("/refresh", authHandler.RefreshToken).Methods("POST", "OPTIONS")
	auth.HandleFunc("/logout", authHandler.Logout).Methods("POST", "OPTIONS")

	// OAuth routes
	oauth := auth.PathPrefix("/oauth").Subrouter()
	oauth.HandleFunc("/google", oauthHandler.GoogleAuth).Methods("GET", "OPTIONS")
	oauth.HandleFunc("/google/callback", oauthHandler.GoogleCallback).Methods("GET", "OPTIONS")

	// Protected routes (perlu authentication)
	protected := api.PathPrefix("").Subrouter()
	protected.Use(middleware.AuthMiddleware(jwtManager))
	protected.HandleFunc("/auth/me", authHandler.Me).Methods("GET", "OPTIONS")

	// Task routes (perlu authentication)
	tasks := protected.PathPrefix("/tasks").Subrouter()
	tasks.HandleFunc("", taskHandler.GetTasks).Methods("GET", "OPTIONS")
	tasks.HandleFunc("", taskHandler.CreateTask).Methods("POST", "OPTIONS")
	tasks.HandleFunc("/{id:[0-9]+}", taskHandler.GetTaskByID).Methods("GET", "OPTIONS")
	tasks.HandleFunc("/{id:[0-9]+}", taskHandler.UpdateTask).Methods("PUT", "OPTIONS")
	tasks.HandleFunc("/{id:[0-9]+}", taskHandler.DeleteTask).Methods("DELETE", "OPTIONS")

	// Admin routes (perlu authentication + admin role)
	admin := api.PathPrefix("/admin").Subrouter()
	admin.Use(middleware.AuthMiddleware(jwtManager))
	admin.Use(middleware.AdminMiddleware())
	admin.HandleFunc("/users", adminHandler.GetAllUsers).Methods("GET", "OPTIONS")

	return r
}
