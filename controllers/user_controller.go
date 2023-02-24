package controllers

import (
	"context"
	"time"

	"github.com/RegiAdi/hatchet/bootstrap"
	"github.com/RegiAdi/hatchet/responses"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUserInfo(c *fiber.Ctx) error {
	userCollection := bootstrap.MongoDB.Database.Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var userResponse responses.UserResponse
	reqHeader := c.GetReqHeaders()

	err := userCollection.FindOne(ctx, bson.D{{"api_token", reqHeader["Authorization"]}}).Decode(&userResponse)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "User not found",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "User retrieved successfully",
		"data":    userResponse,
	})
}
