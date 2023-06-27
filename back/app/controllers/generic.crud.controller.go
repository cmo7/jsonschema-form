package controllers

import (
	"nartex/ngr-stack/database"
	"nartex/ngr-stack/i18n"

	"github.com/gofiber/fiber/v2"
)

type DefaultCrudController[Res Resource] struct {
	Locale             i18n.Locale
	ResourceName       string
	ResourceSlug       string
	ResourcePluralName string
	ResourcePluralSlug string
}

// Function to create a GetAll handler for a given resource.
func (imp *DefaultCrudController[Res]) GetAll() fiber.Handler {
	return func(c *fiber.Ctx) error {
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

		// Extract corresponding page of resources
		var resources []Res
		result := database.DB.
			Scopes(database.Paginate(page, size)).
			Find(&resources)
		if result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).
				JSON(NewResponseBody(
					ErrorStatus,
					result.Error.Error(),
					EmptyBody{},
				))
		}

		// Count total registers
		var resource Res
		var totalRegistersI64 int64
		database.DB.Model(resource).Count(&totalRegistersI64)
		totalRegisters := int(totalRegistersI64)

		// Filter users to response and create pagea
		var resourceResponse []interface{}
		for _, resource := range resources {
			resourceResponse = append(resourceResponse, resource.ToDto())
		}

		return c.Status(fiber.StatusOK).
			JSON(NewResponseBody[*Page[interface{}]](
				SuccessStatus,
				i18n.GetWithValue(imp.Locale, i18n.FOUND, imp.ResourcePluralName),
				NewPage[interface{}](
					resourceResponse,
					page,
					size,
					totalRegisters,
				),
			))
	}
}

// Generic Get One by Id implementation for Resource Res
func (imp *DefaultCrudController[Res]) Get() fiber.Handler {
	// Returns anonymous function parametriced as needed
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		if id == "" {

			return c.Status(fiber.StatusBadRequest).
				JSON(NewResponseBody(
					ErrorStatus,
					"id is required",
					EmptyBody{},
				))
		}
		var resource = new(Res)
		database.DB.First(&resource, "id = ?", id)
		content := *resource
		if content.GetId().String() == "" {
			return c.Status(fiber.StatusNotFound).
				JSON(NewResponseBody(
					ErrorStatus,
					i18n.GetWithValue(imp.Locale, i18n.NOT_FOUND, imp.ResourceName),
					EmptyBody{},
				))
		}

		return c.Status(fiber.StatusOK).
			JSON(NewResponseBody[interface{}](
				SuccessStatus,
				i18n.GetWithValue(imp.Locale, i18n.FOUND, imp.ResourceName),
				content.ToDto(),
			))
	}
}

func (imp *DefaultCrudController[Res]) Create() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var payload Res

		// Parse body
		if err := c.BodyParser(&payload); err != nil {
			return c.Status(fiber.StatusBadRequest).
				JSON(NewResponseBody(
					ErrorStatus,
					err.Error(),
					EmptyBody{},
				))
		}

		// Validate payload
		err := payload.Validate()
		if err != nil {
			return c.Status(fiber.StatusBadRequest).
				JSON(NewResponseBody(
					ErrorStatus,
					fiber.ErrBadRequest.Message,
					EmptyBody{},
				))
		}

		// Create resource
		result := database.DB.Create(&payload)
		if result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).
				JSON(NewResponseBody(
					ErrorStatus,
					result.Error.Error(),
					EmptyBody{},
				))
		}

		return c.Status(fiber.StatusCreated).
			JSON(NewResponseBody[interface{}](
				SuccessStatus,
				i18n.GetWithValue(imp.Locale, i18n.CREATED, imp.ResourceName),
				payload.ToDto(),
			))
	}
}

func (imp *DefaultCrudController[Res]) Update() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		if id == "" {
			return c.Status(fiber.StatusBadRequest).
				JSON(NewResponseBody(
					ErrorStatus,
					i18n.GetWithValue(imp.Locale, i18n.REQUIRED, "id"),
					EmptyBody{},
				))
		}

		var payload Res

		// Parse body
		if err := c.BodyParser(&payload); err != nil {
			return c.Status(fiber.StatusBadRequest).
				JSON(NewResponseBody(
					ErrorStatus,
					err.Error(),
					EmptyBody{},
				))
		}

		// Validate payload
		err := payload.Validate()
		if err != nil {
			return c.Status(fiber.StatusBadRequest).
				JSON(NewResponseBody(
					ErrorStatus,
					fiber.ErrBadRequest.Message,
					EmptyBody{},
				))
		}

		// Update resource
		var resource = new(Res)
		database.DB.First(resource, "id = ?", id)
		content := *resource
		if content.GetId().String() == "" {
			return c.Status(fiber.StatusNotFound).
				JSON(NewResponseBody(
					ErrorStatus,
					i18n.GetWithValue(imp.Locale, i18n.NOT_FOUND, imp.ResourceName),
					EmptyBody{},
				))
		}

		result := database.DB.Model(resource).Updates(payload)
		if result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).
				JSON(NewResponseBody(
					ErrorStatus,
					result.Error.Error(),
					EmptyBody{},
				))
		}

		return c.Status(fiber.StatusOK).
			JSON(NewResponseBody[interface{}](
				SuccessStatus,
				i18n.GetWithValue(imp.Locale, i18n.UPDATED, imp.ResourceName),
				payload.ToDto(),
			))
	}
}

func (imp *DefaultCrudController[Res]) Delete() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		if id == "" {
			return c.Status(fiber.StatusBadRequest).
				JSON(NewResponseBody(
					ErrorStatus,
					i18n.GetWithValue(imp.Locale, i18n.REQUIRED, "id"),
					EmptyBody{},
				))
		}

		var resource = new(Res)
		database.DB.First(resource, "id = ?", id)
		content := *resource
		if content.GetId().String() == "" {
			return c.Status(fiber.StatusNotFound).
				JSON(NewResponseBody(
					ErrorStatus,
					i18n.GetWithValue(imp.Locale, i18n.NOT_FOUND, imp.ResourceName),
					EmptyBody{},
				))
		}

		err := database.DB.Delete(resource)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).
				JSON(NewResponseBody(
					ErrorStatus,
					i18n.GetWithValue(imp.Locale, i18n.NOT_DELETED, imp.ResourceName),
					EmptyBody{},
				))
		}
		return c.Status(fiber.StatusOK).
			JSON(NewResponseBody(
				SuccessStatus,
				i18n.GetWithValue(imp.Locale, i18n.DELETED, imp.ResourceName),
				EmptyBody{},
			))
	}
}
