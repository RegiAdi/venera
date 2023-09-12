package controllers

import (
	"context"
	"time"

	"github.com/RegiAdi/hatchet/helpers"
	"github.com/RegiAdi/hatchet/kernel"
	"github.com/RegiAdi/hatchet/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetProduct(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	productCollection := kernel.Mongo.DB.Collection("products")

	var product models.Product

	objID, _ := primitive.ObjectIDFromHex(c.Params("id"))
	err := productCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&product)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Product not found",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Product retrieved successfully",
		"data":    product,
	})
}

func GetProducts(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	productCollection := kernel.Mongo.DB.Collection("products")

	var products []models.Product

	results, err := productCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Products not found",
			"error":   err,
		})
	}

	defer results.Close(ctx)

	for results.Next(ctx) {
		var product models.Product

		if err = results.Decode(&product); err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"message": "Products not found",
				"error":   err,
			})
		}

		products = append(products, product)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Products retrieved successfully",
		"data":    products,
	})
}

func CreateProduct(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	productCollection := kernel.Mongo.DB.Collection("products")

	product := new(models.Product)

	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}

	product.CreatedAt = helpers.GetCurrentTime()
	product.UpdatedAt = helpers.GetCurrentTime()

	result, err := productCollection.InsertOne(ctx, product)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Create new product failed",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data":    result,
		"success": true,
		"message": "Product created successfully",
	})
}

func UpdateProduct(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	productCollection := kernel.Mongo.DB.Collection("products")

	product := new(models.Product)

	if err := c.BodyParser(product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}

	objID, err := primitive.ObjectIDFromHex(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Product not found",
			"error":   err,
		})
	}

	update := bson.M{
		"$set": product,
	}

	result, err := productCollection.UpdateOne(ctx, bson.M{"_id": objID}, update)

	if err != nil || result.ModifiedCount < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to update Product",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Product updated successfully",
	})
}

func DeleteProduct(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	productCollection := kernel.Mongo.DB.Collection("products")

	objID, _ := primitive.ObjectIDFromHex(c.Params("id"))
	result, err := productCollection.DeleteOne(ctx, bson.M{"_id": objID})

	if err != nil || result.DeletedCount < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to delete product",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Product deleted successfully",
	})
}
