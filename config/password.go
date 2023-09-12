package config

import (
	"os"
	"strconv"
)

func GetPasswordHashCost() int {
	cost, err := strconv.ParseInt(os.Getenv("PASSWORD_HASH_COST"), 10, 0)

	if err != nil {
		cost = 10
	}

	return int(cost)
}