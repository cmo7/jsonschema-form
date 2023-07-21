package controllers

import (
	"nartex/ngr-stack/app/models"

	"github.com/gofiber/fiber/v2"
)

type CrudController interface {
	GetAll() fiber.Handler
	Get() fiber.Handler
	Create() fiber.Handler
	Update() fiber.Handler
	Delete() fiber.Handler
	GetOptions() CrudControllerOptions
}

// Resources must implement the Resource interface. This interface is composed
// of the Identifiable interface, the Validable interface and the Dtoable interface
type Resource interface {
	models.Entity
}

type CrudControllerOptions struct {
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

// Valid StatusType values are "success" and "error"
type StatusType string

const (
	SuccessStatus StatusType = "success"
	ErrorStatus   StatusType = "error"
)

// ResponseBody is a generic type for a response body with data of type T
type ResponseBody[T any] struct {
	Status  StatusType `json:"status"`
	Message string     `json:"message"`
	Data    T          `json:"data"`
}

// NewResponseBody creates a new ResponseBody instance for the given data
func NewResponseBody[T interface{}](status StatusType, message string, data T) *ResponseBody[T] {
	return &ResponseBody[T]{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

// EmptyData is a generic type for a response body with no data
type EmptyData struct{}

// Generic type for a page of a given resource Res
// Res is the resource type, it must implement the Resource interface
type Page struct {
	Content        []models.DTO `json:"content"`
	Page           int          `json:"page"`
	Size           int          `json:"size"`
	TotalRegisters int          `json:"total"`
}

// NewPage creates a new Page instance for the given resource
// Res is the resource type, it must implement the Resource interface
func NewPage(content []models.DTO, page int, size int, totalRegisters int) *Page {
	return &Page{
		Content:        content,
		Page:           page,
		Size:           size,
		TotalRegisters: totalRegisters,
	}
}
