package database

import (
	"example/json-schema/config"
	"example/json-schema/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectToDataBase() {
	fmt.Println("Connecting to database")
	// Select the database engine from the config file
	switch config.Options.Database.Engine {
	case config.Postgres:
		DB = connectToPostgres()
	case config.SQLServer:
		DB = connectToSQLServer()
	case config.Mysql:
		DB = connectToMySQL()
	case config.Sqlite3:
		DB = connectToSQLite()
	default:
		log.Fatalf("Database engine %s not supported", config.Options.Database.Engine)
	}

	if config.Options.Logger.QueryLogger {
		DB.Logger = logger.Default.LogMode(logger.Info)
	}
}

// SyncDataBase syncs the database with the models using gorm AutoMigrate
func SyncDataBase() {
	log.Println("Syncing database")
	err := DB.AutoMigrate(
		&models.User{},
	)
	if err != nil {
		log.Fatal("Error syncing database")
	}
	log.Println("Database synced")
}

// ConnectToPostgres specificly connects to a postgres database using gorm
func connectToPostgres() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Options.Database.Host,
		config.Options.Database.Port,
		config.Options.Database.User,
		config.Options.Database.Password,
		config.Options.Database.Database,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database")
	}
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	return db
}

// ConnectToSQLServer specificly connects to a sql server database using gorm
func connectToSQLServer() *gorm.DB {
	log.Fatal("SQL Server not supported yet")
	return nil
}

// ConnectToMySQL specificly connects to a mysql database using gorm
func connectToMySQL() *gorm.DB {
	log.Fatal("MySQL not supported yet")
	return nil
}

// ConnectToSQLite specificly connects to a sqlite database using gorm
func connectToSQLite() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database")
	}
	return db
}
