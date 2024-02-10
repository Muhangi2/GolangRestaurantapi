package routes

import (
	controller "golang-Restaurantbooking/controllers"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/order", controller.GetOrders())
	incomingRoutes.GET("/order/:order_id", controller.GetOrder())
	incomingRoutes.POST("/order", controller.CreateOrder())
	incomingRoutes.PATCH("/order/:order_id", controller.UpdateOrder())
	// incomingRoutes.DELETE("/order/:order_id", controller.DeleteOrder())
}
