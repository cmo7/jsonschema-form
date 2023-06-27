package seeders

import (
	"log"
	"nartex/ngr-stack/app/models"
	"nartex/ngr-stack/database"
	"nartex/ngr-stack/database/factories"
)

func Seed() {

	log.Println("Seeding...")

	log.Println("Refreshing database...")
	// Refresh the database
	database.DB.Exec("TRUNCATE TABLE user_roles CASCADE")
	database.DB.Exec("TRUNCATE TABLE roles CASCADE")
	database.DB.Exec("TRUNCATE TABLE users CASCADE")

	log.Println("Creating roles and users...")
	// Create roles
	adminRole := factories.RoleFactory{}.CreateOneWithData(&models.Role{
		Name: "admin",
	})
	userRole := factories.RoleFactory{}.CreateOneWithData(&models.Role{
		Name:               "user",
		DefaultForNewUsers: true,
	})

	database.DB.Create([]*models.Role{adminRole, userRole})
	// Create users
	normalUsers := factories.UserFactory{}.CreateMany(100)
	database.DB.Create(&normalUsers)
	// Assign roles to users
	for _, user := range normalUsers {
		database.DB.Model(&user).Association("Roles").Append([]*models.Role{userRole})
	}

	// Create admin user
	adminUser := factories.UserFactory{}.CreateOne()
	database.DB.Create(&adminUser)
	// Assign roles to admin user
	database.DB.Model(&adminUser).Association("Roles").Append([]*models.Role{userRole, adminRole})

	newUserRole := factories.RoleFactory{}.CreateOneWithData(&models.Role{
		Name:               "newUser",
		DefaultForNewUsers: true,
	})
	database.DB.Create(newUserRole)

	// Create posts
	posts := factories.PostFactory{}.CreateMany(1)
	database.DB.Create(&posts)
}
