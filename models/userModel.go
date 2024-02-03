package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" validate:"required,min=2,max=100"`
	User_id       string             `json:"user_id"`
	Username      *string            `json:"username" validate:"required,min=2,max=100"`
	Password      *string            `json:"password" validate:"required,min=2"`
	Email         *string            `json:"email" validate:"required,email"`
	Avatar        *string            `json:"avatar" `
	Phone         *string            `json:"phone"  validate:"required,"`
	Token         *string            `json:"token" `
	Refresh_token *string            `json:"refresh_token" `
	Created_at    *string            `json:"created_at" `
	Updated_at    *string            `json:"updated_at" `
}
