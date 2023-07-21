package middleware

import (
	"nartex/ngr-stack/app/controllers"
	"nartex/ngr-stack/config"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Paginate is a middleware that gets the page and size query parameters and sets them in the context
func Paginate(c *fiber.Ctx) error {

	pageParam := c.Query("page", "1")
	page, err := strconv.Atoi(pageParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(controllers.NewResponseBody(
				controllers.ErrorStatus,
				err.Error(),
				controllers.EmptyData{},
			))
	}
	sizeParam := c.Query("size", "0")
	size, err := strconv.Atoi(sizeParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(controllers.NewResponseBody(
				controllers.ErrorStatus,
				err.Error(),
				controllers.EmptyData{},
			))
	}

	// In case the page or size are invalid, set them to the default values
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = config.Pagination.DefaultPageSize
	}
	if size > config.Pagination.MaxPageSize {
		size = config.Pagination.DefaultPageSize
	}

	c.Locals("page", page)
	c.Locals("size", size)

	return c.Next()
}
