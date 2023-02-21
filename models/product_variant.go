package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ProductVariant struct {
	Name      string             `json:"name,omitempty" bson:"name,omitempty"`
	Price     uint64             `json:"price,omitempty" bson:"price,omitempty"`
	Units     []ProductUnit      `json:"units,omitempty" bson:"units,omitempty"`
	CreatedAt primitive.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt primitive.DateTime `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
