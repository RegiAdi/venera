package responses

import (
	"time"
)

type UserLoginResponse struct {
	ID         string    `json:"id,omitempty" bson:"_id,omitempty"`
	Username   string    `json:"username,omitempty" bson:"username,omitempty"`
	APIToken   string    `json:"api_token,omitempty" bson:"api_token,omitempty"`
	DeviceName string    `json:"device_name,omitempty" bson:"device_name,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
