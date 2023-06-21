package controllers

import (
	"example/json-schema/lib"

	"github.com/gofiber/fiber/v2"
)

func GetSchema(c *fiber.Ctx) error {
	schemaName := c.Params("schemaName")

	schema, err := lib.GetSchema(schemaName)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": fiber.ErrNotFound.Message,
			"data":    nil,
		})
	}
	return c.JSON(schema)

}
