package main

import (
	"example/json-schema/app"
	"example/json-schema/config"
	"example/json-schema/database"
	"example/json-schema/initializers"
	lib "example/json-schema/lib/banners"
	"example/json-schema/routes"
	"fmt"
	"log"
)

// Enviroments
const (
	development   = "development"
	preproduction = "preproduction"
	production    = "production"
	container     = "container"
)

func init() {
	config.LoadConfig(development)
	database.ConnectToDataBase()

	generate := config.Options.Generate
	if generate.AutoMigrate {
		log.Print(lib.BoxBanner("Sincronizando Base de Datos"))
		database.SyncDataBase()
	}
	if generate.FrontTypes {
		log.Print(lib.BoxBanner("Generando Tipos de Datos"))
		initializers.GenerateFrontTypes()
	}
	log.Print(lib.BoxBanner("Generando Schemas de Formularios"))
	initializers.GenerateJsonFormSchemas()
}

func main() {
	config := config.Options
	// Create App
	app := app.BootstrapApp(config)

	// Mount API App Routes
	app.Mount("/api", routes.ApiRoutes())

	// Run Server
	webConfig := config.WebServer
	portString := fmt.Sprintf(":%d", webConfig.Port)

	if webConfig.TLS {
		err := app.ListenTLS(portString, webConfig.CertFile, webConfig.KeyFile)
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
