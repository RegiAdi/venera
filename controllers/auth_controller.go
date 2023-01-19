package controllers

import (
	"context"
	"log"
	"time"

	"github.com/RegiAdi/pos-mobile-backend/bootstrap"
	"github.com/RegiAdi/pos-mobile-backend/helpers"
	"github.com/RegiAdi/pos-mobile-backend/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Login(c *fiber.Ctx) error {
	userCollection := bootstrap.MongoDB.Database.Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	var request models.User
	var user models.User

	if err := c.BodyParser(&request); err != nil {
		log.Println(err)

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error": err,
		})
	}

	err := userCollection.FindOne(ctx, bson.D{{"username", request.Username}}).Decode(&user)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "User not found",
			"error": err,
		})
	}

	if !helpers.CheckPasswordHash(request.Password, user.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Password do not match",
			"error": nil,
		})
	}
	
	apiToken, _ := helpers.GenerateAPIToken()
	apiTokenExpirationDate := helpers.GenerateAPITokenExpiration()
	filter := bson.D{{"_id", user.ID}}
	update := bson.D{
		{"$set", bson.D{
			{"api_token", apiToken}, 
			{"token_expires_at", apiTokenExpirationDate},
			{"updated_at", helpers.GetCurrentTime()},
		},
	}}

	err = userCollection.FindOneAndUpdate(context.TODO(), filter, update, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&user)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to generate API Token",
			"error": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "User authenticated successfully",
		"data": user,
	})
}

func Register(c *fiber.Ctx) error {
	userCollection := bootstrap.MongoDB.Database.Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	user := new(models.User)

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error": err,
		})
	}	

	usernameCount, _ := userCollection.CountDocuments(ctx, bson.D{{"username", user.Username}})
	if usernameCount > 0 {
        return c.Status(fiber.StatusConflict).JSON(fiber.Map{
            "success": false,
            "message": "Username already exist",
            "error": nil,
        })
	}

	emailCount, _ := userCollection.CountDocuments(ctx, bson.D{{"email", user.Email}})
	if emailCount > 0 {
        return c.Status(fiber.StatusConflict).JSON(fiber.Map{
            "success": false,
            "message": "Email already exist",
            "error": nil,
        })
	}

	password, err := helpers.HashPassword(user.Password)
	if err != nil {
        return c.Status(500).JSON(fiber.Map{
            "success": false,
            "message": "User registration failed",
            "error": err,
        })
	}

	user.Password = password
	user.CreatedAt = primitive.NewDateTimeFromTime(helpers.GetCurrentTime())
	user.UpdatedAt = primitive.NewDateTimeFromTime(helpers.GetCurrentTime())

	result, err := userCollection.InsertOne(ctx, user)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{
            "success": false,
            "message": "User registration failed",
            "error":   err,
        })
    }

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "data": result,
        "success": true,
        "message": "User created successfully",
    })
}