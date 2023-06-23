package routes

import (
	"example/json-schema/controllers"
	"example/json-schema/middleware"

	"github.com/gofiber/fiber/v2"
)

// AdminRoutes is the router for the admin routes
// Creates a new fiber app and mounts the admin routes

func adminRoutes() *fiber.App {
	router := fiber.New()

	// Administration Routes Group, token and admin role required
	administration := router.Group("/administration").
		Use(middleware.ValidateToken).     // Routes in the group require a valid token
		Use(middleware.OnlyAdmin).         // Routes in the group require the user to be an admin
		Use(middleware.RefreshAccessToken) // Routes in the group Refresh the token if it's valid

		// Administration Routes Group, token and admin role required
	administration.Get("/config", controllers.GetServerConfiguration).Name("Get App Configuration")
	administration.Get("/routes", controllers.GetAllRoutes).Name("Get Routes List")

	return router
}
