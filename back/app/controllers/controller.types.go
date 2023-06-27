package controllers

import (
	"nartex/ngr-stack/app/models"
	"nartex/ngr-stack/i18n"

	"github.com/gofiber/fiber/v2"
)

type CrudController interface {
	GetAll() fiber.Handler
	Get() fiber.Handler
	Create() fiber.Handler
	Update() fiber.Handler
	Delete() fiber.Handler
}

type Resource interface {
	models.Identifiable
	models.Validable
	models.Dtoable
}

type CrudControllerOptions struct {
	Locale             i18n.Locale
	ResourceName       string
	ResourceSlug       string
	ResourcePluralName string
	ResourcePluralSlug string
}

// NewStandardCrudController creates a new CrudController instance for the given resource
// T is the resource type, it must implement the Validable interface and the Dtoable interface
func NewStandardCrudController[T Resource](options CrudControllerOptions) DefaultCrudController[T] {
	return DefaultCrudController[T](options)
}

type StatusType string

const (
	SuccessStatus StatusType = "success"
	ErrorStatus   StatusType = "error"
)

type ResponseBody[T any] struct {
	Status  StatusType `json:"status"`
	Message string     `json:"message"`
	Data    T          `json:"data"`
}

func NewResponseBody[T interface{}](status StatusType, message string, data T) *ResponseBody[T] {
	return &ResponseBody[T]{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

type EmptyBody struct{}

type MiscelaneousController struct {
	Locale   i18n.Locale
	Handlers map[string]fiber.Handler
}

// Generic type for a page of a given resource Res
type Page[Res interface{}] struct {
	Content        []Res `json:"content"`
	Page           int   `json:"page"`
	Size           int   `json:"size"`
	TotalRegisters int   `json:"total"`
}

// NewPage creates a new Page instance for the given resource
// Res is the resource type, it must implement the Resource interface
func NewPage[Res interface{}](content []Res, page int, size int, totalRegisters int) *Page[Res] {
	return &Page[Res]{
		Content:        content,
		Page:           page,
		Size:           size,
		TotalRegisters: totalRegisters,
	}
}
