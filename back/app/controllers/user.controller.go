package controllers

import (
	"log"
	"nartex/ngr-stack/app/models"
	"nartex/ngr-stack/app/types"
	"nartex/ngr-stack/database"
	"nartex/ngr-stack/services/data"

	"github.com/gofiber/fiber/v2"
)

// We create a struct to hold the handlers
// and we initialize it in the init function
// This way we can use the handlers in the routes
// without contaminating the package namespace with handlers
var User struct {
	GetAll fiber.Handler
	Get    fiber.Handler
	Create fiber.Handler
	Update fiber.Handler
	Delete fiber.Handler
}

func init() {
	User.GetAll = userGetAll
	User.Get = userGet
	User.Create = userCreate
	User.Update = userUpdate
	User.Delete = userDelete
}

func userGet(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {

		return c.Status(fiber.StatusBadRequest).
			JSON(types.NewResponseBody(
				types.Error,
				"id is required",
				types.Empty{},
			))
	}
	var user = new(models.User)
	database.DB.First(&user, "id = ?", id)
	if user.ID == nil || user.ID.String() == "" {
		return c.Status(fiber.StatusNotFound).
			JSON(types.NewResponseBody(
				types.Error,
				"User not found",
				types.Empty{},
			))
	}

	return c.Status(fiber.StatusOK).
		JSON(types.NewResponseBody[models.UserResponse](
			types.Success,
			"User found",
			models.FilterUserRecord(user),
		))
}

// TODO: Add filters
func userGetAll(c *fiber.Ctx) error {
	log.Println("userGetAll")
	// Parse query parameters
	page, size, err := pageParams(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(types.NewResponseBody(
				types.Error,
				err.Error(),
				types.Empty{},
			))
	}

	// Extract corresponding page of users
	var users []models.User
	result := database.DB.
		Scopes(database.Paginate(page, size)).
		Find(&users)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(types.NewResponseBody(
				types.Error,
				result.Error.Error(),
				types.Empty{},
			))
	}

	// Count total registers
	totalRegisters := data.Count(&models.User{})

	// Filter users to response and create pagea
	var usersResponse []models.UserResponse
	for x, user := range users {
		log.Println("user", x, " -> ", user)
		usersResponse = append(usersResponse, models.FilterUserRecord(&user))
	}

	return c.Status(fiber.StatusOK).
		JSON(types.NewResponseBody[*models.Page[models.UserResponse]](
			types.Success,
			"Users found",
			models.NewPage[models.UserResponse](
				usersResponse,
				page,
				size,
				totalRegisters,
			),
		))
}

// Todo: Implement this method
func userCreate(c *fiber.Ctx) error {
	var payload models.SignUpInput
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(types.NewResponseBody(
				types.Error,
				err.Error(),
				types.Empty{},
			))
	}
	err := models.ValidateStruct(payload)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(types.NewResponseBody(
				types.Error,
				fiber.ErrBadRequest.Message,
				types.Empty{},
			))
	}

	user := models.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: payload.Password,
		Avatar:   payload.Avatar,
		Provider: "created_by_admin",
	}

	database.DB.Create(&user)

	return c.Next()
}

func userUpdate(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {

		return c.Status(fiber.StatusBadRequest).
			JSON(types.NewResponseBody(
				types.Error,
				"id is required",
				types.Empty{},
			))
	}

	var payload models.SignUpInput
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(types.NewResponseBody(
				types.Error,
				err.Error(),
				types.Empty{},
			))
	}
	err := models.ValidateStruct(payload)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(types.NewResponseBody(
				types.Error,
				fiber.ErrBadRequest.Message,
				types.Empty{},
			))
	}

	user := models.User{}
	database.DB.First(&models.User{}, "id = ?", id)
	if user.ID == nil || user.ID.String() == "" {
		return c.Status(fiber.StatusNotFound).
			JSON(types.NewResponseBody(
				types.Error,
				"User not found",
				types.Empty{},
			))
	}

	user.Name = payload.Name
	user.Email = payload.Email
	user.Password = payload.Password
	user.Avatar = payload.Avatar
	user.Provider = "modified_by_admin"

	database.DB.Save(&user)
	return c.Next()
}

func userDelete(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).
			JSON(types.NewResponseBody(
				types.Error,
				"id is required",
				types.Empty{},
			))
	}

	user := models.User{}
	database.DB.First(&models.User{}, "id = ?", id)
	if user.ID == nil || user.ID.String() == "" {
		return c.Status(fiber.StatusNotFound).
			JSON(types.NewResponseBody(
				types.Error,
				"User not found",
				types.Empty{},
			))
	}

	database.DB.Delete(&user)

	return c.Next()
}
