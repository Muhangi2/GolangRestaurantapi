package routes
import (
	"github.com/gin-gonic/gin"
   controller	"golang-Restaurantbooking/controllers"
)
 func OrderItemRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/orderItem",controller.GetOrderItems())
	incomingRoutes.GET("/orderItem/:orderItem_id",controller.GetOrderItem())
	incomingRoutes.POST("/orderItem",controller.CreateOrderItem())
	incomingRoutes.PATCH("/orderItem/:orderItem_id",controller.UpdateOrderItem())
	incomingRoutes.DELETE("/orderItem/:orderItem_id",controller.DeleteOrderItem())
 }