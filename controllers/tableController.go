package controllers

import (
	"context"
	"fmt"
	"golang-Restaurantbooking/database"
	"golang-Restaurantbooking/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var tableCollection *mongo.Collection = database.OpenCollection(database.Client, "table")

func GetTables() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		result, err := tableCollection.Find(context.TODO(), bson.M{})
		defer cancel()
		if err != nil {
			c.JSON(500, gin.H{"error": "Error while fetching the table"})
		}
		var allTables []bson.M
		if err = result.All(ctx, &allTables); err != nil {
			c.JSON(500, gin.H{"error": "Error while fetching the table"})
		}
		c.JSON(200, allTables)

	}
}
func GetTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		tableId := c.Param("table_id")
		var table models.Table
		err := tableCollection.FindOne(ctx, bson.M{"table_id": tableId}).Decode(&table)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while fetching the food and doesnt exist "})
		}
		defer cancel()
		c.JSON(200, table)

	}
}
func CreateTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var table models.Table

		if err := c.BindJSON(&table); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		validateErr := validate.Struct(table)
		if validateErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validateErr.Error()})
			defer cancel()
			return
		}
		table.Created_at = time.Now()
		table.Updated_at = time.Now()
		table.ID = primitive.NewObjectID()
		table.Table_id = table.ID.Hex()

		result, issertErr := tableCollection.InsertOne(ctx, table)
		defer cancel()
		if issertErr != nil {
			msg := fmt.Sprintf("table wasnt created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		c.JSON(200, result)
	}
}

func UpdateTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel() // Defer context cancellation to the end of the function

		var table models.Table
		if err := c.BindJSON(&table); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		tableID := c.Param("table_id")
		filter := bson.M{"table_id": tableID}

		var updateObj bson.D
		if table.Number_of_guests != nil {
			updateObj = append(updateObj, bson.E{"number_of_guests", table.Number_of_guests})
		}
		if table.Table_number != nil {
			updateObj = append(updateObj, bson.E{"table_number", table.Table_number})
		}

		table.Updated_at = time.Now()

		opt := options.Update().SetUpsert(true)
		result, err := tableCollection.UpdateOne(
			ctx,
			filter,
			bson.D{{"$set", updateObj}},
			opt,
		)
		if err != nil {
			msg := "Table was not updated"
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		c.JSON(http.StatusOK, result)
	}
}
