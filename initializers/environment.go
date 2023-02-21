package initializers

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to load dotenv file")
	}
	log.Println("Dotenv file loaded")
}
