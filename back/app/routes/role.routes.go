package routes

import (
	"fmt"
	"nartex/ngr-stack/app/controllers"
	"nartex/ngr-stack/app/middleware"

	"github.com/gofiber/fiber/v2"
)

// userRoutes is the router for the user routes
// Creates a new fiber app and mounts the user routes
func roleRoutes() *fiber.App {
	router := fiber.New()
	controller := controllers.RoleController

	// Public routes, no token required
	public := router.Group("/")

	public.
		Get("/", controller.GetAll()).
		Name(fmt.Sprintf("Get All %s", controller.ResourcePluralName))

	public.Get("/:id", controller.Get()).
		Name(fmt.Sprintf("Get %s", controller.ResourceName))

	// Protected routes, token required, only admin
	protected := router.Group("/").
		Use(middleware.ValidateToken).
		Use(middleware.OnlyAdmin)

	protected.
		Post("/", controller.Create()).
		Name(fmt.Sprintf("Create %s", controller.ResourceName))

	protected.
		Put("/:id", controller.Update()).
		Name(fmt.Sprintf("Update %s", controller.ResourceName))

	protected.
		Delete("/:id", controller.Delete()).
		Name(fmt.Sprintf("Delete %s", controller.ResourceName))

	return router
}
