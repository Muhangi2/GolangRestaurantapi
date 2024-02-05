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

var menuCollection *mongo.Collection = database.OpenCollection(database.Client, "menu")

func GetMenus() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		result, err := menuCollection.Find(context.TODO(), bson.M{})
		defer cancel()
		if err != nil {
			c.JSON(500, gin.H{"error": "Error while fetching the menu"})
		}
		var allMenus []bson.M
		if err = result.All(ctx, &allMenus); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allMenus)

	}

}
func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		menuId := c.Param("menu_id")
		var menu models.Menu
		err := menuCollection.FindOne(ctx, bson.M{"menu_id": menuId}).Decode(&menu)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error occured while fetching the menu or menu does not exist",
			})

		}
		c.JSON(http.StatusOK, menu)
	}
}
func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout((context.Background()), 100*time.Second)
		var menu models.Menu

		if err := c.BindJSON(&menu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
			return
		}
		validateError := validate.Struct(menu)
		if validateError != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validateError.Error()})
			return
		}

		//preparing to add items..
		menu.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		menu.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		menu.ID = primitive.NewObjectID()
		menu.Menu_id = menu.ID.Hex()

		result, issertErr := menuCollection.InsertOne(ctx, menu)
		if issertErr != nil {
			msg := fmt.Sprintf("Food item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, result)

	}
}
func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		var menu models.Menu
		if err := c.BindJSON(&menu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
			return
		}
		menuId := c.Param("menu_id")
		filter := bson.M{"menu_id": menuId}

		var updateObj primitive.D
		if !menu.Start_date.IsZero() && !menu.End_date.IsZero() {
			if !inTimeSpan(menu.Start_date, menu.End_date, time.Now()) {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Start date and end date must be in the future"})
				defer cancel()
				return
			}
			updateObj = append(updateObj, bson.E{"start_date", menu.Start_date})
			updateObj = append(updateObj, bson.E{"end_date", menu.End_date})

			//menu name
			if menu.Name != nil {
				updateObj = append(updateObj, bson.E{"name", menu.Name})
			}
			//menu-item
			if menu.Category != nil {
				updateObj = append(updateObj, bson.E{"category", menu.Category})
			}
			//since its update funcion we consider update and update it
			menu.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			updateObj = append(updateObj, bson.E{"updated_at", menu.Updated_at})

			upsert := true

			opt := options.UpdateOptions{
				Upsert: &upsert,
			}
			result, err := menuCollection.UpdateOne(
				ctx,
				filter,
				bson.D{
					{"$set", updateObj},
				},
				&opt,
			)
			//check for the error
			if err != nil {
				msg := fmt.Sprintf("Menu was not updated")
				c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
				defer cancel()
				return

			}
			defer cancel()
			c.JSON(http.StatusOK, result)

		}

	}
}

func DeleteMenu() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

//function for checking the the end_date is after start_date

func inTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}
