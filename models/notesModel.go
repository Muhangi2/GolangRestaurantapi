package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Note struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" `
	Title      *string            `json:"title"`
	Text       *string            `json:"text"`
	Note_id    string             `json:"note_id" `
	Created_at time.Time          `json:"created_at" `
	Updated_at time.Time          `json:"updated_at" `
}
