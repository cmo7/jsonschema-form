package routes

import (
	"nartex/ngr-stack/app/controllers"
	"nartex/ngr-stack/app/middleware"
	"nartex/ngr-stack/i18n"

	"github.com/gofiber/fiber/v2"
)

func postRoutes() *fiber.App {
	router := fiber.New()
	controller := controllers.PostController

	public := router.Group("/")

	public.Get("/", controller.GetAll()).
		Name(i18n.S(i18n.GET_ALL, controller.ResourcePluralName))

	public.Get("/:id", controller.Get()).
		Name(i18n.S(i18n.GET, controller.ResourceName))

	protected := router.Group("/").
		Use(middleware.ValidateToken).
		Use(middleware.OnlyAdmin)

	protected.
		Post("/", controller.Create()).
		Name(i18n.S(i18n.CREATE, controller.ResourceName))

	protected.
		Put("/:id", controller.Update()).
		Name(i18n.S(i18n.UPDATE, controller.ResourceName))

	return router
}
