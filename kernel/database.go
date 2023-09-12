package kernel

import (
	"context"
	"log"
	"time"

	"github.com/RegiAdi/hatchet/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoConnection struct {
	Client *mongo.Client
	DB     *mongo.Database
}

var Mongo MongoConnection

func NewMongoConnection() {
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

	db := client.Database(config.GetMongoDatabase())

	Mongo = MongoConnection{
		Client: client,
		DB:     db,
	}
}
