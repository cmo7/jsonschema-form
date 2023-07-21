package middleware

import (
	"nartex/ngr-stack/app/models"
	"nartex/ngr-stack/database"
	"time"

	"github.com/gofiber/fiber/v2"
)

// RequestTimer is a middleware that logs the time it takes to process a request
func RequestTimer(c *fiber.Ctx) error {
	start := time.Now()
	err := c.Next()
	stop := time.Now()
	a := models.Analytic{}
	a.Route = c.Route().Path
	a.Time = stop.Sub(start)
	database.DB.Create(&a)
	return err
}
