package models

type Page[T any] struct {
	Content        []T `json:"content"`
	Page           int `json:"page"`
	Size           int `json:"size"`
	TotalRegisters int `json:"total"`
}

func NewPage[T any](content []T, page int, size int, totalRegisters int) *Page[T] {
	return &Page[T]{
		Content:        content,
		Page:           page,
		Size:           size,
		TotalRegisters: totalRegisters,
	}
}
