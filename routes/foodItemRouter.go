package routes
import (
	"github.com/gin-gonic/gin"
   controller	"golang-Restaurantbooking/controllers"
)
func FoodRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/food",controller.GetFoods())
	incomingRoutes.GET("/food/:food_id",controller.GetFood())
	incomingRoutes.POST("/food",controller.CreateFood())
	incomingRoutes.PATCH("/food/:food_id",controller.UpdateFood())
	incomingRoutes.DELETE("/food/:food_id",controller.DeleteFood())
}