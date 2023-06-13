package routes

import (
	"example/json-schema/controllers"

	"github.com/gofiber/fiber/v2"
)

func AddRoutes(app *fiber.App) {

	api := app.Group("/api")

	// User Routes
	userRoutes(api)
}

func userRoutes(api fiber.Router) {
	user := api.Group("/user")
	user.Get("/schema", controllers.UserSchema).Name("userSchema")
	user.Post("/create", controllers.UserCreate).Name("userCreate")
	user.Get("/all", controllers.UserGetAll).Name("userGetAll")
	user.Get("/:id", controllers.UserGet).Name("userGet")
	user.Put("/:id", controllers.UserUpdate).Name("userUpdate")
	user.Delete("/:id", controllers.UserDelete).Name("userDelete")
}
