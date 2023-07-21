package controllers

import (
	"nartex/ngr-stack/config"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// This function is used to parse the page and size parameters from the query string
// and return the values as integers
// If the page or size parameters are invalid, the default values are returned
func pageParams(c *fiber.Ctx) (page int, size int, err error) {
	pageParam := c.Query("page", "1")
	page, err = strconv.Atoi(pageParam)
	if err != nil {
		return 0, 0, err
	}
	sizeParam := c.Query("size", "0")
	size, err = strconv.Atoi(sizeParam)
	if err != nil {
		return 0, 0, err
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

	return page, size, nil
}
