package controllers

import (
	"example/json-schema/initializers"
	"example/json-schema/models"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/swaggest/jsonschema-go"
	"gorm.io/gorm/clause"
)

func UserSchema(c *fiber.Ctx) error {
	reflector := jsonschema.Reflector{}
	schema, err := reflector.Reflect(models.UserDTO{})
	log.Println(schema)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": fiber.ErrInternalServerError.Message,
			"data":    err,
		})
	}
	return c.JSON(schema)
}

func UserCreate(c *fiber.Ctx) error {
	user := new(models.User)
	log.Println(c.Body())
	// PARSE
	err := c.BodyParser(user)
	friends := user.Friends
	log.Printf("Out: %+v", friends)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{
				"message": fiber.ErrBadRequest.Message,
				"data":    err,
			})
	}
	// STORE
	err = initializers.DB.Create(&user).Error
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{
				"status":  fiber.StatusInternalServerError,
				"message": fiber.ErrInternalServerError.Message,
			})
	}
	// ASSOCIATE
	for _, friendId := range friends {
		var friend models.User
		initializers.DB.First(&friend, friendId)
		initializers.DB.Model(&user).Association("Friends").Append(&friend)
	}

	// RESPONSE
	return c.Status(fiber.StatusCreated).
		JSON(fiber.Map{
			"status":  fiber.StatusCreated,
			"message": "Created user",
			"data":    user.ID,
		})
}

func UserGetAll(c *fiber.Ctx) error {
	var users []models.User
	if initializers.DB.
		Model(&models.User{}).
		Preload(clause.Associations).
		Find(&users).Error != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{
				"message": fiber.ErrInternalServerError.Message,
			})
	}
	var userDTOs []models.UserDTO
	for _, user := range users {
		userDTOs = append(userDTOs, ToDTO(&user))
	}
	return c.JSON(userDTOs)
}

func UserGet(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if initializers.DB.First(&user, id).Error != nil {
		return c.Status(fiber.StatusNotFound).
			JSON(fiber.Map{
				"message": fiber.ErrNotFound.Message,
			})
	}
	return c.JSON(user)
}

func UserUpdate(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if initializers.DB.First(&user, id).Error != nil {
		return c.Status(fiber.StatusNotFound).
			JSON(fiber.Map{
				"message": fiber.ErrNotFound.Message,
			})
	}

	err := c.BodyParser(user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{
				"message": fiber.ErrBadRequest.Message,
				"data":    err,
			})
	}

	err = initializers.DB.Create(&user).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{
				"status":  fiber.StatusInternalServerError,
				"message": fiber.ErrInternalServerError.Message,
				"data":    err,
			})
	}

	return c.Status(fiber.StatusCreated).
		JSON(fiber.Map{
			"status":  fiber.StatusCreated,
			"message": "Created user",
			"data":    user.ID,
		})
}

func UserDelete(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	initializers.DB.Delete(&user, id)
	return c.JSON(fiber.Map{
		"message": "Successfully deleted",
	})
}

func ToDTO(u *models.User) models.UserDTO {

	friends := []models.UserDTO{}
	for _, friend := range u.Friends {
		log.Println("Procesando un amigo: $x", friend.ID)
		friends = append(friends, ToDTO(&friend))
	}
	return models.UserDTO{
		ID:        u.ID,
		Email:     u.Email,
		Friends:   friends,
		CreatedAt: u.CreatedAt,
	}
}

func toEntity(u *models.UserDTO) models.User {
	user := models.User{}
	initializers.DB.First(&user, u.ID)
	return user
}
