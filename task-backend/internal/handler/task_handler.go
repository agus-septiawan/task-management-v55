package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/Mahathirrr/task-management-backend/internal/middleware"
	"github.com/Mahathirrr/task-management-backend/internal/model"
	"github.com/Mahathirrr/task-management-backend/internal/service"
	"github.com/Mahathirrr/task-management-backend/pkg/response"
	"github.com/Mahathirrr/task-management-backend/pkg/validator"
)

type TaskHandler struct {
	taskService service.TaskService
}

func NewTaskHandler(taskService service.TaskService) *TaskHandler {
	return &TaskHandler{
		taskService: taskService,
	}
}

// CreateTask menangani pembuatan task baru
func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	// Get user from context
	claims, ok := middleware.GetUserFromContext(r)
	if !ok {
		response.Error(w, http.StatusUnauthorized, model.ErrUnauthorized)
		return
	}

	var req model.TaskCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	// Validate input
	if validationErrors := validator.ValidateStruct(req); len(validationErrors) > 0 {
		response.ValidationError(w, validationErrors)
		return
	}

	// Create task
	task, err := h.taskService.CreateTask(claims.UserID, &req)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, model.ErrInternalServer)
		return
	}

	response.Created(w, task)
}

// GetTasks menangani pengambilan tasks dengan pagination dan filter
func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	// Get user from context
	claims, ok := middleware.GetUserFromContext(r)
	if !ok {
		response.Error(w, http.StatusUnauthorized, model.ErrUnauthorized)
		return
	}

	// Parse query parameters
	page, limit := parsePageAndLimit(r)
	status := r.URL.Query().Get("status")
	search := r.URL.Query().Get("search")

	var tasksResp *model.TasksResponse
	var err error

	// Admin can see all tasks, users only their own
	if claims.Role == string(model.UserRoleAdmin) {
		tasksResp, err = h.taskService.GetAllTasks(page, limit, status, search)
	} else {
		tasksResp, err = h.taskService.GetUserTasks(claims.UserID, page, limit, status, search)
	}

	if err != nil {
		response.Error(w, http.StatusInternalServerError, model.ErrInternalServer)
		return
	}

	response.JSON(w, http.StatusOK, tasksResp)
}

// GetTaskByID menangani pengambilan task berdasarkan ID
func (h *TaskHandler) GetTaskByID(w http.ResponseWriter, r *http.Request) {
	// Get user from context
	claims, ok := middleware.GetUserFromContext(r)
	if !ok {
		response.Error(w, http.StatusUnauthorized, model.ErrUnauthorized)
		return
	}

	// Get task ID from URL
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid task ID")
		return
	}

	// Get task
	isAdmin := claims.Role == string(model.UserRoleAdmin)
	task, err := h.taskService.GetTaskByID(taskID, claims.UserID, isAdmin)
	if err != nil {
		switch err.Error() {
		case model.ErrTaskNotFound:
			response.Error(w, http.StatusNotFound, err.Error())
		case model.ErrForbidden:
			response.Error(w, http.StatusForbidden, err.Error())
		default:
			response.Error(w, http.StatusInternalServerError, model.ErrInternalServer)
		}
		return
	}

	response.JSON(w, http.StatusOK, task)
}

// UpdateTask menangani update task
func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	// Get user from context
	claims, ok := middleware.GetUserFromContext(r)
	if !ok {
		response.Error(w, http.StatusUnauthorized, model.ErrUnauthorized)
		return
	}

	// Get task ID from URL
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid task ID")
		return
	}

	var req model.TaskUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	// Validate input
	if validationErrors := validator.ValidateStruct(req); len(validationErrors) > 0 {
		response.ValidationError(w, validationErrors)
		return
	}

	// Update task
	isAdmin := claims.Role == string(model.UserRoleAdmin)
	task, err := h.taskService.UpdateTask(taskID, claims.UserID, &req, isAdmin)
	if err != nil {
		switch err.Error() {
		case model.ErrTaskNotFound:
			response.Error(w, http.StatusNotFound, err.Error())
		case model.ErrForbidden:
			response.Error(w, http.StatusForbidden, err.Error())
		default:
			response.Error(w, http.StatusInternalServerError, model.ErrInternalServer)
		}
		return
	}

	response.JSON(w, http.StatusOK, task)
}

// DeleteTask menangani penghapusan task
func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	// Get user from context
	claims, ok := middleware.GetUserFromContext(r)
	if !ok {
		response.Error(w, http.StatusUnauthorized, model.ErrUnauthorized)
		return
	}

	// Get task ID from URL
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid task ID")
		return
	}

	// Delete task
	isAdmin := claims.Role == string(model.UserRoleAdmin)
	err = h.taskService.DeleteTask(taskID, claims.UserID, isAdmin)
	if err != nil {
		switch err.Error() {
		case model.ErrTaskNotFound:
			response.Error(w, http.StatusNotFound, err.Error())
		case model.ErrForbidden:
			response.Error(w, http.StatusForbidden, err.Error())
		default:
			response.Error(w, http.StatusInternalServerError, model.ErrInternalServer)
		}
		return
	}

	response.Success(w, model.MsgTaskDeleted)
}

// parsePageAndLimit parses page and limit query parameters
func parsePageAndLimit(r *http.Request) (int, int) {
	page := 1
	limit := 10

	if pageStr := r.URL.Query().Get("page"); pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}

	if limitStr := r.URL.Query().Get("limit"); limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 && l <= 100 {
			limit = l
		}
	}

	return page, limit
}