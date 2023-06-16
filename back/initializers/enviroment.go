package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnviromentVariables(enviroment string) {
	err := godotenv.Load(".env." + enviroment)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
