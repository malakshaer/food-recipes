package routes

import (
	controller "golang-food-recipes/controllers"

	"github.com/gin-gonic/gin"
)

func RecipeRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/recipe", controller.CreateRecipe())

}
