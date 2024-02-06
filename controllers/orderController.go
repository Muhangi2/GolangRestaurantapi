package controllers

import (
	"context"
	"golang-Restaurantbooking/database"
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
    var ctx,cancel =context.WithTimeout(context.Background(),100*time.Second)
	orderId:=c.Params("o")
	}
}
func CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
func UpdateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
func DeleteOrder() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
