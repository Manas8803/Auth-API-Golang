package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func SQLURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("SQLURI")
}
