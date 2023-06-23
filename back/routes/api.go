package routes

import (
	"example/json-schema/controllers"

	"github.com/gofiber/fiber/v2"
)

// ApiRoutes is the main router for the API
// Creates a new fiber app and mounts the auth and administration routes, also adds a health check route
// Mount new routes here
func ApiRoutes() *fiber.App {

	api := fiber.New()

	// Health Check
	api.Get("/healthcheck", controllers.HealthCheck).Name("Health Check")
	// Json Schemas for Forms Routes
	api.Get("/schema/:schemaName", controllers.GetSchema).Name("Get Json Form Schema")

	api.Mount("/auth", authRoutes())
	api.Mount("/administration", adminRoutes())

	return api
}
