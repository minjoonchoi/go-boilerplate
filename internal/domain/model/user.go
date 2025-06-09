package model

// User represents a user in the domain
type User struct {
	ID       int    `json:"id" example:"1"`
	Username string `json:"username" example:"johndoe"`
	Email    string `json:"email" example:"john@example.com"`
	Name     string `json:"name" example:"John Doe"`
}

// CreateUserRequest represents the request to create a new user
type CreateUserRequest struct {
	Username string `json:"username" binding:"required" example:"johndoe"`
	Email    string `json:"email" binding:"required" example:"john@example.com"`
	Name     string `json:"name" binding:"required" example:"John Doe"`
}

// UpdateUserRequest represents the request to update an existing user
type UpdateUserRequest struct {
	Email string `json:"email" example:"john.new@example.com"`
	Name  string `json:"name" example:"John Smith"`
}
