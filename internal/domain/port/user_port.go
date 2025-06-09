package port

import (
	"context"
	"go-boilerplate/internal/domain/model"
)

// UserRepositoryPort defines the interface for user persistence
type UserRepositoryPort interface {
	Create(ctx context.Context, user *model.User) error
	GetByID(ctx context.Context, id int) (*model.User, error)
	GetByUsername(ctx context.Context, username string) (*model.User, error)
	List(ctx context.Context) ([]*model.User, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id int) error
}

// UserServicePort defines the interface for user business logic
type UserServicePort interface {
	CreateUser(ctx context.Context, req *model.CreateUserRequest) (*model.User, error)
	GetUser(ctx context.Context, id int) (*model.User, error)
	ListUsers(ctx context.Context) ([]*model.User, error)
	UpdateUser(ctx context.Context, id int, req *model.UpdateUserRequest) (*model.User, error)
	DeleteUser(ctx context.Context, id int) error
}
