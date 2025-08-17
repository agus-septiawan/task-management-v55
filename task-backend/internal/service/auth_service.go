package service

import (
	"errors"
	"fmt"

	"github.com/Mahathirrr/task-management-backend/internal/model"
	"github.com/Mahathirrr/task-management-backend/internal/repository"
	"github.com/Mahathirrr/task-management-backend/pkg/jwt"
	"github.com/Mahathirrr/task-management-backend/pkg/oauth"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(req *model.UserRegisterRequest) (*model.AuthResponse, error)
	Login(req *model.UserLoginRequest) (*model.AuthResponse, error)
	OAuthLogin(oauthUser *oauth.OAuthUser) (*model.AuthResponse, error)
	RefreshToken(refreshToken string) (*model.TokenResponse, error)
	GetUserProfile(userID int) (*model.User, error)
}

type authService struct {
	userRepo   repository.UserRepository
	jwtManager *jwt.JWTManager
}

func NewAuthService(userRepo repository.UserRepository, jwtManager *jwt.JWTManager) AuthService {
	return &authService{
		userRepo:   userRepo,
		jwtManager: jwtManager,
	}
}

// Register mendaftarkan user baru
func (s *authService) Register(req *model.UserRegisterRequest) (*model.AuthResponse, error) {
	// Cek apakah email sudah ada
	existingUser, err := s.userRepo.GetByEmail(req.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to check existing user: %w", err)
	}
	if existingUser != nil {
		return nil, errors.New(model.ErrEmailAlreadyExists)
	}

	// Buat user baru
	user := &model.User{
		Email:    req.Email,
		Name:     req.Name,
		Password: req.Password,
		Role:     model.UserRoleUser,
	}

	err = s.userRepo.Create(user)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	// Generate token pair
	tokenPair, err := s.jwtManager.GenerateTokenPair(user.ID, user.Email, string(user.Role))
	if err != nil {
		return nil, fmt.Errorf("failed to generate tokens: %w", err)
	}

	// Hapus password dari response
	user.Password = ""

	return &model.AuthResponse{
		AccessToken: tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
		User:        *user,
	}, nil
}

// Login melakukan autentikasi user
func (s *authService) Login(req *model.UserLoginRequest) (*model.AuthResponse, error) {
	// Cari user berdasarkan email
	user, err := s.userRepo.GetByEmail(req.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	if user == nil {
		return nil, errors.New(model.ErrInvalidCredentials)
	}

	// Verifikasi password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New(model.ErrInvalidCredentials)
	}

	// Generate token pair
	tokenPair, err := s.jwtManager.GenerateTokenPair(user.ID, user.Email, string(user.Role))
	if err != nil {
		return nil, fmt.Errorf("failed to generate tokens: %w", err)
	}

	// Hapus password dari response
	user.Password = ""

	return &model.AuthResponse{
		AccessToken: tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
		User:        *user,
	}, nil
}

// RefreshToken menggenerate access token baru dari refresh token
func (s *authService) RefreshToken(refreshToken string) (*model.TokenResponse, error) {
	// Validasi refresh token
	claims, err := s.jwtManager.ValidateRefreshToken(refreshToken)
	if err != nil {
		return nil, errors.New("invalid refresh token")
	}

	// Generate access token baru
	tokenPair, err := s.jwtManager.GenerateTokenPair(claims.UserID, claims.Email, claims.Role)
	if err != nil {
		return nil, fmt.Errorf("failed to generate new access token: %w", err)
	}

	return &model.TokenResponse{
		AccessToken: tokenPair.AccessToken,
	}, nil
}

// GetUserProfile mengambil profile user
func (s *authService) GetUserProfile(userID int) (*model.User, error) {
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user profile: %w", err)
	}
	if user == nil {
		return nil, errors.New(model.ErrUserNotFound)
	}

	// Hapus password dari response
	user.Password = ""

	return user, nil
}

// OAuthLogin handles OAuth login/registration
func (s *authService) OAuthLogin(oauthUser *oauth.OAuthUser) (*model.AuthResponse, error) {
	// Check if user exists with OAuth provider
	user, err := s.userRepo.GetByOAuth(oauthUser.Provider, oauthUser.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to check OAuth user: %w", err)
	}

	// If user doesn't exist, check by email
	if user == nil {
		user, err = s.userRepo.GetByEmail(oauthUser.Email)
		if err != nil {
			return nil, fmt.Errorf("failed to check user by email: %w", err)
		}

		if user != nil {
			// Link OAuth account to existing user
			user.OauthProvider = &oauthUser.Provider
			user.OauthID = &oauthUser.ID
			err = s.userRepo.Update(user)
			if err != nil {
				return nil, fmt.Errorf("failed to link OAuth account: %w", err)
			}
		} else {
			// Create new user from OAuth
			user = &model.User{
				Email:         oauthUser.Email,
				Name:          oauthUser.Name,
				Role:          model.UserRoleUser,
				OauthProvider: &oauthUser.Provider,
				OauthID:       &oauthUser.ID,
			}

			err = s.userRepo.Create(user)
			if err != nil {
				return nil, fmt.Errorf("failed to create OAuth user: %w", err)
			}
		}
	}

	// Generate token pair
	tokenPair, err := s.jwtManager.GenerateTokenPair(user.ID, user.Email, string(user.Role))
	if err != nil {
		return nil, fmt.Errorf("failed to generate tokens: %w", err)
	}

	// Clear password from response
	user.Password = ""

	return &model.AuthResponse{
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
		User:         *user,
	}, nil
}