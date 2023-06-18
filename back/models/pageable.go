package models

type Page[T any] struct {
	Content []T `json:"content"`
	Total   int `json:"total"`
	Page    int `json:"page"`
	Size    int `json:"size"`
}

func NewPage[T any](content []T, total int, page int, size int) *Page[T] {
	return &Page[T]{
		Content: content,
		Total:   total,
		Page:    page,
		Size:    size,
	}
}

type PageableRequest struct {
	Page int `json:"page,omitempty"`
	Size int `json:"size,omitempty"`
}
