package controllers

import (
	"fmt"
	"nartex/ngr-stack/database"
	"nartex/ngr-stack/utils/validation"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CrudController interface {
	GetAll() fiber.Handler
	Get() fiber.Handler
	Create() fiber.Handler
	Update() fiber.Handler
	Delete() fiber.Handler
}

type Resource interface {
	Identifiable
	validation.Validable
	Dtoable
}

type Dtoable interface {
	ToDto() interface{}
}

type Identifiable interface {
	GetId() uuid.UUID
}

type CrudControllerImpl[Res Resource] struct {
	ResourceName       string
	ResourceSlug       string
	ResourcePluralName string
	ResourcePluralSlug string
}

func (imp *CrudControllerImpl[Res]) GetAll() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// TODO: implement
		return nil
	}
}

func (imp *CrudControllerImpl[Res]) Get() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// TODO: implement
		return nil
	}
}

func (imp *CrudControllerImpl[Res]) Create() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// TODO: implement
		return nil
	}
}

func (imp *CrudControllerImpl[Res]) Update() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// TODO: implement
		return nil
	}
}

func (imp *CrudControllerImpl[Res]) Delete() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		if id == "" {
			return c.Status(fiber.StatusBadRequest).
				JSON(NewResponseBody(
					Error,
					"id is required",
					Empty{},
				))
		}

		var resource = new(Res)
		database.DB.First(resource, "id = ?", id)
		content := *resource
		if content.GetId().String() == "" {
			return c.Status(fiber.StatusNotFound).
				JSON(NewResponseBody(
					Error,
					fmt.Sprintf(messages["notFound"], imp.ResourceName),
					Empty{},
				))
		}

		err := database.DB.Delete(resource)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).
				JSON(NewResponseBody(
					Error,
					fmt.Sprintf(messages["notDeleted"], imp.ResourceName),
					Empty{},
				))
		}
		return c.Status(fiber.StatusOK).
			JSON(NewResponseBody(
				Success,
				fmt.Sprintf(messages["deleted"], imp.ResourceName),
				Empty{},
			))
	}
}

type NewCrudControllerOptions struct {
	ResourceName       string
	ResourceSlug       string
	ResourcePluralName string
	ResourcePluralSlug string
	CustomGetAll       fiber.Handler
	CustomGet          fiber.Handler
	CustomCreate       fiber.Handler
	CustomUpdate       fiber.Handler
	CustomDelete       fiber.Handler
}

// NewCrudController creates a new CrudController instance for the given resource
// T is the resource type, it must implement the Validable interface and the Dtoable interface
func NewCrudController[T Resource](options NewCrudControllerOptions) CrudControllerImpl[T] {
	controller := CrudControllerImpl[T]{
		ResourceName:       options.ResourceName,
		ResourceSlug:       options.ResourceSlug,
		ResourcePluralName: options.ResourcePluralName,
		ResourcePluralSlug: options.ResourcePluralSlug,
	}

	return controller
}

var messages = map[string]string{
	"required":          "%s is required",
	"notFound":          "%s not found",
	"invalid":           "%s is invalid",
	"alreadyExists":     "%s already exists",
	"notFoundOrInvalid": "%s not found or invalid",
	"found":             "%s found",
	"deleted":           "%s deleted",
	"updated":           "%s updated",
	"created":           "%s created",
	"notDeleted":        "%s not deleted",
}

type StatusType string

const (
	Success StatusType = "success"
	Error   StatusType = "error"
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

type Empty struct{}
