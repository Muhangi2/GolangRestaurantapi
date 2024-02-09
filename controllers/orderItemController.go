package controllers

import (
	"context"
	"fmt"
	"golang-Restaurantbooking/database"
	"golang-Restaurantbooking/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type neworderItem struct {
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
		if err := c.BindJSON(&orderItem); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
		}
		var orderitemStructure neworderItem
		var orders models.Order
		orders.Order_Date = time.Now()
		orders.Table_id = orderitemStructure.Table_id
		orderItemsTobeCreated := []interface{}{}
		order_id := OrderItemOrderCreator(order)

		for _, orderItem := range orderitemStructure.Order_items {
			orderItem.Order_id = order_id
			validateError := validate.Struct(orderItem)
			if validateError != nil {
				c.JSON(400, gin.H{"error": validateError.Error()})
			}

			orderItem.ID = primitive.NewObjectID()
			orderItem.Order_item_id = orderItem.ID.Hex()
			orderItem.Created_at = time.Now()
			orderItem.Updated_at = time.Now()
			var num = toFixed(*orderItem.Unit_price, 2)
			orderItem.Unit_price = &num
			orderItemsTobeCreated = append(orderItemsTobeCreated, orderItem)

		}

		result, insertErr := orderItemCollection.InsertMany(ctx, orderItemsTobeCreated)

		if insertErr != nil {
			log.Fatal(insertErr)
			c.JSON(500, gin.H{"error": "Error while creating the orderItem"})
		}
		defer cancel()
		c.JSON(201, result)

	}
}
func UpdateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		var orderItem models.OrderItem
		if err := c.BindJSON(&orderItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
			return
		}
		orderItemId := c.Param("orderItem_id")
		filter := bson.M{"order_item_id": orderItemId}

		var updateObj primitive.D
		if orderItem.Unit_price != nil {
			updateObj = append(updateObj, bson.E{"unit_price": *orderItem.Unit_price})
		}
		if orderItem.Quantity != nil {
			updateObj = append(updateObj, bson.E{"quantity": *orderItem.Quantity})
		}
		if orderItem.Food_id != nil {
			updateObj = append(updateObj, bson.E{"food_id": *orderItem.Food_id})
		}
		orderItem.Updated_at = time.Now()
		updateObj = append(updateObj, bson.E{"updated_at": orderItem.Updated_at})

		//finalizing with inserting in the mongoDb
		upsert := true
		opt := options.UpdateOptions{
			Upsert: &upsert,
		}
		result, err := orderCollection.UpdateOne(
			ctx, filter, bson.D{{"$set", updateObj}}, &opt,
		)
		if err != nil {
			msg := fmt.Sprintf("orderitem updated ")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()
		c.JSON(201, result)
	}
}
func DeleteOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func ItemsByOrder(id string) (orderItems []primitive.M, error error) {

}
