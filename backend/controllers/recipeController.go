package controller

import (
	"context"
	"golang-food-recipes/database"
	"golang-food-recipes/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var recipeCollection *mongo.Collection = database.OpenCollection("recipes")

func CreateRecipe() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the user details from the context
		user, exists := c.Get("user")
		if !exists {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to to retrieve user details"})
			return
		}

		// Get recipe
		var input models.Recipe
		if err := c.BindJSON(&input); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Create the recipe object
		recipe := models.Recipe{
			ID:             primitive.NewObjectID(),
			Name:           input.Name,
			Ingredients:    input.Ingredients,
			Instructions:   input.Instructions,
			TotalTime:      input.TotalTime,
			RecipeCategory: input.RecipeCategory,
			RecipeImage:    input.RecipeImage,
			RecipeAuthor:   user.(models.User).ID,
			RecipeDate:     time.Now().Format(time.RFC3339),
		}

		// Store the recipe in the database
		if _, err := recipeCollection.InsertOne(context.Background(), recipe); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Add the recipe to the user's recipes list in the database
		update := bson.M{"$push": bson.M{"recipes": recipe}}
		if _, err := userCollection.UpdateOne(context.Background(), bson.M{"_id": user.(models.User).ID}, update); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Return the recipe
		c.JSON(http.StatusOK, gin.H{"recipe": recipe})
	}
}
