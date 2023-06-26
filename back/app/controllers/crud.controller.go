package controllers

import (
	"nartex/ngr-stack/utils/validation"

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
	validation.Validable
	Dtoable
}

type Dtoable interface {
	ToDto() interface{}
}

type CrudControllerImpl[T Resource] struct {
}

func (c *CrudControllerImpl[T]) GetAll() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// TODO: implement
		return nil
	}
}

func (c *CrudControllerImpl[T]) Get() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// TODO: implement
		return nil
	}
}

func (c *CrudControllerImpl[T]) Create() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// TODO: implement
		return nil
	}
}

func (c *CrudControllerImpl[T]) Update() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// TODO: implement
		return nil
	}
}

func (c *CrudControllerImpl[T]) Delete() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return nil
	}
}

// NewCrudController creates a new CrudController instance for the given resource
// T is the resource type, it must implement the Validable interface and the Dtoable interface
func NewCrudController[T Resource]() CrudController {
	return &CrudControllerImpl[T]{}
}
