package models

import (
	"nartex/ngr-stack/database"
	"strings"

	"github.com/google/uuid"
)

func init() {
	// Register models in this file that will be synced with the database
	// Example: database.RegisterModel(User{})
	database.RegisterModel(User{})
}

// User is the model for the user table
// Some decisions were made here:
// - Composition is used with the baseEntity struct, as it is a common pattern
// - Roles are many to many, as a user can have many roles and a role can have many users
// - Posts are one to many, as a user can have many posts and a post can have only one user
// - The avatar is stored as bytes, as it is a small image
// - The provider is used to know if the user is local or from a third party
type User struct {
	ID uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	baseEntity
	Name     string `gorm:"type:varchar(255);not null"`
	Email    string `gorm:"type:varchar(255);not null;unique"`
	Password string `gorm:"type:varchar(255);not null"`
	Avatar   string `gorm:"type:bytes;not null"`
	Provider string `gorm:"type:varchar(255);not null;default:'local'"`
	// Relationships

	Roles []*Role `gorm:"many2many:user_roles;"`
	Posts []*Post
}

// ToDto converts the user model to a user DTO
// This function is used to convert the model to a DTO before sending it to the frontend
// This funciton fulfills the Dtoable interface
func (user *User) ToDto() UserDTO {
	filteredRoles := make([]RoleDto, len(user.Roles))
	for i, role := range user.Roles {
		filteredRoles[i] = role.ToDto().(RoleDto)
	}

	u := UserDTO{}
	u.ID = user.ID
	u.Name = user.Name
	u.Email = user.Email
	u.Avatar = user.Avatar
	u.Roles = filteredRoles
	u.Provider = user.Provider
	u.CreatedAt = user.CreatedAt
	u.UpdatedAt = user.UpdatedAt
	return u
}

// GetId returns the ID of the user
// This function fulfills the Identifiable interface
func (user *User) GetId() uuid.UUID {
	return user.ID
}

// Matches returns true if the user matches the query
// This function fulfills the Searchable interface
func (user *User) Matches(query string) bool {
	return strings.Contains(user.Name, query) || strings.Contains(user.Email, query)
}

// UserDTO is the response payload for the user model
// Some decisions were made here:
// - Composition is used with the baseDTO struct, as it is a common pattern
// - The roles are converted to strings, as the frontend does not need the full role DTO
// - The provider is used to know if the user is local or from a third party
// - The avatar is stored as bytes, as it is a small image

type UserDTO struct {
	ID uuid.UUID `json:"id,omitempty" ts_type:"string" ts_transform:"__VALUE__.toString()"`
	baseDTO
	Name     string    `json:"name,omitempty"`
	Email    string    `json:"email,omitempty"`
	Avatar   string    `json:"avatar,omitempty"`
	Roles    []RoleDto `json:"roles,omitempty" ts_type:"Role[]" ts_transform:"__VALUE__.map((role: Role) => role.Name.toString())"`
	Provider string    `json:"provider,omitempty"`
}

// GetId returns the ID of the user
// This function fulfills the Identifiable interface
func (userDTO UserDTO) GetId() uuid.UUID {
	return userDTO.ID
}

// Validate validates the user DTO
// This function fulfills the Validable interface
func (userDTO UserDTO) Validate() []*ErrorResponse {
	return ValidateStruct(userDTO)
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

// Validate validates the signup input
// This function fulfills the Validable interface
func (signup SignUpInput) Validate() []*ErrorResponse {
	return ValidateStruct(signup)
}

// LogInInput is the request payload for user login
// As the model represents a form in the frontend, the json tags are used to generate the form
// The validate tags are used to validate the input
// The title tags are used to generate the form labels
// The widget tags are used to select the widget to use in the form
type LogInInput struct {
	Email    string `json:"email" validate:"required,email" title:"Email" widget:"email"`
	Password string `json:"password" validate:"required,min=8" title:"Password" widget:"password"`
}

// Validate validates the login input
// This function fulfills the Validable interface
func (login LogInInput) Validate() []*ErrorResponse {
	return ValidateStruct(login)
}
