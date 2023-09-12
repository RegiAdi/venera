package config

import (
	"os"
	"strconv"
)

func GetTokenDuration() int {
	APITokenDuration, _ := strconv.Atoi(os.Getenv("APITOKEN_DURATIONDAYS"))

	return APITokenDuration
}
