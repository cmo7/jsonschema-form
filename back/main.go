package main

import (
	"fmt"
	"log"
	"nartex/ngr-stack/app"
	"nartex/ngr-stack/app/routes"
	"nartex/ngr-stack/codegen"
	"nartex/ngr-stack/config"
	"nartex/ngr-stack/database"
	"nartex/ngr-stack/database/seeders"
	"os/exec"
)

// Initialize the services and do codegen tasks.
func init() {
	// Load the configuration from env files.
	config.LoadConfig()
	// Establish the connection to the database.
	database.ConnectToDataBase()

	// If enabled in config, migrate registered models to the database
	if config.Generate.AutoMigrate {
		database.SyncDataBase()
	}

	// If enabled in config, use Tygo (as an external executable) to generate the
	// front end types
	if config.Generate.FrontTypes {
		cmd := exec.Command("tygo", "generate")
		if err := cmd.Run(); err != nil {
			fmt.Println("Error: ", err)
		}
	}

	// Pre-generate the Json Schemas for the registered models.
	codegen.GenerateJsonFormSchemas()

	// If selected, seed the database with mock data.
	if config.Debug.DatabaseSeed {
		seeders.Seed()
	}
}

func main() {
	// Create App
	app := app.BootstrapApp()

	// Mount API App Routes
	app.Mount("/api", routes.ApiRoutes())

	// Run Server
	portString := fmt.Sprintf(":%d", config.WebServer.Port)

	if config.WebServer.TLS {
		err := app.ListenTLS(portString, config.WebServer.CertFile, config.WebServer.KeyFile)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := app.Listen(portString)
		if err != nil {
			log.Fatal(err)
		}
	}

}
