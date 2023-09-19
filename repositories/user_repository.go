package repositories

import (
	"context"
	"time"

	"github.com/RegiAdi/hatchet/helpers"
	"github.com/RegiAdi/hatchet/models"
	"github.com/RegiAdi/hatchet/responses"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (userRepository *UserRepository) GetAuthenticatedUser(APIToken string) (models.User, error) {
	userCollection := userRepository.DB.Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User

	err := userCollection.FindOne(ctx, bson.D{{Key: "api_token", Value: APIToken}}).Decode(&user)

	return user, err
}

func (userRepository *UserRepository) GetUserByAPIToken(APIToken string) (responses.UserResponse, error) {
	userCollection := userRepository.DB.Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var userResponse responses.UserResponse

	err := userCollection.FindOne(ctx, bson.D{{Key: "api_token", Value: APIToken}}).Decode(&userResponse)

	return userResponse, err
}

func (userRepository *UserRepository) UpdateAPITokenExpirationTime(userID string) error {
	userCollection := userRepository.DB.Collection("users")

	currentTime := helpers.GetCurrentTime()
	userObjectID, _ := primitive.ObjectIDFromHex(userID)

	filter := bson.D{{Key: "_id", Value: userObjectID}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "api_token", Value: ""},
			{Key: "token_expires_at", Value: time.Time{}},
			{Key: "token_last_used_at", Value: time.Time{}},
			{Key: "updated_at", Value: currentTime},
		},
		}}

	_, err := userCollection.UpdateOne(context.TODO(), filter, update)

	return err
}

func (userRepository *UserRepository) UpdateAPITokenLastUsedTime(userID string) error {
	userCollection := userRepository.DB.Collection("users")

	currentTime := helpers.GetCurrentTime()
	userObjectID, _ := primitive.ObjectIDFromHex(userID)

	filter := bson.D{{Key: "_id", Value: userObjectID}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "token_last_used_at", Value: currentTime},
			{Key: "updated_at", Value: currentTime},
		},
		}}

	_, err := userCollection.UpdateOne(context.TODO(), filter, update)

	return err
}
