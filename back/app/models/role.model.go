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
	database.RegisterModel(Role{})
}

type Role struct {
	ID                 uuid.UUID      `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name               string         `gorm:"type:varchar(255);not null;unique"`
	DefaultForNewUsers bool           `gorm:"type:boolean;not null;default:false"`
	CreatedAt          time.Time      `gorm:"type:timestamp;not null"`
	UpdatedAt          time.Time      `gorm:"type:timestamp;not null"`
	DeletedAt          gorm.DeletedAt `gorm:"type:timestamp;null"`
	// Relationships
	Users []User `gorm:"many2many:user_roles;"`
}

type RoleDto struct {
	ID        uuid.UUID `json:"id,omitempty" ts_type:"string" ts_transform:"__VALUE__.toString()"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"created,omitempty" ts_type:"Date" ts_transform:"new Date(__VALUE__)"`
	UpdatedAt time.Time `json:"updated,omitempty" ts_type:"Date" ts_transform:"new Date(__VALUE__)"`
}

func (role Role) ToDto() interface{} {
	return RoleDto{
		ID:        role.ID,
		Name:      role.Name,
		CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt,
	}
}

func (role Role) GetId() uuid.UUID {
	return role.ID
}

func (role Role) Validate() []*validation.ErrorResponse {
	return validation.ValidateStruct(role)
}
