package responses

import "time"

type UserResponse struct {
	ID         string    `json:"id,omitempty" bson:"_id,omitempty"`
	Usercode   string    `json:"usercode,omitempty" bson:"usercode,omitempty" faker:"uuid_digit"`
	Username   string    `json:"username,omitempty" bson:"username,omitempty" faker:"username"`
	Email      string    `json:"email,omitempty" bson:"email,omitempty" faker:"email"`
	Fullname   string    `json:"fullname,omitempty" bson:"fullname,omitempty" faker:"name"`
	Phone      string    `json:"phone,omitempty" bson:"phone,omitempty" faker:"phone_number"`
	Address    string    `json:"address,omitempty" bson:"address,omitempty" faker:"real_address"`
	DeviceName string    `json:"device_name,omitempty" bson:"device_name,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
