package config

import (
	"os"
	"strconv"
)

func GetTokenDuration() int {
	apiTokenDuration, _ := strconv.Atoi(os.Getenv("APITOKEN_DURATIONDAYS"))

	return apiTokenDuration
}