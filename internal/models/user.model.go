package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"type:varchar(100);not null" json:"name"`
	Email     string         `gorm:"uniqueIndex;type:varchar(100);not null" json:"email"`
	Password  string         `gorm:"type:varchar(255);not null" json:"-"` // never expose password
	CreatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP;autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // soft delete
}

// CreateUserRequest represents input payload when creating a new user
// swagger:model
type CreateUserRequest struct {
	// User's full name
	// required: true
	// example: John Doe
	Name string `json:"name" binding:"required"`

	// User's email (must be unique)
	// required: true
	// example: johndoe@example.com
	Email string `json:"email" binding:"required,email"`

	// User's password (will be hashed before saving)
	// required: true
	// example: password123
	Password string `json:"password" binding:"required"`
}

// UpdateUserRequest represents input payload when updating an existing user
// swagger:model
type UpdateUserRequest struct {
	// User's full name
	// example: John Doe
	Name string `json:"name,omitempty"`

	// User's email
	// example: johndoe@example.com
	Email string `json:"email,omitempty"`

	// User's password
	// example: newpassword123
	Password string `json:"password,omitempty"`
}

// UserResponse represents the user data returned in API responses
// swagger:model
type UserResponse struct {
	// User ID
	// example: 1
	ID uint `json:"id"`

	// User's full name
	// example: John Doe
	Name string `json:"name"`

	// User's email
	// example: johndoe@example.com
	Email string `json:"email"`

	// Timestamp when user was created
	// example: 2025-09-06T12:34:56Z
	CreatedAt string `json:"created_at"`

	// Timestamp when user was last updated
	// example: 2025-09-06T12:34:56Z
	UpdatedAt string `json:"updated_at"`
}
