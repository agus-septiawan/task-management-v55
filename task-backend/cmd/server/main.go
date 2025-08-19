package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Mahathirrr/task-management-backend/internal/config"
	"github.com/Mahathirrr/task-management-backend/internal/database"
	"github.com/Mahathirrr/task-management-backend/internal/handler"
	"github.com/Mahathirrr/task-management-backend/internal/repository"
	"github.com/Mahathirrr/task-management-backend/internal/router"
	"github.com/Mahathirrr/task-management-backend/internal/service"
	"github.com/Mahathirrr/task-management-backend/pkg/jwt"
	"github.com/Mahathirrr/task-management-backend/pkg/oauth"
)

func main() {
	// Load configuration

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize database
	if err := database.InitDatabase(&cfg.Database); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Initialize JWT manager
	jwtManager := jwt.NewJWTManager(
		cfg.JWT.AccessSecret,
		cfg.JWT.RefreshSecret,
		cfg.JWT.AccessExpire,
		cfg.JWT.RefreshExpire,
	)

	// Initialize OAuth manager
	oauthManager := oauth.InitializeOAuth(
		cfg.OAuth.Google.ClientID,
		cfg.OAuth.Google.ClientSecret,
		cfg.OAuth.Google.RedirectURL,
	)

	// Initialize repositories
	userRepo := repository.NewUserRepository(database.GetDB())
	taskRepo := repository.NewTaskRepository(database.GetDB())

	// Initialize services
	authService := service.NewAuthService(userRepo, jwtManager)
	userService := service.NewUserService(userRepo)
	taskService := service.NewTaskService(taskRepo)

	// Initialize handlers
	authHandler := handler.NewAuthHandler(authService)
	oauthHandler := handler.NewOAuthHandler(authService, oauthManager)
	taskHandler := handler.NewTaskHandler(taskService)
	adminHandler := handler.NewAdminHandler(userService)

	// Setup routes
	routerHandler := router.SetupRoutes(authHandler, oauthHandler, taskHandler, adminHandler, jwtManager, &cfg.CORS)

	// Start server
	addr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	log.Printf("Server starting on: %s", addr)
	if err := http.ListenAndServe(addr, routerHandler); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
