package database

import (
	"fmt"
	"log"
	"nartex/ngr-stack/config"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectToDataBase() {
	fmt.Println("Connecting to database")
	// Select the database engine from the config file
	fmt.Printf("Engine: %s\n", config.Database.Engine)
	switch config.Database.Engine {
	case config.Postgres:
		DB = connectToPostgres()
	case config.SQLServer:
		DB = connectToSQLServer()
	case config.Mysql:
		DB = connectToMySQL()
	case config.Sqlite3:
		DB = connectToSQLite()
	default:
		log.Fatalf("Database engine %s not supported", config.Database.Engine)
	}

	if config.Logger.QueryLogger {
		DB.Logger = logger.Default.LogMode(logger.Info)
	}
}
