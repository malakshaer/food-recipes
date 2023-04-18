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

func UnSaveRecipe() gin.HandlerFunc {
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

		// Check if recipe exists in user's saved recipes
		found := false
		for _, savedRecipe := range user.(models.User).SavedRecipes {
			if savedRecipe.RecipeID == id {
				found = true
				break
			}
		}

		if !found {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Recipe not found in saved recipes"})
			return
		}

		// Remove the recipe from the user's saved recipes
		if _, err := userCollection.UpdateOne(context.Background(), bson.M{"_id": user.(models.User).ID}, bson.M{"$pull": bson.M{"savedrecipes": bson.M{"RecipeID": id}}}); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Recipe unsaved successfully",
		})
	}
}

func LikeRecipe() gin.HandlerFunc {
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

		// Check if the user already liked the recipe
		for _, likedRecipe := range user.(models.User).LikedRecipes {
			if likedRecipe.RecipeID == recipe.ID {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Recipe already liked"})
				return
			}
		}

		// Create liked recipe
		var liked models.LikedRecipes
		liked.ID = primitive.NewObjectID()
		liked.RecipeID = id
		liked.UserID = user.(models.User).ID

		// Update the user's liked recipes
		if _, err := userCollection.UpdateOne(context.Background(), bson.M{"_id": user.(models.User).ID}, bson.M{"$push": bson.M{"likedrecipes": liked}}); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Update the recipe's likes
		if _, err := recipeCollection.UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$inc": bson.M{"Likes": 1}}); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Recipe liked successfully",
		})
	}
}

func UnLikeRecipe() gin.HandlerFunc {
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

		// Check if recipe exists in user's liked recipes
		found := false
		for _, likedRecipe := range user.(models.User).LikedRecipes {
			if likedRecipe.RecipeID == id {
				found = true
				break
			}
		}

		if !found {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Recipe not found in liked recipes"})
			return
		}

		// Remove the recipe from the user's liked recipes
		if _, err := userCollection.UpdateOne(context.Background(), bson.M{"_id": user.(models.User).ID}, bson.M{"$pull": bson.M{"likedrecipes": bson.M{"RecipeID": id}}}); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Update the recipe's likes
		if _, err := recipeCollection.UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$inc": bson.M{"likes": -1}}); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Recipe unLiked successfully",
		})

	}
}

func SearchRecipe() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the user details from the context
		_, exists := c.Get("user")
		if !exists {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user details"})
			return
		}

		// Get the search query from the request
		query := c.Query("q")

		// Get the search results
		var results []models.Recipe

		// Search by category or name
		cursor, err := recipeCollection.Find(context.Background(), bson.M{"$or": []bson.M{
			{"Name": bson.M{"$regex": query, "$options": "i"}},
			{"RecipeCategory": bson.M{"$regex": query, "$options": "i"}},
		}})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		var categoryResults []models.Recipe
		if err := cursor.All(context.Background(), &categoryResults); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		results = append(results, categoryResults...)

		// Remove duplicates
		uniqueResults := make([]models.Recipe, 0, len(results))
		idMap := make(map[primitive.ObjectID]bool)
		for _, recipe := range results {
			if _, ok := idMap[recipe.ID]; !ok {
				uniqueResults = append(uniqueResults, recipe)
				idMap[recipe.ID] = true
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"results": uniqueResults,
		})
	}
}
