package routes

import (
	"example/json-schema/controllers"
	"example/json-schema/middleware"

	"github.com/gofiber/fiber/v2"
)

func AddRoutes(app *fiber.App) {

	api := app.Group("/api")

	// Health Check
	api.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": "Server is running",
		})
	})
	// User Routes
	authRoutes(&api)
	userRoutes(&api)
}

func authRoutes(api *fiber.Router) {
	router := *api
	auth := router.Group("/auth")
	auth.Post("/register", controllers.SignUpUser).Name("signUpUser")
	auth.Post("/login", controllers.LogInUser).Name("logInUser")
	auth.Get("/logout", controllers.LogOutUser).Name("logOutUser")
	auth.Get("/getCurrentUser", middleware.DeserializeUser, controllers.GetCurrentUser).Name("getCurrentUser")
}

func userRoutes(api *fiber.Router) {
	router := *api
	user := router.Group("/user")
	// Auxiliar Routes
	user.Get("/schema/:schemaName", controllers.UserSchema).Name("userSchema")

	user.Get("/:id", controllers.GetUser).Name("getUser")
	user.Get("/", controllers.GetAllUsers).Name("getUsers")

}
