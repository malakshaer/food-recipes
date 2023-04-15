package routes

import (
	controller "golang-food-recipes/controllers"

	"github.com/gin-gonic/gin"
)

func ActionRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("save_recipe/:id", controller.SaveRecipe())
	incomingRoutes.POST("unsave_recipe/:id", controller.UnSaveRecipe())
}
