package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Usercode        string             `json:"usercode,omitempty" bson:"usercode,omitempty"`
	Username        string             `json:"username,omitempty" bson:"username,omitempty"`
	Password        string             `json:"password,omitempty" bson:"password,omitempty"`
	Email           string             `json:"email,omitempty" bson:"email,omitempty"`
	Fullname        string             `json:"fullname,omitempty" bson:"fullname,omitempty"`
	Phone           string             `json:"phone,omitempty" bson:"phone,omitempty"`
	Address         string             `json:"address,omitempty" bson:"address,omitempty"`
	APIToken        string             `json:"api_token,omitempty" bson:"api_token,omitempty"`
	DeviceName      string             `json:"device_name,omitempty" bson:"device_name,omitempty"`
	TokenLastUsedAt primitive.DateTime `json:"token_last_used_at,omitempty" bson:"token_last_used_at,omitempty"`
	TokenExpiresAt  primitive.DateTime `json:"token_expires_at,omitempty" bson:"token_expires_at,omitempty"`
	CreatedAt       primitive.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt       primitive.DateTime `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
