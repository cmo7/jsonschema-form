package routes

import (
	"nartex/ngr-stack/app/controllers"

	"github.com/gofiber/fiber/v2"
)

// schemaRoutes is the router for the schema routes
// Creates a new fiber app and mounts the schema routes
func schemaRoutes() *fiber.App {
	router := fiber.New()
	router.Get("/:schemaName", controllers.GetSchema).Name("Get Json Form Schema")
	router.Get("/ui/:schemaName", controllers.GetUiSchema).Name("Get Json Form Ui Schema")
	return router
}
