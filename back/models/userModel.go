package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User is the model for the user table
type User struct {
	ID        *uuid.UUID     `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name      string         `gorm:"type:varchar(255);not null"`
	Email     string         `gorm:"type:varchar(255);not null;unique"`
	Password  string         `gorm:"type:varchar(255);not null"`
	Avatar    string         `gorm:"type:varchar(255);not null"`
	Role      string         `gorm:"type:varchar(255);not null;default:'user'"`
	Provider  string         `gorm:"type:varchar(255);not null;default:'local'"`
	CreatedAt time.Time      `gorm:"type:timestamp;not null"`
	UpdatedAt time.Time      `gorm:"type:timestamp;not null"`
	DeletedAt gorm.DeletedAt `gorm:"type:timestamp;null"`
}

// SignUpInput is the request payload for user signup
type SignUpInput struct {
	Name            string `json:"name" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=8"`
	PasswordConfirm string `json:"password_confirm" validate:"required,min=8,eqfield=Password"`
	Avatar          string `json:"avatar"`
}

// LogInInput is the request payload for user login
type LogInInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

// UserResponse is the response payload for the user model
type UserResponse struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Avatar    string    `json:"avatar,omitempty"`
	Role      string    `json:"role,omitempty"`
	Provider  string    `json:"provider,omitempty"`
	CreatedAt time.Time `json:"created,omitempty"`
	UpdatedAt time.Time `json:"updated,omitempty"`
}

// FilterUserRecord is a function to convert the user model to user response payload
func FilterUserRecord(user *User) UserResponse {
	return UserResponse{
		ID:        *user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Avatar:    user.Avatar,
		Role:      user.Role,
		Provider:  user.Provider,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

// ErrorResponse is the response payload for error response
type ErrorResponse struct {
	FailedField string `json:"failed_field"`
	Tag         string `json:"tag"`
	Value       string `json:"value,omitempty"`
}
