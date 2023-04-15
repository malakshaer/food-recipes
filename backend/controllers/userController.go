package controller

import (
	"context"
	"encoding/base64"
	"golang-food-recipes/models"
	"golang-food-recipes/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the user details from the context
		user, exists := c.Get("user")
		if !exists {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user details"})
			return
		}
		// Get the user input
		var input models.User
		if err := c.BindJSON(&input); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Retrieve the user data from the database
		var dbUser models.User
		err := userCollection.FindOne(context.Background(), bson.M{"_id": user.(models.User).ID}).Decode(&dbUser)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user details"})
			return
		}
		// Check if username already exists
		if input.Username != "" {
			if _, err := userCollection.Find(context.Background(), bson.M{"username": input.Username}); err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
				return
			}
			dbUser.Username = input.Username
		}
		// Check if email already exists
		if input.Email != "" {
			if _, err := userCollection.Find(context.Background(), bson.M{"email": input.Email}); err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
				return
			}
			dbUser.Email = input.Email
		}
		// Check if password and confirm password are equal
		if input.Password != "" && input.ConfirmPassword != "" {
			// Hash password
			hashedPassword, err := utils.HashPassword(input.Password)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			if input.Password != input.ConfirmPassword {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Password and confirm password do not match"})
				return
			}
			dbUser.Password = hashedPassword
		}
		// Update profile image if provided
		if input.ProfileImage != "" {
			Image, err := base64.StdEncoding.DecodeString(input.ProfileImage)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Something went wrong"})
				return
			}
			dbUser.ProfileImage = string(Image)
		}
		// Update profile bio if provided
		if input.ProfileBio != "" {
			dbUser.ProfileBio = input.ProfileBio
		}
		// Update the user in database
		if _, err := userCollection.UpdateOne(context.Background(), bson.M{"_id": user.(models.User).ID}, bson.M{"$set": dbUser}); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// Return the updated user
		c.JSON(http.StatusOK, gin.H{"data": dbUser})
	}
}

func GetUserProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the user details from the context
		user, exists := c.Get("user")
		if !exists {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user details"})
			return
		}
		// Retrieve the user data from the database
		var dbUser models.User
		err := userCollection.FindOne(context.Background(), bson.M{"_id": user.(models.User).ID}).Decode(&dbUser)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user details"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": dbUser})
	}
}
