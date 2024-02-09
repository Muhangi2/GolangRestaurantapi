package controllers

import (
	"context"
	"golang-Restaurantbooking/database"
	"golang-Restaurantbooking/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type orderItem struct {
	Table_id    *string
	Order_items []models.OrderItem
}

var orderItemCollection *mongo.Collection = database.OpenCollection(database.Client, "orderItem")

func GetOrderItems() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		result, err := orderItemCollection.Find(context.TODO(), bson.M{})
		defer cancel()
		if err != nil {
			c.JSON(500, gin.H{"error": "Error while fetching the orderItem"})
		}
		var allOrderItems []bson.M
		if err = result.All(ctx, &allOrderItems); err != nil {
			c.JSON(500, gin.H{"error": "Error while fetching the orderItem"})
		}
		c.JSON(200, allOrderItems)

	}

}
func GetOrderItemsByOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		orderItemId := c.Param("orderItem_id")
		allOrderItems, err := ItemsByOrder(orderItemId)

		if err != nil {
			c.JSON(500, gin.H{"error": "error occured while fetching the orderItem"})
		}
		c.JSON(200, allOrderItems)
	}
}
func GetOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		orderItemId := c.Param("orderItem_id")
		var orderItem models.OrderItem
		err := orderItemCollection.FindOne(ctx, bson.M{"order_item_id": orderItemId}).Decode(&orderItem)
		defer cancel()
		if err != nil {
			c.JSON(500, gin.H{"error": "error occured while fetching the orderItem"})
		}
		c.JSON(200, orderItem)

	}
}
func CreateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
     var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	 var orderItem models.OrderItem
	  if err:=  c.BindJSON(&orderItem);
	  err != nil {
		  c.JSON(400, gin.H{"error": err.Error()})
	  }
	  //validate the orderItem
	  validateError := validate.Struct(orderItem)
    if validateError != nil {
		c.JSON(400, gin.H{"error": validateError.Error()})
	}

	  orderItem.Created_at = time.Now()
	  orderItem.Updated_at = time.Now()
	  orderItem.ID = primitive.NewObjectID()
	  orderItem.Order_item_id = orderItem.ID.Hex()
	  result, insertErr := orderItemCollection.InsertOne(ctx, orderItem)
      	  if insertErr != nil {
					  c.JSON(500, gin.H{"error": "Error while creating the orderItem"})
		  }
		  defer cancel()
		  c.JSON(201, orderItem)

	}
}
func UpdateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
func DeleteOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func ItemsByOrder(id string) (orderItems []primitive.M, error error) {

}
