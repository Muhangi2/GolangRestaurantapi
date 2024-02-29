package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" `
	User_id       string             `json:"user_id"`
	Username      *string            `json:"user_name" validate:"required,min=2,max=100"`
	FirstName     *string            `json:"first_name"`
	SecondName    *string            `json:"second_name"`
	Password      *string            `json:"password" validate:"required,min=2"`
	Email         *string            `json:"email" validate:"required,email"`
	Avatar        *string            `json:"avatar" `
	Phone         *string            `json:"phone" validate:"required,numeric"`
	Token         *string            `json:"token" `
	Refresh_token *string            `json:"refresh_token" `
	Created_at    time.Time          `json:"created_at" `
	Updated_at    time.Time          `json:"updated_at" `
}
