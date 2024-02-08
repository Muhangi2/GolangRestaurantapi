package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Invoice struct {
	ID               primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" validate:"required,min=2,max=100"`
	Invoice_id       string             `json:"invoice_id" `
	Order_id         string             `json:"order_id" `
	Payment_method   *string            `json:"payment_method" validate:"eq=creditcard|eq=debitcard|eq=cash|eq=paypal|eq=''"`
	Payment_status   *string            `json:"payment_status" validate:"eq=paid|eq=unpaid|eq=pending"`
	Payment_due_date time.Time          `json:"payment_due_date" `
	Created_at       time.Time          `json:"created_at"`
	Updated_at       time.Time          `json:"updated_at"`
}
