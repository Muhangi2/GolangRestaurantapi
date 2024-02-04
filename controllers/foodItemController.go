package controllers

import (
	"context"
	"fmt"
	"golang-Restaurantbooking/database"
	"golang-Restaurantbooking/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//defining food collection

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

// importing validate package
var validate = validator.New()

func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		foodId := c.Param("food_id")
		var food models.FoodItem
		err := foodCollection.FindOne(ctx, bson.M{"food_id": foodId}).Decode(&food)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error occured while fetching the food or food does not exist",
			})

		}
		c.JSON(http.StatusOK, food)
	}
}
func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout((context.Background()), 100*time.Second)
		var food models.FoodItem
		var menu models.Menu

		if err := c.BindJSON(&food); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
			return
		}
		validateError := validate.Struct(food)
		if validateError != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validateError.Error()})
			return
		}
		//checking if the menu exists
		err := foodCollection.FindOne(ctx, bson.M{"menu_id": food.Menu_id}).Decode(&menu)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Menu does not exist"})
			return
		}
		//preparing to add items..
		food.CreatedAt = time.Now()
		food.UpdatedAt = time.Now()
		food.ID = primitive.NewObjectID()
		food.Food_id = food.ID.Hex()

		var num = toFixed(*food.Price, 2)
		food.Price = &num

		result, issertErr := foodCollection.InsertOne(ctx, food)
		if issertErr != nil {
			msg := fmt.Sprintf("Food item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, result)

	}
}
func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
func DeleteFood() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// other functions
func round(num float64) int {
}
func toFixed(num float64, precision int) float64 {

}
