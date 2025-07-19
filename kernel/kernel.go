package kernel

import (
	"context"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type AppKernel struct {
	DB          *mongo.Database
	Server      *fiber.App
	mongoClient *mongo.Client
}

func NewAppKernel() (*AppKernel, error) {
	LoadEnv()

	mongoConnection, err := NewMongoConnection()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database connection: %w", err)
	}

	return &AppKernel{
		DB:          mongoConnection.DB,
		Server:      fiber.New(),
		mongoClient: mongoConnection.Client,
	}, nil
}

// Shutdown gracefully shuts down the application.
func (k *AppKernel) Shutdown(ctx context.Context) {
	log.Println("Shutting down server...")
	if err := k.Server.ShutdownWithContext(ctx); err != nil {
		log.Printf("Server shutdown failed: %v", err)
	}

	if k.mongoClient != nil {
		log.Println("Closing MongoDB connection...")
		if err := k.mongoClient.Disconnect(ctx); err != nil {
			log.Printf("MongoDB disconnection failed: %v", err)
		}
	}
}

const (
	StatusFailed       = "failed"
	StatusSuccess      = "success"
	StatusUnauthorized = fiber.StatusUnauthorized
	StatusBadRequest   = fiber.StatusBadRequest
	StatusNotFound     = fiber.StatusNotFound
	StatusOK           = fiber.StatusOK
	StatusCreated      = fiber.StatusCreated
)
