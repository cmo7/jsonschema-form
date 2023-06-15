// Package controllers ...
// Este paquete contiene los controladores de la aplicación.
// Cada controlador es un archivo que contiene las funciones callback de las rutas.
// Cada función callback recibe un parámetro de tipo *fiber.Ctx y devuelve un error.
// El parámetro *fiber.Ctx contiene toda la información sobre la petición HTTP y la respuesta HTTP.
// El error se utiliza para devolver errores HTTP como 404 Not Found, 500 Internal Server Error, etc.

package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func MainPage(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Hello World",
	})
}
