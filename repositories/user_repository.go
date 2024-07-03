package repositories

import (
	"context"
	"time"

	"github.com/RegiAdi/venera/helpers"
	"github.com/RegiAdi/venera/kernel"
	"github.com/RegiAdi/venera/models"
	"github.com/RegiAdi/venera/responses"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (userRepository *UserRepository) GetUserByUsername(username string) (models.User, error) {
	userCollection := userRepository.DB.Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User

	err := userCollection.FindOne(ctx, bson.D{{Key: "username", Value: username}}).Decode(&user)

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

func (userRepository *UserRepository) UpdateAPIToken(user models.User) (responses.UserLoginResponse, error) {
	userCollection := userRepository.DB.Collection("users")

	APIToken, err := helpers.GenerateAPIToken()
	if err != nil {
		return responses.UserLoginResponse{}, kernel.ErrGenerateAPITokenFailed
	}

	APITokenExpirationDate := helpers.GenerateAPITokenExpiration()

	objectID, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return responses.UserLoginResponse{}, kernel.ErrInvalidObjectID
	}

	filter := bson.D{{Key: "_id", Value: objectID}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "api_token", Value: APIToken},
			{Key: "device_name", Value: user.DeviceName},
			{Key: "token_expires_at", Value: APITokenExpirationDate},
			{Key: "updated_at", Value: helpers.GetCurrentTime()},
		},
		}}

	var userLoginResponse responses.UserLoginResponse
	err = userCollection.FindOneAndUpdate(context.TODO(), filter, update, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&userLoginResponse)

	if err != nil {
		return responses.UserLoginResponse{}, kernel.ErrUserUpdateFailed
	}

	return userLoginResponse, err
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
