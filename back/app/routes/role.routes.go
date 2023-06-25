package routes

import (
	"nartex/ngr-stack/app/controllers"
	"nartex/ngr-stack/app/middleware"

	"github.com/gofiber/fiber/v2"
)

// userRoutes is the router for the user routes
// Creates a new fiber app and mounts the user routes
func roleRoutes() *fiber.App {
	router := fiber.New()
	// Public routes, no token required
	public := router.Group("/")
	public.Get("/", controllers.Role.GetAll).Name("Get All Users")
	public.Get("/:id", controllers.Role.Get).Name("Get User")
	// Protected routes, token required, only admin
	protected := router.Group("/").Use(middleware.ValidateToken).Use(middleware.OnlyAdmin)
	protected.Post("/", controllers.Role.Create).Name("Create User")
	protected.Put("/:id", controllers.Role.Update).Name("Update User")
	protected.Delete("/:id", controllers.Role.Delete).Name("Delete User")

	return router
}
