package routes

import (
	"nartex/ngr-stack/app/controllers"
	"nartex/ngr-stack/app/middleware"
	"nartex/ngr-stack/i18n"

	"github.com/gofiber/fiber/v2"
)

func genericRoutes(controller controllers.CrudController) *fiber.App {
	router := fiber.New()

	options := controller.GetOptions()

	public := router.Group("/")

	public.
		Get("/", controller.GetAll()).
		Name(i18n.S(i18n.GET_ALL, options.ResourcePluralName))

	public.
		Get("/:id", controller.Get()).
		Name(i18n.S(i18n.GET, options.ResourceName))

	protected := router.Group("/").
		Use(middleware.ValidateToken).
		Use(middleware.OnlyAdmin)

	protected.
		Post("/", controller.Create()).
		Name(i18n.S(i18n.CREATE, options.ResourceName))
	protected.
		Put("/:id", controller.Update()).
		Name(i18n.S(i18n.UPDATE, options.ResourceName))
	protected.
		Delete("/:id", controller.Delete()).
		Name(i18n.S(i18n.DELETE, options.ResourceName))

	return router
}
