package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Menu struct {
	ID         primitive.ObjectID `json:"_id," bson:"_id," validate:"required,min=2,max=100"`
	Menu_id    string             `json:"food_id" `
	Name       *string            `json:"name" `
	Category   *string            `json:"category" `
	Start_date time.Time          `json:"start_date" `
	End_date   time.Time          `json:"end_date" `
	Created_at time.Time          `json:"created_at" `
	Updated_at time.Time          `json:"updated_at" `
}
