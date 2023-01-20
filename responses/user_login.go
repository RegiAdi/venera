package responses

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserLoginResponse struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username   string             `json:"username,omitempty" bson:"username,omitempty"`
	APIToken   string             `json:"api_token,omitempty" bson:"api_token,omitempty"`
	DeviceName string             `json:"device_name,omitempty" bson:"device_name,omitempty"`
	CreatedAt  primitive.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt  primitive.DateTime `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}