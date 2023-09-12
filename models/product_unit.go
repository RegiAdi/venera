package models

import "time"

type ProductUnit struct {
	Name      string    `json:"name,omitempty" bson:"name,omitempty"`
	Price     uint64    `json:"price,omitempty" bson:"price,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
