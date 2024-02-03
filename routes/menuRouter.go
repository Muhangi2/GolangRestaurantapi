package routes
import (
	"github.com/gin-gonic/gin"
   controller	"golang-Restaurantbooking/controllers"
)
func MenuRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/menu",controller.GetMenus())
	incomingRoutes.GET("/menu/:menu_id",controller.GetMenu())
	incomingRoutes.POST("/menu",controller.CreateMenu())
	incomingRoutes.PATCH("/menu/:menu_id",controller.UpdateMenu())
	incomingRoutes.DELETE("/menu/:menu_id",controller.DeleteMenu())
}