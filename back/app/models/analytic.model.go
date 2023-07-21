package models

import (
	"nartex/ngr-stack/database"
	"time"

	"github.com/google/uuid"
)

func init() {
	database.RegisterModel(Analytic{})
}

type Analytic struct {
	ID uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	baseEntity
	Route string        `gorm:"type:varchar(255);not null"`
	Time  time.Duration `gorm:"type:int;not null"`
}

func (analytic Analytic) GetId() uuid.UUID {
	return analytic.ID
}

func (analytic Analytic) ToDto() DTO {
	a := AnalyticDTO{}
	a.ID = analytic.ID
	a.Route = analytic.Route
	a.Time = analytic.Time
	a.CreatedAt = analytic.CreatedAt
	a.UpdatedAt = analytic.UpdatedAt
	return a
}

type AnalyticDTO struct {
	ID uuid.UUID `json:"id,omitempty"`
	baseDTO
	Route string        `json:"route"`
	Time  time.Duration `json:"time"`
}

func (analyticDTO AnalyticDTO) GetId() uuid.UUID {
	return analyticDTO.ID
}

func (analyticDTO AnalyticDTO) Validate() []*ErrorResponse {
	return ValidateStruct(analyticDTO)
}
