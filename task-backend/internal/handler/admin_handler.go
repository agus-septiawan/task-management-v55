package handler

import (
	"net/http"
	"strconv"

	"github.com/Mahathirrr/task-management-backend/internal/model"
	"github.com/Mahathirrr/task-management-backend/internal/service"
	"github.com/Mahathirrr/task-management-backend/pkg/response"
)

type AdminHandler struct {
	userService service.UserService
}

func NewAdminHandler(userService service.UserService) *AdminHandler {
	return &AdminHandler{
		userService: userService,
	}
}

// GetAllUsers menangani get semua users (admin only)
func (h *AdminHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	page := 1
	limit := 10

	if pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}

	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 && l <= 100 {
			limit = l
		}
	}

	// Ambil semua users
	usersResp, err := h.userService.GetAllUsers(page, limit)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, model.ErrInternalServer)
		return
	}

	response.JSON(w, http.StatusOK, usersResp)
}