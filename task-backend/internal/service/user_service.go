package service

import (
	"fmt"

	"github.com/Mahathirrr/task-management-backend/internal/model"
	"github.com/Mahathirrr/task-management-backend/internal/repository"
)

type UserService interface {
	GetAllUsers(page, limit int) (*model.UsersResponse, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

// GetAllUsers mengambil semua user dengan pagination (admin only)
func (s *userService) GetAllUsers(page, limit int) (*model.UsersResponse, error) {
	users, total, err := s.userRepo.GetAll(page, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}

	// Hapus password dari semua users
	for i := range users {
		users[i].Password = ""
	}

	return &model.UsersResponse{
		Users: users,
		Total: total,
		Page:  page,
		Limit: limit,
	}, nil
}