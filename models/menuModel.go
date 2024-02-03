package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Menu struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" validate:"required,min=2,max=100"`
	Menu_id    string             `json:"menu_id,omitempty" `
	Name       *string            `json:"name,omitempty" `
	Category   *string            `json:"category,omitempty" `
	Start_date  time.Time         `json:"start_date,omitempty" `
	End_date    time.Time          `json:"end_date,omitempty" `
	Created_at  time.Time         `json:"created_at,omitempty" `
	Updated_at time.Time          `json:"updated_at,omitempty" `
}
