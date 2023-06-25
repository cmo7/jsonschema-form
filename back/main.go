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

func init() {
	config.LoadConfig()
	database.ConnectToDataBase()

	if config.Generate.AutoMigrate {
		//banners.Print("Sincronizando Base de Datos")
		database.SyncDataBase()
	}
	if config.Generate.FrontTypes {
		//banners.Print("Generando Tipos de Datos")
		//codegen.GenerateFrontTypes()
		cmd := exec.Command("tygo", "generate")
		if err := cmd.Run(); err != nil {
			fmt.Println("Error: ", err)
		}
	}
	//banners.Print("Generando Schemas de Formularios")
	codegen.GenerateJsonFormSchemas()

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
