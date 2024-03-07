package routes
import (
	"github.com/gin-gonic/gin"
	controller "golang-Restaurantbooking/controllers"
)
func TableRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/table",controller.GetTables())
	incomingRoutes.GET("/table/:table_id",controller.GetTable())
	incomingRoutes.POST("/table",controller.CreateTable())
	incomingRoutes.PATCH("/table/:table_id",controller.UpdateTable())
	// incomingRoutes.DELETE("/table/:table_id",controller.DeleteTable())
}