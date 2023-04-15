package routes

import (
	controller "golang-food-recipes/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.PUT("/user", controller.UpdateUser())
}
