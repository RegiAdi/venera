package kernel

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type AppKernel struct {
	DB     *mongo.Database
	Server *fiber.App
}

func NewAppKernel() *AppKernel {
	LoadEnv()
	NewMongoConnection()

	return &AppKernel{
		DB:     Mongo.DB,
		Server: fiber.New(),
	}
}

const (
	StatusFailed       = "FAILED"
	StatusSuccess      = "SUCCESS"
	StatusUnauthorized = fiber.StatusUnauthorized
	StatusBadRequest   = fiber.StatusBadRequest
)
