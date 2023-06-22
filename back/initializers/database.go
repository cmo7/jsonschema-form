package initializers

import (
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
	switch Config.Database.Engine {
	case Postgres:
		DB = connectToPostgres()
	case SQLServer:
		DB = connectToSQLServer()
	case Mysql:
		DB = connectToMySQL()
	case Sqlite3:
		DB = connectToSQLite()
	default:
		log.Fatalf("Database engine %s not supported", Config.Database.Engine)
	}

	if Config.Logger.QueryLogger {
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
		Config.Database.Host,
		Config.Database.Port,
		Config.Database.User,
		Config.Database.Password,
		Config.Database.Database,
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
