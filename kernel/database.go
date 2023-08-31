package kernel

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/RegiAdi/hatchet/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDBInstance struct {
	Client   *mongo.Client
	Database *mongo.Database
}

var MongoDB MongoDBInstance

func connectDB() {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.GetMongoURI()))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connected!")

	MongoDB = MongoDBInstance{
		Client:   client,
		Database: client.Database(config.GetMongoDatabase()),
	}
}
