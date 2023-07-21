package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Entity is an interface for all entities.
// It embeds the Identifiable and Dtoable interfaces
// Entitys must be identifiable (with a UUID) and convertable to DTOs
type Entity interface {
	Identifiable
	Dtoable
}

// DTO is an interface for all DTOs.
// It embeds the Identifiable and Validable interfaces
// DTOs must be identifiable (with a UUID) and validable (using json tags)
type DTO interface {
	Identifiable
	Validable
}

// Structs that implement this interface can be converted to DTOs
type Dtoable interface {
	ToDto() DTO
}

// Structs that implement this interface can be identified by an UUID
type Identifiable interface {
	GetId() uuid.UUID
}

// Structs that implement this interface can be validated using json tags
type Validable interface {
	Validate() []*ErrorResponse
}

// Structs that implement this interface can be searched using a query string
type Searchable interface {
	Matches(query string) bool
}

// ErrorResponse is the response payload for error response
// This struct is used to return validation errors
type ErrorResponse struct {
	FailedField string `json:"failed_field"`
	Tag         string `json:"tag"`
	Value       string `json:"value,omitempty"`
}

// baseEntity is a embedded struct for the common fields of all Entities
// Some decisions were made here:
// - IDs are UUIDs and not sequential
// - DeletedAt is a timestamp and not a boolean, so that we can query for deleted entities
// - CreatedAt and UpdatedAt are timestamps and not dates, as they are more precise
// - CreatedAt and UpdatedAt are not nullable, as they are always set by the database
// - DeletedAt is nullable, as it is set by the database when the entity is deleted

type baseEntity struct {
	// Seems like there is a bug with gorm and embedded structs with UUIDs.
	//ID        uuid.UUID      `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	CreatedAt time.Time      `gorm:"type:timestamp;not null"`
	UpdatedAt time.Time      `gorm:"type:timestamp;not null"`
	DeletedAt gorm.DeletedAt `gorm:"type:timestamp;null"`
}

// baseDTO is a embedded struct for the common fields of all DTOs
// purposedly it does not include the DeletedAt field, as deleted entities should not be returned
// IDs are exposed as they are UUIDs and are not sequential
// As we expect to be able to convert our DTOs to Typescript types, we need to specify the types and transformations
type baseDTO struct {
	// Embedded structs are not supported by the ts_transform tag
	//ID        uuid.UUID `json:"id,omitempty" ts_type:"string" ts_transform:"__VALUE__.toString()"`
	CreatedAt time.Time `json:"created,omitempty" ts_type:"Date" ts_transform:"new Date(__VALUE__)"`
	UpdatedAt time.Time `json:"updated,omitempty" ts_type:"Date" ts_transform:"new Date(__VALUE__)"`
}
