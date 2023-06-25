package controllers

import (
	"nartex/ngr-stack/app/models"
	"nartex/ngr-stack/app/types"
	"nartex/ngr-stack/database"
	"nartex/ngr-stack/services/data"

	"github.com/gofiber/fiber/v2"
)

var Role struct {
	GetAll fiber.Handler
	Get    fiber.Handler
	Create fiber.Handler
	Update fiber.Handler
	Delete fiber.Handler
}

// We create a struct to hold the handlers
// and we initialize it in the init function
// This way we can use the handlers in the routes
// without contaminating the package namespace with handlers
func init() {
	Role.GetAll = roleGetAll
	Role.Get = roleGet
	Role.Create = roleCreate
	Role.Update = roleUpdate
	Role.Delete = roleDelete
}

// TODO: Add pagination
// TODO: Review Copilot suggestions
// TODO: Add filters
func roleGetAll(c *fiber.Ctx) error {

	page, size, err := pageParams(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(types.NewResponseBody(
				types.Error,
				err.Error(),
				types.Empty{},
			))
	}
	var roles []models.Role
	result := database.DB.
		Scopes(database.Paginate(page, size)).
		Find(&roles)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(types.NewResponseBody(
				types.Error,
				result.Error.Error(),
				types.Empty{},
			))
	}

	var rolesResponse []models.RoleResponse
	for _, role := range roles {
		rolesResponse = append(rolesResponse, models.FilterRoleRecord(&role))
	}

	totalRegisters := data.Count(&models.Role{})

	return c.Status(fiber.StatusOK).
		JSON(types.NewResponseBody[*models.Page[models.RoleResponse]](
			types.Success,
			"Roles found",
			models.NewPage[models.RoleResponse](
				rolesResponse,
				page,
				size,
				totalRegisters,
			),
		))
}

func roleGet(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).
			JSON(types.NewResponseBody(
				types.Error,
				"id is required",
				types.Empty{},
			))
	}
	var role = new(models.Role)
	database.DB.First(&role, "id = ?", id)
	if role.ID == nil || role.ID.String() == "" {
		return c.Status(fiber.StatusNotFound).
			JSON(types.NewResponseBody(
				types.Error,
				"Role not found",
				types.Empty{},
			))
	}

	return c.Status(fiber.StatusOK).
		JSON(types.NewResponseBody[models.RoleResponse](
			types.Success,
			"Role found",
			models.FilterRoleRecord(role),
		))
}

func roleCreate(c *fiber.Ctx) error {
	var role models.Role
	if err := c.BodyParser(&role); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(types.NewResponseBody(
				types.Error,
				"Could not parse JSON",
				types.Empty{},
			))
	}

	database.DB.Create(&role)

	return c.Status(fiber.StatusCreated).
		JSON(types.NewResponseBody[models.RoleResponse](
			types.Success,
			"Role created",
			models.FilterRoleRecord(&role),
		))
}

func roleUpdate(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).
			JSON(types.NewResponseBody(
				types.Error,
				"id is required",
				types.Empty{},
			))
	}

	var role = new(models.Role)
	database.DB.First(&role, "id = ?", id)
	if role.ID == nil || role.ID.String() == "" {
		return c.Status(fiber.StatusNotFound).
			JSON(types.NewResponseBody(
				types.Error,
				"Role not found",
				types.Empty{},
			))
	}

	if err := c.BodyParser(&role); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(types.NewResponseBody(
				types.Error,
				"Could not parse JSON",
				types.Empty{},
			))
	}

	database.DB.Save(&role)

	return c.Status(fiber.StatusOK).
		JSON(types.NewResponseBody[models.RoleResponse](
			types.Success,
			"Role updated",
			models.FilterRoleRecord(role),
		))
}

func roleDelete(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).
			JSON(types.NewResponseBody(
				types.Error,
				"id is required",
				types.Empty{},
			))
	}

	var role = new(models.Role)
	database.DB.First(&role, "id = ?", id)
	if role.ID == nil || role.ID.String() == "" {
		return c.Status(fiber.StatusNotFound).
			JSON(types.NewResponseBody(
				types.Error,
				"Role not found",
				types.Empty{},
			))
	}

	database.DB.Delete(&role)

	return c.Status(fiber.StatusOK).
		JSON(types.NewResponseBody[models.RoleResponse](
			types.Success,
			"Role deleted",
			models.FilterRoleRecord(role),
		))
}
