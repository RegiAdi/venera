package controllers

import (
	"context"
	"time"

	"github.com/RegiAdi/venera/helpers"
	"github.com/RegiAdi/venera/kernel"
	"github.com/RegiAdi/venera/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func Logout(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	userCollection := kernel.Mongo.DB.Collection("users")

	reqHeader := c.GetReqHeaders()

	filter := bson.D{{Key: "api_token", Value: reqHeader["Authorization"]}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "api_token", Value: nil},
			{Key: "token_expires_at", Value: time.Time{}},
			{Key: "updated_at", Value: helpers.GetCurrentTime()},
		},
		}}

	result, err := userCollection.UpdateOne(ctx, filter, update)

	if err != nil || result.ModifiedCount < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to delete API Token",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "User logged out successfully",
	})
}

func Register(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	userCollection := kernel.Mongo.DB.Collection("users")

	user := new(models.User)

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}

	usernameCount, _ := userCollection.CountDocuments(ctx, bson.D{{Key: "username", Value: user.Username}})
	if usernameCount > 0 {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"success": false,
			"message": "Username already exist",
		})
	}

	emailCount, _ := userCollection.CountDocuments(ctx, bson.D{{Key: "email", Value: user.Email}})
	if emailCount > 0 {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"success": false,
			"message": "Email already exist",
		})
	}

	password, err := helpers.HashPassword(user.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "User registration failed",
			"error":   err,
		})
	}

	user.Password = password
	user.CreatedAt = helpers.GetCurrentTime()
	user.UpdatedAt = helpers.GetCurrentTime()

	result, err := userCollection.InsertOne(ctx, user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "User registration failed",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data":    result,
		"success": true,
		"message": "User created successfully",
	})
}
