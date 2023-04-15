package controller

import (
	"context"
	"golang-food-recipes/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SaveRecipe() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the user details from the context
		user, exists := c.Get("user")
		if !exists {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user details"})
			return
		}
		// Get the recipe id from the request
		id, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Retrieve the recipe data from the database
		var recipe models.Recipe
		if err := recipeCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&recipe); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve recipe details"})
			return
		}
		// Check if recipe already exists in user's saved recipes
		for _, savedRecipe := range user.(models.User).SavedRecipes {
			if savedRecipe.ID == recipe.ID {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Recipe already saved"})
				return
			}
		}
		// Create saved recipe
		var saved models.SavedRecipes
		saved.ID = primitive.NewObjectID()
		saved.RecipeID = id
		saved.UserID = user.(models.User).ID
		saved.Name = recipe.Name
		saved.Ingredients = recipe.Ingredients
		saved.Instructions = recipe.Instructions
		saved.TotalTime = recipe.TotalTime
		saved.Category = recipe.RecipeCategory
		saved.RecipeImage = recipe.RecipeImage

		// Update the user's saved recipes
		if _, err := userCollection.UpdateOne(context.Background(), bson.M{"_id": user.(models.User).ID}, bson.M{"$push": bson.M{"savedrecipes": saved}}); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Recipe saved successfully",
		})
	}
}
