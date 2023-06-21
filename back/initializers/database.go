package initializers

import (
	"example/json-schema/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectToDataBase() {
	switch Config.Database.Engine {
	case "postgres":
		connectToPostgres()
	case "sqlserver":
		connectToSQLServer()
	default:
		log.Fatal("Database engine not supported")
	}

	if os.Getenv("DB_QUERY_LOGGING") == "true" {
		DB.Logger = logger.Default.LogMode(logger.Info)
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
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		Config.Database.Host,
		Config.Database.Port,
		Config.Database.User,
		Config.Database.Password,
		Config.Database.Database,
	)
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database")
	}
	DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")

}

func connectToSQLServer() {
	log.Fatal("SQL Server not supported yet")
}
