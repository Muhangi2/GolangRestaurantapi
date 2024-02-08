package controllers

import (
	"context"
	"golang-Restaurantbooking/database"
	"golang-Restaurantbooking/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// defining order collection
var orderCollection *mongo.Collection = database.OpenCollection(database.Client, "order")

func GetOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		result, err := orderCollection.Find(context.TODO(), bson.M{})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while fetching the food"})
		}
		var allorders []bson.M
		if err = result.All(ctx, &allorders); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allorders)
	}

}
func GetOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		orderId := c.Param("Order_id")
		var order models.OrderItem
		err := orderCollection.FindOne(ctx, bson.M{"order_id": orderId}).Decode(&order)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured fetching the single order"})
		}

	}
}

func CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		orderId := c.Param("Order_id")
		var order models.Order
		var table models.Table

		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
		}

		validateError := validate.Struct(order)
		if validateError != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validateError.Error()})
			return
		}
		//checking if the t exits
		err := orderCollection.FindOne(ctx, bson.M{"Order_id": order.Table_id}).Decode(&order)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Menu does not exist"})
			return
		}

	}
}
func UpdateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// func DeleteOrder() gin.HandlerFunc {
// 	return func(c *gin.Context) {

// 	}
// }
