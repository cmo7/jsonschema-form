package database

import (
	"log"
)

var registeredModels = []interface{}{}

// RegisterModel registers a model to be synced with the database
func RegisterModel(model interface{}) {
	registeredModels = append(registeredModels, model)
}

// SyncDataBase syncs the database with the models using gorm AutoMigrate
func SyncDataBase() {
	log.Println("Syncing database")
	err := DB.AutoMigrate(
		//&models.User{},
		registeredModels...,
	)
	if err != nil {
		log.Fatal("Error syncing database")
	}
	log.Println("Database synced")
}
