package routes

import (
	controller "golang-food-recipes/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.PUT("/user", controller.UpdateUser())
	incomingRoutes.GET("/user", controller.GetUserProfile())
	incomingRoutes.GET("/users", controller.GetAllUsers())
	incomingRoutes.GET("/user/:id", controller.GetUserById())
}
