package controllers

import "github.com/gofiber/fiber/v2"

type CrudController interface {
	GetAll() fiber.Handler
	Get() fiber.Handler
	Create() fiber.Handler
	Update() fiber.Handler
	Delete() fiber.Handler
}
