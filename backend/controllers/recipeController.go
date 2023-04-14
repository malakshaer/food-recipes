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

func GetAllRecipes() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get all recipes from the database
		var recipes []models.Recipe
		cursor, err := recipeCollection.Find(context.Background(), bson.M{})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer cursor.Close(context.Background())
		for cursor.Next(context.Background()) {
			var recipe models.Recipe
			if err := cursor.Decode(&recipe); err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			recipes = append(recipes, recipe)
		}
		if err := cursor.Err(); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Return all recipes
		c.JSON(http.StatusOK, gin.H{"recipes": recipes})
	}
}

func GetRecipeById() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the recipe ID from the request
		id, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Get the recipe from the database
		var recipe models.Recipe
		if err := recipeCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&recipe); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Return the recipe
		c.JSON(http.StatusOK, gin.H{"recipe": recipe})
	}
}
