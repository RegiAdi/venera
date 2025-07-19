package kernel

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/RegiAdi/venera/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConnection struct {
	Client *mongo.Client
	DB     *mongo.Database
}

// NewMongoConnection establishes a new connection to the MongoDB database.
// It returns a MongoConnection struct containing the client and database instances,
// or an error if the connection or ping fails.
func NewMongoConnection() (*MongoConnection, error) {
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(config.GetMongoURI()).SetServerAPIOptions(serverAPI)

	// Use a context with a timeout to prevent the application from hanging
	// during startup if the database is unavailable.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create a new client and connect to the server
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongo: %w", err)
	}

	// Send a ping to confirm a successful connection
	db := client.Database(config.GetMongoDatabase())
	if err := db.RunCommand(ctx, bson.D{{"ping", 1}}).Decode(&bson.M{}); err != nil {
		_ = client.Disconnect(context.Background()) // Attempt to clean up
		return nil, fmt.Errorf("failed to ping mongo: %w", err)
	}
	log.Println("Pinged your deployment. You successfully connected to MongoDB!")

	return &MongoConnection{
		Client: client,
		DB:     db,
	}, nil
}
