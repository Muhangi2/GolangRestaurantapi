package main

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"golang-Restaurantbooking/database"
	"golang-Restaurantbooking/middleware"
	"golang-Restaurantbooking/routes"
	"os"
)

var foodCollection *mongo.Collection=database.OpenCollection(database.client,"food")

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	routes.UserRoutes(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	router.invoiceRouter(router)
	routes.TableRoutes(router)

	router.Run(":" + port)

}
