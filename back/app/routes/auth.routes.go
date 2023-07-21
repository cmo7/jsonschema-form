package routes

import (
	"nartex/ngr-stack/app/controllers"
	"nartex/ngr-stack/app/middleware"

	"github.com/gofiber/fiber/v2"
)

// AuthRoutes is the router for the auth routes
// Creates a new fiber app and mounts the auth routes
func authRoutes() *fiber.App {

	router := fiber.New()

	// Public Routes Group, no token required
	public := router.Group("/")
	public.Post("/register", controllers.AuthController.SignUp()).Name("Register User")
	public.Post("/login", controllers.AuthController.LogIn()).Name("Login User")

	// Protected Routes Group, token required, refresh token if valid
	protected := router.Group("/").
		Use(middleware.ValidateToken).     // Routes in the group require a valid token
		Use(middleware.RefreshAccessToken) // Routes in the group Refresh the token if it's valid
	protected.Get("/logout", controllers.AuthController.LogOut()).Name("Logout User")
	protected.Get("/refresh", controllers.AuthController.RefreshAccessToken()).Name("Refresh Token")
	protected.Use(middleware.DeserializeUser).Get("/getCurrentUser", controllers.AuthController.GetCurrentUser()).Name("Get User Corresponding to Token")

	return router
}
