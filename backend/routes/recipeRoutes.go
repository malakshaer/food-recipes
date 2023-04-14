package routes

import (
	controller "golang-food-recipes/controllers"

	"github.com/gin-gonic/gin"
)

func RecipeRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/recipe", controller.CreateRecipe())
	incomingRoutes.GET("/recipe", controller.GetAllRecipes())
	incomingRoutes.GET("/recipe/:id", controller.GetRecipeById())
	incomingRoutes.PUT("/recipe/:id", controller.UpdateRecipeById())
	incomingRoutes.DELETE("/recipe/:id", controller.DeleteRecipeById())
}
