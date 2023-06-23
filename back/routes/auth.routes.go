package routes

import (
	"example/json-schema/controllers"
	"example/json-schema/middleware"

	"github.com/gofiber/fiber/v2"
)

// AuthRoutes is the router for the auth routes
// Creates a new fiber app and mounts the auth routes
func authRoutes() *fiber.App {

	router := fiber.New()

	// Public Routes Group, no token required
	public := router.Group("/public")

	// Protected Routes Group, token required, refresh token if valid
	protected := router.Group("/").
		Use(middleware.ValidateToken).     // Routes in the group require a valid token
		Use(middleware.RefreshAccessToken) // Routes in the group Refresh the token if it's valid

		// Public auth routes
	public.Post("/register", controllers.SignUpUser).Name("Register User")
	public.Post("/login", controllers.LogInUser).Name("Login User")

	// Protected auth routes, token required
	protected.Get("/logout", controllers.LogOutUser).Name("Logout User")
	protected.Get("/refresh", controllers.RefreshAccessToken).Name("Refresh Token")
	protected.Use(middleware.DeserializeUser).Get("/getCurrentUser", controllers.GetCurrentUser).Name("Get User Corresponding to Token")

	return router
}
