package initializers

import (
	"example/json-schema/models"
	"fmt"
	"log"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectToDataBase() {
	databaseEngine := os.Getenv("DATABASE_ENGINE")
	switch databaseEngine {
	case "postgres":
		connectToPostgres()
	case "sqlserver":
		connectToSQLServer()
	default:
		log.Fatal("Database engine not supported")
	}
}

func SyncDataBase() {
	log.Println("Syncing database")
	err := DB.AutoMigrate(
		&models.User{},
	)
	if err != nil {
		log.Fatal("Error syncing database")
	}
}

func connectToPostgres() {
	p := os.Getenv("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		log.Fatal("Error parsing port")
	}
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		port,
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database")
	}
	DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	DB.Logger = logger.Default.LogMode(logger.Info)
}

func connectToSQLServer() {
	log.Fatal("SQL Server not supported yet")
}
