package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FoodItem struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" validate:"required,min=2,max=100"`
	Name       *string            `json:"name,omitempty" `
	Price      *float64           `json:"price,omitempty"`
	Food_image *string            `json:"food_image,omitempty"  validate:"required"`
	CreatedAt  time.Time          `json:"created_at,omitempty" `
	UpdatedAt  time.Time          `json:"updated_at,omitempty" `
	Food_id    string            `json:"food_id,omitempty" `
	Menu_id    *string            `json:"menu_id,omitempty"  validate:"required"`
}
