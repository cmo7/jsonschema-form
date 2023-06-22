package initializers

import "github.com/gofiber/fiber/v2"

var App *fiber.App

func InitializeApp() {
	App = fiber.New()
}
