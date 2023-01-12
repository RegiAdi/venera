package config

import (
	"os"
)

func GetAppENV() string {
	return os.Getenv("APP_ENV")
}

func GetAppPort() string {
	return os.Getenv("APP_PORT")
}