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

	err := userCollection.FindOne(ctx, bson.D{{"api_token", apiToken}}).Decode(&userResponse)

	return userResponse, err 
}