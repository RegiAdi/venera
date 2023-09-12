package config

import (
	"os"
)

func GetAppEnv() string {
	return os.Getenv("APP_ENV")
}

func GetAppPort() string {
	return os.Getenv("APP_PORT")
}

func GetAppTimezone() string {
	return os.Getenv("APP_TIMEZONE")
}
