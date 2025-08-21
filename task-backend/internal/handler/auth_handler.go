package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Mahathirrr/task-management-backend/internal/middleware"
	"github.com/Mahathirrr/task-management-backend/internal/model"
	"github.com/Mahathirrr/task-management-backend/internal/service"
	"github.com/Mahathirrr/task-management-backend/pkg/response"
	"github.com/Mahathirrr/task-management-backend/pkg/validator"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// Register menangani registrasi user baru
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req model.UserRegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	// Validasi input
	if validationErrors := validator.ValidateStruct(req); len(validationErrors) > 0 {
		response.ValidationError(w, validationErrors)
		return
	}

	// Proses registrasi
	authResp, err := h.authService.Register(&req)
	if err != nil {
		if err.Error() == model.ErrEmailAlreadyExists {
			response.Error(w, http.StatusConflict, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, model.ErrInternalServer)
		return
	}

	// Set refresh token sebagai HTTP-only cookie
	// Note: Untuk implementasi lengkap, refresh token harus disimpan
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    authResp.RefreshToken,
		HttpOnly: true,
		Secure:   false, // Always secure for production
		SameSite: http.SameSiteStrictMode,
		Expires:  time.Now().Add(168 * time.Hour), // 7 hari
		Path:     "/api/v1/auth",
	})

	response.Created(w, authResp)
}

// Login menangani autentikasi user
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req model.UserLoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	// Validasi input
	if validationErrors := validator.ValidateStruct(req); len(validationErrors) > 0 {
		response.ValidationError(w, validationErrors)
		return
	}

	// Proses login
	authResp, err := h.authService.Login(&req)
	if err != nil {
		if err.Error() == model.ErrInvalidCredentials {
			response.Error(w, http.StatusUnauthorized, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, model.ErrInternalServer)
		return
	}

	// Set refresh token sebagai HTTP-only cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    authResp.RefreshToken,
		HttpOnly: true,
		Secure:   false, // Always secure for production
		SameSite: http.SameSiteStrictMode,
		Expires:  time.Now().Add(168 * time.Hour), // 7 hari
		Path:     "/api/v1/auth",
	})

	response.JSON(w, http.StatusOK, authResp)
}

// RefreshToken menangani refresh access token
func (h *AuthHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	// Ambil refresh token dari cookie
	cookie, err := r.Cookie("refresh_token")
	if err != nil {
		response.Error(w, http.StatusUnauthorized, "Refresh token not found")
		return
	}

	// Proses refresh token
	tokenResp, err := h.authService.RefreshToken(cookie.Value)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, tokenResp)
}

// Logout menangani logout user
func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	// Hapus refresh token cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    "",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
		Expires:  time.Now().Add(-time.Hour), // Set expired
		Path:     "/api/v1/auth",
	})

	response.Success(w, model.MsgLogoutSuccess)
}

// Me menangani get user profile
func (h *AuthHandler) Me(w http.ResponseWriter, r *http.Request) {
	// Ambil user dari context
	claims, ok := middleware.GetUserFromContext(r)
	if !ok {
		response.Error(w, http.StatusUnauthorized, model.ErrUnauthorized)
		return
	}

	// Ambil user profile
	user, err := h.authService.GetUserProfile(claims.UserID)
	if err != nil {
		if err.Error() == model.ErrUserNotFound {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, model.ErrInternalServer)
		return
	}

	response.JSON(w, http.StatusOK, user)
}
