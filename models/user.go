package models

import "time"

type User struct {
	Id              string    `json:"id,omitempty" bson:"_id,omitempty"`
	Usercode        string    `json:"usercode,omitempty" bson:"usercode,omitempty"`
	Username        string    `json:"username,omitempty" bson:"username,omitempty"`
	Password        string    `json:"password,omitempty" bson:"password,omitempty"`
	Email           string    `json:"email,omitempty" bson:"email,omitempty"`
	Fullname        string    `json:"fullname,omitempty" bson:"fullname,omitempty"`
	Phone           string    `json:"phone,omitempty" bson:"phone,omitempty"`
	Address         string    `json:"address,omitempty" bson:"address,omitempty"`
	ApiToken        string    `json:"api_token,omitempty" bson:"api_token,omitempty"`
	DeviceName      string    `json:"device_name,omitempty" bson:"device_name,omitempty"`
	TokenLastUsedAt time.Time `json:"token_last_used_at,omitempty" bson:"token_last_used_at,omitempty"`
	TokenExpiresAt  time.Time `json:"token_expires_at,omitempty" bson:"token_expires_at,omitempty"`
	CreatedAt       time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt       time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
