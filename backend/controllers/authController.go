package controller

import (
	"context"
	"golang-food-recipes/database"
	"golang-food-recipes/models"
	"golang-food-recipes/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection("users")

func Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Check if the email already exists in the database
		_, err := utils.GetUserByEmail(user.Email)
		if err == nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "User with this email already exists"})
			return
		}

		// Hash the user's password before storing it in the database
		hashedPassword, err := utils.HashPassword(user.Password)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		user.Password = hashedPassword

		// Hash the user's confirm password before storing it in the database
		hashedConfirmPassword, err := utils.HashPassword(user.ConfirmPassword)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		user.ConfirmPassword = hashedConfirmPassword

		// Generate an ID for the user
		user.ID = primitive.NewObjectID()

		// Store the user in the database
		if _, err := userCollection.InsertOne(context.Background(), user); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Generate a JWT token for the user
		token, err := utils.GenerateToken(user)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		user.Token = token

		// Update the user's token in the database
		update := bson.M{"$set": bson.M{"token": token}}
		if _, err := userCollection.UpdateOne(context.Background(), bson.M{"_id": user.ID}, update); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Return success response with user ID and JWT token
		c.JSON(http.StatusOK, gin.H{
			"id":    user.ID.Hex(),
			"token": token,
		})
	}
}

func Hello() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello Me"})
	}
}
