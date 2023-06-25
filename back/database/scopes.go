package database

import (
	"gorm.io/gorm"
)

func Paginate(page int, size int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		offset := (page - 1) * size
		limit := size

		return db.Offset(offset).Limit(limit)
	}
}
