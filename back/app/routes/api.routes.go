package routes

import (
	"nartex/ngr-stack/app/controllers"

	"github.com/gofiber/fiber/v2"
)

// ApiRoutes is the main router for the API
// Creates a new fiber app and mounts the api routes
// Mount new routes here
func ApiRoutes() *fiber.App {

	api := fiber.New()

	// Health Check
	api.Get("/healthcheck", controllers.HealthCheck).Name("Health Check")
	// Json Schemas for Forms Routes

	api.Mount("/auth", authRoutes())
	api.Mount("/administration", adminRoutes())
	api.Mount("/schema", schemaRoutes())
	api.Mount("/api-info", apiInfoRoutes())
	api.Mount("/user", userRoutes())
	api.Mount("/role", roleRoutes())
	api.Mount("/post", postRoutes())

	return api
}
