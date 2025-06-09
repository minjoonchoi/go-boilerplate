package domain

import "errors"

// Repository errors
var (
	// ErrNotFound is returned when a requested resource is not found
	ErrNotFound = errors.New("resource not found")
	// ErrDuplicate is returned when trying to create a resource that already exists
	ErrDuplicate = errors.New("resource already exists")
)

// Todo business logic errors
var (
	// ErrInvalidTodoTitle is returned when todo title is empty or invalid
	ErrInvalidTodoTitle = errors.New("todo title cannot be empty")
	// ErrTodoAlreadyCompleted is returned when trying to complete an already completed todo
	ErrTodoAlreadyCompleted = errors.New("todo is already completed")
)

// User business logic errors
var (
	// ErrInvalidUsername is returned when username is empty or invalid
	ErrInvalidUsername = errors.New("username cannot be empty")
	// ErrUsernameDuplicate is returned when trying to create a user with existing username
	ErrUsernameDuplicate = errors.New("username already exists")
	// ErrInvalidEmail is returned when email format is invalid
	ErrInvalidEmail = errors.New("invalid email format")
)
