package kernel

import (
	"log"

	"github.com/joho/godotenv"
)

func loadENV() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
