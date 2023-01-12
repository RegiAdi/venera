package config

import (
	"os"
)

func GetMongoURI() string {
	return os.Getenv("MONGO_URI")
}

func GetMongoDatabase() string {
	return os.Getenv("MONGO_DATABASE")
}