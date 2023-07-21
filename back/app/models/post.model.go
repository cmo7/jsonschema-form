package models

import (
	"nartex/ngr-stack/database"

	"github.com/google/uuid"
)

func init() {
	// Register models in this file that will be synced with the database
	// Example: database.RegisterModel(User{})
	database.RegisterModel(Post{})
}

type Post struct {
	ID uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	baseEntity
	Title   string `gorm:"type:varchar(255);not null"`
	Content string `gorm:"type:string;not null"`
	// Relationships

	// Belongs to User
	// Both fields are required, UserID is the foreign key and User is the relationship
	UserID uuid.UUID `gorm:"type:uuid;not null"`
	User   User      `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (post Post) ToDto() DTO {
	p := PostDTO{}
	p.ID = post.ID
	p.Title = post.Title
	p.Content = post.Content
	p.CreatedAt = post.CreatedAt
	p.UpdatedAt = post.UpdatedAt
	return p
}

func (post Post) GetId() uuid.UUID {
	return post.ID
}

func (post *Post) Matches(query string) bool {
	// TODO: Implement
	return false
}

type PostDTO struct {
	ID uuid.UUID `json:"id,omitempty"`
	baseDTO
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (postDTO PostDTO) GetId() uuid.UUID {
	return postDTO.ID
}

func (postDTO PostDTO) Validate() []*ErrorResponse {
	return ValidateStruct(postDTO)
}
