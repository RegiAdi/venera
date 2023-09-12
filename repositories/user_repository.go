package repositories

import (
	"context"
	"time"

	"github.com/RegiAdi/hatchet/responses"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	DB *mongo.Database
}

func NewUserRepository(DB *mongo.Database) *UserRepository {
	return &UserRepository{
		DB,
	}
}

func (userRepository *UserRepository) GetUserByAPIToken(APIToken string) (responses.UserResponse, error) {
	userCollection := userRepository.DB.Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var userResponse responses.UserResponse

	err := userCollection.FindOne(ctx, bson.D{{Key: "api_token", Value: APIToken}}).Decode(&userResponse)

	return userResponse, err
}
