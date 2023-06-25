package controllers

import (
	"nartex/ngr-stack/utils/jsonschemasgenerator"

	"github.com/gofiber/fiber/v2"
)

func GetSchema(c *fiber.Ctx) error {
	schemaName := c.Params("schemaName")

	schema, err := jsonschemasgenerator.GetSchema(schemaName)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": fiber.ErrNotFound.Message,
			"data":    nil,
		})
	}
	return c.JSON(schema)

}

func GetUiSchema(c *fiber.Ctx) error {
	schemaName := c.Params("schemaName")

	schema, err := jsonschemasgenerator.GetUiSchema(schemaName)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": fiber.ErrNotFound.Message,
			"data":    nil,
		})
	}
	return c.JSON(schema)

}
