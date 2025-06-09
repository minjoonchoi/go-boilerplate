package persistence

import (
	"context"
	"sync"

	"go-boilerplate/internal/domain"
	"go-boilerplate/internal/domain/model"
)

// UserRepository implements the UserRepositoryPort interface
type UserRepository struct {
	users         map[int]*model.User
	usernameIndex map[string]*model.User
	mu            sync.RWMutex
	nextID        int
}

// NewUserRepository creates a new UserRepository
func NewUserRepository() *UserRepository {
	return &UserRepository{
		users:         make(map[int]*model.User),
		usernameIndex: make(map[string]*model.User),
		nextID:        1,
	}
}

// Create creates a new user
func (r *UserRepository) Create(ctx context.Context, user *model.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	user.ID = r.nextID
	r.nextID++
	r.users[user.ID] = user
	r.usernameIndex[user.Username] = user
	return nil
}

// GetByID retrieves a user by ID
func (r *UserRepository) GetByID(ctx context.Context, id int) (*model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, exists := r.users[id]
	if !exists {
		return nil, domain.ErrNotFound
	}
	return user, nil
}

// GetByUsername retrieves a user by username
func (r *UserRepository) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, exists := r.usernameIndex[username]
	if !exists {
		return nil, domain.ErrNotFound
	}
	return user, nil
}

// List retrieves all users
func (r *UserRepository) List(ctx context.Context) ([]*model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	users := make([]*model.User, 0, len(r.users))
	for _, user := range r.users {
		users = append(users, user)
	}
	return users, nil
}

// Update updates an existing user
func (r *UserRepository) Update(ctx context.Context, user *model.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[user.ID]; !exists {
		return domain.ErrNotFound
	}

	r.users[user.ID] = user
	r.usernameIndex[user.Username] = user
	return nil
}

// Delete deletes a user
func (r *UserRepository) Delete(ctx context.Context, id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	user, exists := r.users[id]
	if !exists {
		return domain.ErrNotFound
	}

	delete(r.users, id)
	delete(r.usernameIndex, user.Username)
	return nil
}
