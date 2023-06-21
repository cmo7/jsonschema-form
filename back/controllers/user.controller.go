package controllers

import (
	"example/json-schema/initializers"
	"example/json-schema/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "id is required",
			"data":    nil,
		})
	}
	var user = new(models.User)
	initializers.DB.First(&user, "id = ?", id)
	if user.ID == nil || user.ID.String() == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": fiber.ErrNotFound.Message,
			"status":  "error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.FilterUserRecord(user))
}

func GetAllUsers(c *fiber.Ctx) error {
	var payload models.PageableRequest
	err := c.BodyParser(&payload)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}
	log.Println(payload)

	errors := models.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": errors,
		})
	}

	pageSize := payload.Size
	page := payload.Page

	if pageSize == 0 {
		pageSize = 10
	}
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * pageSize

	var users []models.User
	initializers.DB.Limit(pageSize).Offset(offset).Find(&users)
	var count int64
	initializers.DB.Model(&users).Count(&count)

	log.Println("page: ", page)
	log.Println("pageSize: ", pageSize)
	log.Println("offset: ", offset)
	log.Println("count: ", count)
	log.Println(users)
	var usersResponse []models.UserResponse
	for _, user := range users {
		usersResponse = append(usersResponse, models.FilterUserRecord(&user))
	}

	pageable := models.NewPage[models.UserResponse](usersResponse, len(users), page, pageSize)

	return c.Status(fiber.StatusOK).JSON(pageable)
}
