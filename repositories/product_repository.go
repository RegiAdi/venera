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
)

type ProductRepository struct {
	DB *mongo.Database
}

func NewProductRepository(DB *mongo.Database) *ProductRepository {
	return &ProductRepository{
		DB: DB,
	}
}

func (repository *ProductRepository) GetProductByID(id string) (models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var product models.Product
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return product, kernel.ErrInvalidObjectID
	}

	err = repository.DB.Collection("products").FindOne(ctx, bson.M{"_id": objID}).Decode(&product)
	if err != nil {
		return product, kernel.ErrProductNotFound
	}

	return product, nil
}

func (repository *ProductRepository) GetAllProducts() ([]models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var products []models.Product
	results, err := repository.DB.Collection("products").Find(ctx, bson.M{})
	if err != nil {
		return products, kernel.ErrProductNotFound
	}

	defer results.Close(ctx)

	for results.Next(ctx) {
		var product models.Product
		if err = results.Decode(&product); err != nil {
			return products, kernel.ErrProductNotFound
		}
		products = append(products, product)
	}

	return products, nil
}

func (repository *ProductRepository) CreateProduct(product models.Product) (responses.ProductResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	product.CreatedAt = helpers.GetCurrentTime()
	product.UpdatedAt = helpers.GetCurrentTime()

	result, err := repository.DB.Collection("products").InsertOne(ctx, product)
	if err != nil {
		return responses.ProductResponse{}, kernel.ErrProductCreateFailed
	}

	return responses.ProductResponse{
		ID:          result.InsertedID.(primitive.ObjectID),
		Name:        product.Name,
		Description: product.Description,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}, nil
}

func (repository *ProductRepository) UpdateProduct(id string, product models.Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return kernel.ErrInvalidObjectID
	}

	product.UpdatedAt = helpers.GetCurrentTime()
	update := bson.M{
		"$set": product,
	}

	result, err := repository.DB.Collection("products").UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil || result.ModifiedCount < 1 {
		return kernel.ErrProductUpdateFailed
	}

	return nil
}

func (repository *ProductRepository) DeleteProduct(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return kernel.ErrInvalidObjectID
	}

	result, err := repository.DB.Collection("products").DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil || result.DeletedCount < 1 {
		return kernel.ErrProductDeleteFailed
	}

	return nil
}
