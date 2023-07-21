package database

import (
	"log"
)

// registeredModels is a slice of all models that will be synced with the database
var registeredModels = []interface{}{}

// RegisterModel registers a model to be synced with the database
func RegisterModel(model interface{}) {
	registeredModels = append(registeredModels, model)
}

func clearDatabase() {
	log.Println("Clearing database")
	err := DB.Exec("DROP SCHEMA public CASCADE;").Error
	DB.Exec("CREATE SCHEMA IF NOT EXISTS " + "public")
	if err != nil {
		log.Fatal("Error clearing database")
	}
	log.Println("Database cleared")
}

// SyncDataBase syncs the database with the models using gorm AutoMigrate
func SyncDataBase() {
	log.Println("Syncing database")
	err := DB.Debug().AutoMigrate(
		//&models.User{},
		registeredModels...,
	)
	if err != nil {
		log.Fatal("Error syncing database")
	}
	log.Println("Database synced")
}
