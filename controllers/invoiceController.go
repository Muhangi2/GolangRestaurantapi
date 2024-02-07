package controllers

import (
	"context"
	"golang-Restaurantbooking/database"
	"golang-Restaurantbooking/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type InvoiceViewFormat struct {
	Invoice_id       string
	Order_id         string
	Payment_method   *string
	Payment_status   *string
	Payment_due_date time.Time
	Payment_due      interface{}
	Table_number     interface{}
	Order_details    interface{}
}

var invoiceCollection *mongo.Collection = database.OpenCollection(database.Client, "invoice")

func GetInvoices() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		result, err := invoiceCollection.Find(context.TODO(), bson.M{})
		defer cancel()
		if err != nil {
			c.JSON(500, gin.H{"error": "Error while fetching the invoice"})
		}
		var allInvoices []bson.M
		if err = result.All(ctx, &allInvoices); err != nil {
			c.JSON(500, gin.H{"error": "Error while fetching the invoice"})
		}
		c.JSON(http.StatusOk, allInvoices)
	}

}
func GetInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		invoiceId := c.Param("invoice_id")
		var invoice models.Invoice
		err := invoiceCollection.FindOne(ctx, bson.M{"invoice_id": invoiceId}).Decode(&invoice)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "error occured while fetching the food and food doesnt exist",
			})
		}

		var invoiceView InvoiceViewFormat

		allOrderItems, err := ItemsByOrder(invoice.Order_id)
		invoiceView.Order_id = invoice.Order_id
		invoiceView.Payment_due_date = invoice.Payment_due_date

		invoiceView.Payment_method="null"
		if invoice.Payment_method!=nil{
		  invoiceView.Payment_method=*&invoice.Payment_method
		}
		
	}
}
func CreateInvoice() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
func UpdateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func DeleteInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
