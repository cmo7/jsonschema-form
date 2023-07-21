package database

import (
	"gorm.io/gorm"
)

// Paginate creates a pagination scope for gorm
func Paginate(page int, size int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		offset := (page - 1) * size
		limit := size

		return db.Offset(offset).Limit(limit)
	}
}

// WithId creates a scope for gorm that filters by id
func WithId(id string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", id)
	}
}
