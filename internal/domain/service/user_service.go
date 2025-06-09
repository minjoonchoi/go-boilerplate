package service

import (
	"context"
	"strings"

	"go-boilerplate/internal/domain"
	"go-boilerplate/internal/domain/model"
	"go-boilerplate/internal/domain/port"
)

// UserService implements the UserServicePort interface
type UserService struct {
	repo port.UserRepositoryPort
}

// NewUserService creates a new UserService
func NewUserService(repo port.UserRepositoryPort) *UserService {
	return &UserService{
		repo: repo,
	}
}

// CreateUser creates a new user
func (s *UserService) CreateUser(ctx context.Context, req *model.CreateUserRequest) (*model.User, error) {
	// Business logic validation
	if strings.TrimSpace(req.Username) == "" {
		return nil, domain.ErrInvalidUsername
	}

	// Check for duplicate username
	if existingUser, _ := s.repo.GetByUsername(ctx, req.Username); existingUser != nil {
		return nil, domain.ErrUsernameDuplicate
	}

	user := &model.User{
		Username: req.Username,
		Email:    req.Email,
		Name:     req.Name,
	}

	if err := s.repo.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

// GetUser retrieves a user by ID
func (s *UserService) GetUser(ctx context.Context, id int) (*model.User, error) {
	return s.repo.GetByID(ctx, id)
}

// ListUsers retrieves all users
func (s *UserService) ListUsers(ctx context.Context) ([]*model.User, error) {
	return s.repo.List(ctx)
}

// UpdateUser updates an existing user
func (s *UserService) UpdateUser(ctx context.Context, id int, req *model.UpdateUserRequest) (*model.User, error) {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Name != "" {
		user.Name = req.Name
	}

	if err := s.repo.Update(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

// DeleteUser deletes a user
func (s *UserService) DeleteUser(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
