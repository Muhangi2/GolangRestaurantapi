package main

import (
	"golang-Restaurantbooking/middlewares"
	"golang-Restaurantbooking/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	//here i checking if the port is null or not
	if port == "" {
		port = "8080"
	}
	//mking an instance of gin
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(middlewares.Authentication())

	//routers for different path
	routes.UserRoutes(router)
	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)
	routes.TableRoutes(router)
	router.Run(":" + port)

}
