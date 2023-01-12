package bootstrap

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/RegiAdi/pos-mobile-backend/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoInstance struct {
	Client *mongo.Client
	DB *mongo.Database
}

var mongoInstance MongoInstance

func connectDB() {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.GetMongoURI()))	
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
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

	mongoInstance = MongoInstance{
		Client: client,
		DB: client.Database(config.GetMongoDatabase()),
	}
}