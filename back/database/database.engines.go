package database

import (
	"fmt"
	"log"
	"nartex/ngr-stack/config"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ConnectToPostgres specificly connects to a postgres database using gorm
func connectToPostgres() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Database.Host,
		config.Database.Port,
		config.Database.User,
		config.Database.Password,
		config.Database.Database,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database")
	}
	db.Exec("CREATE SCHEMA IF NOT EXISTS " + "public")
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	fmt.Printf("Connected to database %s\n", config.Database.Database)
	return db
}

// ConnectToSQLServer specificly connects to a sql server database using gorm
func connectToSQLServer() *gorm.DB {
	// TODO: Implement
	log.Fatal("SQL Server not supported yet")
	return nil
}

// ConnectToMySQL specificly connects to a mysql database using gorm
func connectToMySQL() *gorm.DB {
	// TODO: Implement
	log.Fatal("MySQL not supported yet")
	return nil
}

// ConnectToSQLite specificly connects to a sqlite database using gorm
func connectToSQLite() *gorm.DB {
	// TODO: Test and complete
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database")
	}
	return db
}
