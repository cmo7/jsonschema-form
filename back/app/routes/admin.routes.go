package routes

import (
	"nartex/ngr-stack/app/controllers"
	"nartex/ngr-stack/app/middleware"

	"github.com/gofiber/fiber/v2"
)

// adminRoutes is the router for the admin routes
// Creates a new fiber app and mounts the admin routes

func adminRoutes() *fiber.App {
	router := fiber.New()

	// Administration Routes Group, token and admin role required

	protected := router.Group("/").
		Use(middleware.ValidateToken).     // Routes in the group require a valid token
		Use(middleware.OnlyAdmin).         // Routes in the group require the user to be an admin
		Use(middleware.RefreshAccessToken) // Routes in the group Refresh the token if it's valid

		// Administration Routes Group, token and admin role required
	protected.Get("/config", controllers.GetServerConfiguration).Name("Get App Configuration")

	return router
}
