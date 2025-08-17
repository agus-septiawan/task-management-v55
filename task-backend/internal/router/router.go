package router

import (
	"net/http"

	"github.com/Mahathirrr/task-management-backend/internal/handler"
	"github.com/Mahathirrr/task-management-backend/internal/middleware"
	"github.com/Mahathirrr/task-management-backend/pkg/jwt"
	"github.com/gorilla/mux"
)

func SetupRoutes(authHandler *handler.AuthHandler, oauthHandler *handler.OAuthHandler, taskHandler *handler.TaskHandler, adminHandler *handler.AdminHandler, jwtManager *jwt.JWTManager) http.Handler {
	r := mux.NewRouter()

	// Apply global middleware
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.LoggingMiddleware())

	// API v1 routes
	api := r.PathPrefix("/api/v1").Subrouter()

	// Public routes (tidak perlu authentication)
	auth := api.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/register", authHandler.Register).Methods("POST")
	auth.HandleFunc("/login", authHandler.Login).Methods("POST")
	auth.HandleFunc("/refresh", authHandler.RefreshToken).Methods("POST")
	auth.HandleFunc("/logout", authHandler.Logout).Methods("POST")

	// OAuth routes
	oauth := auth.PathPrefix("/oauth").Subrouter()
	oauth.HandleFunc("/google", oauthHandler.GoogleAuth).Methods("GET")
	oauth.HandleFunc("/google/callback", oauthHandler.GoogleCallback).Methods("GET")

	// Protected routes (perlu authentication)
	protected := api.PathPrefix("").Subrouter()
	protected.Use(middleware.AuthMiddleware(jwtManager))
	protected.HandleFunc("/auth/me", authHandler.Me).Methods("GET")

	// Task routes (perlu authentication)
	tasks := protected.PathPrefix("/tasks").Subrouter()
	tasks.HandleFunc("", taskHandler.GetTasks).Methods("GET")
	tasks.HandleFunc("", taskHandler.CreateTask).Methods("POST")
	tasks.HandleFunc("/{id:[0-9]+}", taskHandler.GetTaskByID).Methods("GET")
	tasks.HandleFunc("/{id:[0-9]+}", taskHandler.UpdateTask).Methods("PUT")
	tasks.HandleFunc("/{id:[0-9]+}", taskHandler.DeleteTask).Methods("DELETE")

	// Admin routes (perlu authentication + admin role)
	admin := api.PathPrefix("/admin").Subrouter()
	admin.Use(middleware.AuthMiddleware(jwtManager))
	admin.Use(middleware.AdminMiddleware())
	admin.HandleFunc("/users", adminHandler.GetAllUsers).Methods("GET")

	// Health check endpoint
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "ok", "message": "Server is running"}`))
	}).Methods("GET")

	return r
}
