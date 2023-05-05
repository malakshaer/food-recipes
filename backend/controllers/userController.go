package controller

import (
	"context"
	"encoding/base64"
	"golang-food-recipes/models"
	"golang-food-recipes/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		// current password must exists
		if input.CurrentPassword == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Please enter your current password"})
			return
		}
		// Check if current password is correct
		if input.CurrentPassword != "" {
			// Check if current password matches the password in the database
			if !utils.CheckPasswordHash(input.CurrentPassword, dbUser.Password) {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Current password is incorrect"})
				return
			}
		}

		// Check if password and confirm password are equal
		if input.Password != "" && input.ConfirmPassword != "" {
			// Check if current password was entered
			if input.CurrentPassword == "" {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Please enter your current password"})
				return
			}

			// Check if new password and confirm password match
			if input.Password != input.ConfirmPassword {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Password and confirm password do not match"})
				return
			}

			// Hash new password
			hashedPassword, err := utils.HashPassword(input.Password)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			// Update password in database
			dbUser.Password = hashedPassword
		}

		// Update profile image if provided
		if input.ProfileImage != "" {
			// Remove the data url prefix
			encodedImage := strings.Replace(input.ProfileImage, "data:image/png;base64,", "", 1)
			// Decode the base64 image
			decodedImage, err := base64.StdEncoding.DecodeString(encodedImage)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			// Convert the decoded image to a base64 string
			base64Image := base64.StdEncoding.EncodeToString(decodedImage)
			dbUser.ProfileImage = base64Image
		}

		// Update profile bio if provided
		if input.ProfileBio != "" {
			dbUser.ProfileBio = input.ProfileBio
		}
		// Store only username, password,bio and image of dbUser
		dbUser = models.User{
			ID:           dbUser.ID,
			Username:     dbUser.Username,
			Email:        dbUser.Email,
			Password:     dbUser.Password,
			ProfileBio:   dbUser.ProfileBio,
			ProfileImage: dbUser.ProfileImage,
			Token:        dbUser.Token,
			Recipes:      dbUser.Recipes,
			SavedRecipes: dbUser.SavedRecipes,
			LikedRecipes: dbUser.LikedRecipes,
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

func GetAllUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve all users from the database
		cursor, err := userCollection.Find(context.Background(), bson.M{})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer cursor.Close(context.Background())
		var users []models.User
		for cursor.Next(context.Background()) {
			var user models.User
			cursor.Decode(&user)
			users = append(users, user)
		}
		c.JSON(http.StatusOK, gin.H{"data": users})
	}
}

func GetUserById() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the user id from the request
		id, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Retrieve the user data from the database
		var dbUser models.User
		if err := userCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&dbUser); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user details"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": dbUser})
	}
}
