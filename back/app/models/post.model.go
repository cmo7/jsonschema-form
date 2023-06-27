package models

import (
	"nartex/ngr-stack/database"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {
	// Register models in this file that will be synced with the database
	// Example: database.RegisterModel(User{})
	database.RegisterModel(Post{})
}

type Post struct {
	ID        uuid.UUID      `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Title     string         `gorm:"type:varchar(255);not null"`
	Content   string         `gorm:"type:string;not null"`
	CreatedAt time.Time      `gorm:"type:timestamp;not null"`
	UpdatedAt time.Time      `gorm:"type:timestamp;not null"`
	DeletedAt gorm.DeletedAt `gorm:"type:timestamp;null"`
	// Relationships
	User   User `gorm:"foreignKey:UserID"`
	UserID uuid.UUID
}

func (post Post) Validate() []*ErrorResponse {
	return ValidateStruct(post)
}

func (post Post) ToDto() interface{} {
	return PostDTO{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}
}

func (post Post) GetId() uuid.UUID {
	return post.ID
}

func (post *Post) Matches(query string) bool {
	// TODO: Implement
	return false
}

type PostDTO struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
