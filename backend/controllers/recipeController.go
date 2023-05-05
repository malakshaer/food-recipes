package controller

import (
	"context"
	"encoding/base64"
	"golang-food-recipes/database"
	"golang-food-recipes/models"
	"golang-food-recipes/utils"
	"net/http"
	"strings"
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
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Failed to retrieve user details"})
			return
		}
		// Get recipe
		var input models.Recipe
		if err := c.ShouldBindJSON(&input); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}
		// Update recipe image if provided
		if input.RecipeImage != "" {
			// Remove the data url prefix
			encodedImage := strings.Replace(input.RecipeImage, "data:image/png;base64,", "", 1)
			// Decode the base64 image
			decodedImage, err := base64.StdEncoding.DecodeString(encodedImage)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			// Convert the decoded image to a base64 string
			base64Image := base64.StdEncoding.EncodeToString(decodedImage)
			input.RecipeImage = base64Image
		}
		// Check if name exists
		if input.Name == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Recipe name is required"})
			return
		}
		// Check if ingredients exists
		if input.Ingredients == nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Recipe ingredients are required"})
			return
		}
		// Create the recipe object
		recipe := models.Recipe{
			ID:                primitive.NewObjectID(),
			Name:              input.Name,
			Ingredients:       utils.GenerateIngredientIDs(input.Ingredients),
			Instructions:      input.Instructions,
			TotalTime:         input.TotalTime,
			RecipeCategory:    input.RecipeCategory,
			RecipeImage:       input.RecipeImage,
			RecipeAuthorID:    user.(models.User).ID,
			RecipeAuthorName:  user.(models.User).Username,
			RecipeAuthorImage: user.(models.User).ProfileImage,
			RecipeAuthorBio:   user.(models.User).ProfileBio,
			RecipeCreatedAt:   time.Now().Format(time.RFC3339),
		}
		// Store the recipe in the database
		if _, err := recipeCollection.InsertOne(context.Background(), recipe); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Failed to store recipe in the database"})
			return
		}
		var addedRecipe models.Recipe
		addedRecipe.ID = recipe.ID
		addedRecipe.Name = recipe.Name

		// Add the recipe to the user's recipes list in the database
		update := bson.M{"$push": bson.M{"recipes": addedRecipe}}
		if _, err := userCollection.UpdateOne(context.Background(), bson.M{"_id": user.(models.User).ID}, update); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Failed to update user's recipes list in the database"})
			return
		}
		// Return the recipe
		c.JSON(http.StatusCreated, gin.H{"recipe": recipe})
	}
}

func GetAllRecipes() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get all recipes from the database
		var recipes []models.Recipe
		cursor, err := recipeCollection.Find(context.Background(), bson.M{})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		defer cursor.Close(context.Background())
		for cursor.Next(context.Background()) {
			var recipe models.Recipe
			if err := cursor.Decode(&recipe); err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			recipes = append(recipes, recipe)
		}
		if err := cursor.Err(); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Return the recipe
		c.JSON(http.StatusOK, gin.H{"recipe": recipe})
	}
}

func UpdateRecipeById() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the recipe id from the request
		id, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Fail to get recipe id"})
			return
		}

		// Get the user details from the context
		user, exists := c.Get("user")
		if !exists {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Failed to retrieve user details"})
			return
		}

		// Get the updated recipe details from the request
		var input models.Recipe
		if err := c.BindJSON(&input); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}

		// Get the recipe from the database
		var recipe models.Recipe
		if err := recipeCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&recipe); err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Recipe not found"})
			return
		}

		// Check if the user is authorized to update the recipe
		if recipe.RecipeAuthorID != user.(models.User).ID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		// Update the recipe in the database
		update := bson.M{
			"$set": map[string]interface{}{
				"RecipeUpdatedAt": time.Now().Format(time.RFC3339),
			},
		}
		if input.Name != "" {
			update["$set"].(map[string]interface{})["Name"] = input.Name
			recipe.Name = input.Name
		}
		if input.Instructions != "" {
			update["$set"].(map[string]interface{})["Instructions"] = input.Instructions
			recipe.Instructions = input.Instructions
		}
		if input.Ingredients != nil {
			// Loop through the ingredients array and update the matching ingredient by ID
			for _, ingredient := range input.Ingredients {
				for j, recipeIngredient := range recipe.Ingredients {
					if recipeIngredient.ID == ingredient.ID {
						recipe.Ingredients[j].Text = ingredient.Text
						break
					}
				}
			}

			// Set the updated ingredients array in the update map
			update["$set"] = bson.M{"Ingredients": recipe.Ingredients}
		}
		if input.TotalTime != "" {
			update["$set"].(map[string]interface{})["TotalTime"] = input.TotalTime
			recipe.TotalTime = input.TotalTime
		}
		if input.RecipeCategory != "" {
			update["$set"].(map[string]interface{})["RecipeCategory"] = input.RecipeCategory
			recipe.RecipeCategory = input.RecipeCategory
		}
		if input.RecipeImage != "" {
			update["$set"].(map[string]interface{})["RecipeImage"] = input.RecipeImage
			recipe.RecipeImage = input.RecipeImage
		}

		if _, err := recipeCollection.UpdateOne(context.Background(), bson.M{"_id": id}, update); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Fail to update recipe in the database"})
			return
		}

		// Get the updated recipe from the database
		if err := recipeCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&recipe); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Fail to get updated recipe from database"})
			return
		}

		// Get the user from the database
		var dbUser models.User
		if err := userCollection.FindOne(context.Background(), bson.M{"_id": user.(models.User).ID}).Decode(&dbUser); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Fail to get user from database"})
			return
		}

		// Update the list of recipes of the user
		for i, r := range dbUser.Recipes {
			if r.ID == id {
				dbUser.Recipes[i] = recipe
				break
			}
		}

		// Update the user in the database
		if _, err := userCollection.UpdateOne(context.Background(), bson.M{"_id": user.(models.User).ID}, bson.M{"$set": bson.M{"recipes": dbUser.Recipes}}); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Fail to update user in the database"})
			return
		}

		// Return the updated recipe
		c.JSON(http.StatusOK, gin.H{"recipe": recipe})
	}
}

func DeleteRecipeById() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the recipe id from the request
		id, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Get the user details from the context
		user, exists := c.Get("user")
		if !exists {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Failed to retrieve user details"})
			return
		}

		// Get the recipe from the database
		var recipe models.Recipe
		if err := recipeCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&recipe); err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Recipe not found"})
			return
		}

		// Check if the user is authorized to delete the recipe
		if recipe.RecipeAuthorID != user.(models.User).ID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// Delete the recipe from the database
		if _, err := recipeCollection.DeleteOne(context.Background(), bson.M{"_id": id}); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Remove the recipe from the user's list of recipes in the database
		if _, err := userCollection.UpdateOne(context.Background(), bson.M{"_id": user.(models.User).ID}, bson.M{"$pull": bson.M{"recipes": bson.M{"_id": id}}}); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Remove the recipe from the saved list of all users saved the recipe
		if _, err := userCollection.UpdateMany(context.Background(), bson.M{"savedRecipes": bson.M{"$elemMatch": bson.M{"_id": id}}}, bson.M{"$pull": bson.M{"savedRecipes": bson.M{"_id": id}}}); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Return success message
		c.JSON(http.StatusOK, gin.H{"message": "Recipe deleted successfully"})
	}
}
