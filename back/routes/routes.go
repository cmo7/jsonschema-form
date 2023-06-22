package routes

import (
	"example/json-schema/controllers"
	"example/json-schema/middleware"

	"github.com/gofiber/fiber/v2"
)

func AddApiRoutes(app *fiber.App) {

	api := app.Group("/api")

	// Health Check
	api.Get("/healthcheck", controllers.HealthCheck).Name("Health Check")

	// Json Schemas for Forms Routes
	api.Get("/schema/:schemaName", controllers.GetSchema).Name("Get Json Form Schema")

	// Auth Routes
	auth := api.Group("/auth")
	auth.Post("/register", controllers.SignUpUser).Name("Register User")
	auth.Post("/login", controllers.LogInUser).Name("Login User")
	auth.Get("/logout", controllers.LogOutUser).Name("Logout User")
	// Token required auth routes
	authProtected := auth.Group("/", middleware.ValidateToken)
	authProtected.Use(middleware.DeserializeUser).Get("/getCurrentUser", controllers.GetCurrentUser).Name("Get User Corresponding to Token")

	administration := api.Group("/administration").Use(middleware.ValidateToken).Use(middleware.OnlyAdmin)
	administration.Get("/config", controllers.GetServerConfiguration).Name("Get App Configuration")
	administration.Get("/routes", controllers.GetAllRoutes).Name("Get Routes List")
}
