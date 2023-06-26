package models

import (
	"nartex/ngr-stack/database"
	"nartex/ngr-stack/utils/validation"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {
	// Register models in this file that will be synced with the database
	// Example: database.RegisterModel(User{})
	database.RegisterModel(User{})
}

// User is the model for the user table
type User struct {
	ID        uuid.UUID      `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name      string         `gorm:"type:varchar(255);not null"`
	Email     string         `gorm:"type:varchar(255);not null;unique"`
	Password  string         `gorm:"type:varchar(255);not null"`
	Avatar    string         `gorm:"type:bytes;not null"`
	Provider  string         `gorm:"type:varchar(255);not null;default:'local'"`
	CreatedAt time.Time      `gorm:"type:timestamp;not null"`
	UpdatedAt time.Time      `gorm:"type:timestamp;not null"`
	DeletedAt gorm.DeletedAt `gorm:"type:timestamp;null"`
	// Relationships
	Roles []Role `gorm:"many2many:user_roles;"`
}

func (user *User) Validate() []*validation.ErrorResponse {
	return validation.ValidateStruct(user)
}

func (user *User) ToDto() UserDTO {
	filteredRoles := make([]RoleDto, len(user.Roles))
	for i, role := range user.Roles {
		filteredRoles[i] = role.ToDto().(RoleDto)
	}
	return UserDTO{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Avatar:    user.Avatar,
		Roles:     filteredRoles,
		Provider:  user.Provider,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func (user *User) GetId() uuid.UUID {
	return user.ID
}

// SignUpInput is the request payload for user signup
// As the model represents a form in the frontend, the json tags are used to generate the form
// The validate tags are used to validate the input
// The title tags are used to generate the form labels
// The widget tags are used to select the widget to use in the form
// The options tags are used to pass options to the widget
type SignUpInput struct {
	Name            string `json:"name" validate:"required" title:"Name"`
	Email           string `json:"email" validate:"required,email" title:"Email" widget:"email"`
	Password        string `json:"password" validate:"required,min=8" title:"Password" widget:"password"`
	PasswordConfirm string `json:"password_confirm" validate:"required,min=8,eqfield=Password" title:"Password Confirm" widget:"password"`
	Avatar          string `json:"avatar" title:"Avatar" widget:"file" options:"accept:image/*,filePreview:false"`
}

// LogInInput is the request payload for user login
type LogInInput struct {
	Email    string `json:"email" validate:"required,email" title:"Email" widget:"email"`
	Password string `json:"password" validate:"required,min=8" title:"Password" widget:"password"`
}

// UserDTO is the response payload for the user model
type UserDTO struct {
	ID        uuid.UUID `json:"id,omitempty" ts_type:"string" ts_transform:"__VALUE__.toString()"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Avatar    string    `json:"avatar,omitempty"`
	Roles     []RoleDto `json:"roles,omitempty" ts_type:"Role[]" ts_transform:"__VALUE__.map((role: Role) => role.Name.toString())"`
	Provider  string    `json:"provider,omitempty"`
	CreatedAt time.Time `json:"created,omitempty" ts_type:"Date" ts_transform:"new Date(__VALUE__)"`
	UpdatedAt time.Time `json:"updated,omitempty" ts_type:"Date" ts_transform:"new Date(__VALUE__)"`
}

// FilterUserRecord is a function to convert the user model to user response payload
func FilterUserRecord(user *User) UserDTO {

	filteredRoles := make([]RoleDto, len(user.Roles))
	for i, role := range user.Roles {
		filteredRoles[i] = role.ToDto().(RoleDto)
	}
	return UserDTO{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Avatar:    user.Avatar,
		Roles:     filteredRoles,
		Provider:  user.Provider,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
