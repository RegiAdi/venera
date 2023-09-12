package repositories

import (
	"context"
	"time"

	"github.com/RegiAdi/hatchet/responses"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	GetUserByApiToken(apiToken string) (responses.UserResponse, error)
}

type userRepository struct {
	db *mongo.Database
}

func NewUserRepository(db *mongo.Database) *userRepository {
	return &userRepository{
		db,
	}
}

func (userRepo *userRepository) GetUserByApiToken(apiToken string) (responses.UserResponse, error) {
	userCollection := userRepo.db.Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var userResponse responses.UserResponse

	err := userCollection.FindOne(ctx, bson.D{{Key: "api_token", Value: apiToken}}).Decode(&userResponse)

	return userResponse, err
}
