package shrine

import (
	"context"
	"time"

	"github.com/RegiAdi/hatchet/helpers"
	"github.com/RegiAdi/hatchet/kernel"
	"github.com/RegiAdi/hatchet/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func New() fiber.Handler {
	return func(c *fiber.Ctx) error {
		reqHeader := c.GetReqHeaders()

		if reqHeader["Authorization"] == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"message": "Unauthorized access",
				"error":   nil,
			})
		}

		userCollection := kernel.MongoDB.Database.Collection("users")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var user models.User

		err := userCollection.FindOne(ctx, bson.D{{"api_token", reqHeader["Authorization"]}}).Decode(&user)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"message": "Unauthorized access",
				"error":   err,
			})
		}

		currentTime := helpers.GetCurrentTime()

		// check api token expiration time
		if helpers.GetCurrentTime().After(user.TokenExpiresAt.Time()) {
			filter := bson.D{{"_id", user.ID}}
			update := bson.D{
				{"$set", bson.D{
					{"api_token", ""},
					{"token_expires_at", time.Time{}},
					{"token_last_used_at", time.Time{}},
					{"updated_at", currentTime},
				},
				}}

			_, err = userCollection.UpdateOne(context.TODO(), filter, update)

			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"success": false,
					"message": "Failed to delete API Token",
					"error":   err,
				})
			}

			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"message": "API Token expired",
				"error":   nil,
			})

		}

		// save token last used time
		filter := bson.D{{"_id", user.ID}}
		update := bson.D{
			{"$set", bson.D{
				{"token_last_used_at", currentTime},
				{"updated_at", currentTime},
			},
			}}

		_, err = userCollection.UpdateOne(context.TODO(), filter, update)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": "Failed to update API Token Last Used Time",
				"error":   err,
			})
		}

		return c.Next()
	}
}
