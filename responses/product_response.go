package responses

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductResponse struct {
	ID          primitive.ObjectID `json:"_id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
}

type ProductsResponse struct {
	Products []ProductResponse `json:"products"`
}
