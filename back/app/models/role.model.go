package models

import (
	"nartex/ngr-stack/database"

	"github.com/google/uuid"
)

func init() {
	// Register models in this file that will be synced with the database
	// Example: database.RegisterModel(User{})
	database.RegisterModel(Role{})
}

type Role struct {
	ID uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	baseEntity
	Name               string `gorm:"type:varchar(255);not null;unique"`
	DefaultForNewUsers bool   `gorm:"type:boolean;not null;default:false"`
	// Relationships
	Users []*User `gorm:"many2many:user_roles;"`
}

func (role Role) ToDto() DTO {
	r := RoleDto{}
	r.ID = role.ID
	r.Name = role.Name
	r.CreatedAt = role.CreatedAt
	r.UpdatedAt = role.UpdatedAt
	return r
}

func (role Role) GetId() uuid.UUID {
	return role.ID
}

type RoleDto struct {
	ID uuid.UUID `json:"id,omitempty"`
	baseDTO
	Name string `json:"name,omitempty"`
}

func (roleDto RoleDto) GetId() uuid.UUID {
	return roleDto.ID
}

func (roleDto RoleDto) Validate() []*ErrorResponse {
	return ValidateStruct(roleDto)
}
