package controllers

import (
	"nartex/ngr-stack/app/models"
	"nartex/ngr-stack/database"
	"nartex/ngr-stack/i18n"

	"github.com/gofiber/fiber/v2"
)

// We create a struct to hold the handlers
// and we initialize it in the init function
// This way we can use the handlers in the routes
// without contaminating the package namespace with handlers

var UserController struct {
	Locale i18n.Locale

	ResourceName       string
	ResourceSlug       string
	ResourcePluralName string
	ResourcePluralSlug string

	GetAll fiber.Handler
	Get    fiber.Handler
	Create fiber.Handler
	Update fiber.Handler
	Delete fiber.Handler
}

func init() {
	UserController.GetAll = userGetAll
	UserController.Get = userGet
	UserController.Create = userCreate
	UserController.Update = userUpdate
	UserController.Delete = userDelete

	UserController.ResourceName = "User"
	UserController.ResourceSlug = "user"
	UserController.ResourcePluralName = "Users"
	UserController.ResourcePluralSlug = "users"
}

func userGet(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {

		return c.Status(fiber.StatusBadRequest).
			JSON(NewResponseBody(
				ErrorStatus,
				i18n.GetWithValue(UserController.Locale, i18n.REQUIRED, "id"),
				EmptyBody{},
			))
	}
	var user = new(models.User)
	database.DB.First(&user, "id = ?", id)
	if user.ID.String() == "" {
		return c.Status(fiber.StatusNotFound).
			JSON(NewResponseBody(
				ErrorStatus,
				i18n.GetWithValue(UserController.Locale, i18n.NOT_FOUND, "User"),
				EmptyBody{},
			))
	}

	return c.Status(fiber.StatusOK).
		JSON(NewResponseBody[models.UserDTO](
			SuccessStatus,
			"User found",
			user.ToDto(),
		))
}

func userGetAll(c *fiber.Ctx) error {
	// Parse query parameters
	page, size, err := pageParams(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(NewResponseBody(
				ErrorStatus,
				err.Error(),
				EmptyBody{},
			))
	}

	// Extract corresponding page of users
	var users []models.User
	result := database.DB.
		Scopes(database.Paginate(page, size)).
		Find(&users)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(NewResponseBody(
				ErrorStatus,
				result.Error.Error(),
				EmptyBody{},
			))
	}

	// Count total registers
	var totalRegistersI64 int64
	database.DB.Model(&models.User{}).Count(&totalRegistersI64)
	totalRegisters := int(totalRegistersI64)

	// Filter users to response and create pagea
	var usersResponse []models.UserDTO
	for _, user := range users {
		usersResponse = append(usersResponse, user.ToDto())
	}

	return c.Status(fiber.StatusOK).
		JSON(NewResponseBody[*Page[models.UserDTO]](
			SuccessStatus,
			i18n.GetWithValue(UserController.Locale, i18n.FOUND, "Users"),
			NewPage[models.UserDTO](
				usersResponse,
				page,
				size,
				totalRegisters,
			),
		))
}

func userCreate(c *fiber.Ctx) error {
	var payload models.SignUpInput
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(NewResponseBody(
				ErrorStatus,
				err.Error(),
				EmptyBody{},
			))
	}
	err := payload.Validate()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(NewResponseBody(
				ErrorStatus,
				fiber.ErrBadRequest.Message,
				EmptyBody{},
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
			JSON(NewResponseBody(
				ErrorStatus,
				i18n.GetWithValue(UserController.Locale, i18n.REQUIRED, "id"),
				EmptyBody{},
			))
	}

	var payload models.SignUpInput
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(NewResponseBody(
				ErrorStatus,
				err.Error(),
				EmptyBody{},
			))
	}
	err := payload.Validate()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(NewResponseBody(
				ErrorStatus,
				fiber.ErrBadRequest.Message,
				EmptyBody{},
			))
	}

	user := models.User{}
	database.DB.First(&models.User{}, "id = ?", id)
	if user.ID.String() == "" {
		return c.Status(fiber.StatusNotFound).
			JSON(NewResponseBody(
				ErrorStatus,
				i18n.GetWithValue(UserController.Locale, i18n.NOT_FOUND, "User"),
				EmptyBody{},
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
			JSON(NewResponseBody(
				ErrorStatus,
				i18n.GetWithValue(UserController.Locale, i18n.REQUIRED, "id"),
				EmptyBody{},
			))
	}

	user := models.User{}
	database.DB.First(&models.User{}, "id = ?", id)
	if user.ID.String() == "" {
		return c.Status(fiber.StatusNotFound).
			JSON(NewResponseBody(
				ErrorStatus,
				i18n.GetWithValue(UserController.Locale, i18n.NOT_FOUND, "User"),
				EmptyBody{},
			))
	}

	database.DB.Delete(&user)

	return c.Next()
}
