package routes

import (
	"nartex/ngr-stack/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func apiInfoRoutes() *fiber.App {

	router := fiber.New()
	public := router.Group("/")
	public.Get("/routes", controllers.GetAllRoutes).Name("Get Routes List")

	return router
}
