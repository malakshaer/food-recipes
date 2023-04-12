package routes

import (
	controller "golang-food-recipes/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/register", controller.Register())

}
