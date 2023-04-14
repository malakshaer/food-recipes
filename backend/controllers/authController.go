package controller

import (
	"context"
	"fmt"
	"golang-food-recipes/database"
	"golang-food-recipes/models"
	"golang-food-recipes/utils"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = database.OpenCollection("users")
var secretKey string = os.Getenv("SECRET_KEY")

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

		// Check if password and confirm password are equal
		if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(user.ConfirmPassword)); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Password and confirm password must be equal"})
			return
		}

		// Remove the ConfirmPassword field from the user struct
		user.ConfirmPassword = ""

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
			"message": fmt.Sprintf("%s's account created successfully", user.Username),
			"token":   token,
		})
	}
}

func Login() gin.HandlerFunc {
	store := sessions.NewCookieStore([]byte(secretKey))
	return func(c *gin.Context) {
		// Check if user is already logged in
		session, err := store.Get(c.Request, "session-name")
		if err == nil {
			userID := session.Values["user_id"]
			if userID != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "User is already logged in"})
				return
			}
		}

		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Check if the email exists in the database
		userFound, err := utils.GetUserByEmail(user.Email)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid Email or password"})
			return
		}

		// Check if password is correct
		if err := bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(user.Password)); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid Email or password"})
			return
		}

		// Generate a JWT token for the user
		token, err := utils.GenerateToken(user)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		userFound.Token = token

		// Update the user's token in the database
		update := bson.M{"$set": bson.M{"token": token}}
		if _, err := userCollection.UpdateOne(context.Background(), bson.M{"_id": userFound.ID}, update); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Store the user's ID in the session
		session.Values["user_id"] = userFound.ID.Hex()
		if err := session.Save(c.Request, c.Writer); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Return success response with user ID and JWT token
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%s successfully logged in", userFound.Username),
			"token":   token,
		})

	}
}

func Logout() gin.HandlerFunc {
	store := sessions.NewCookieStore([]byte(secretKey))
	return func(c *gin.Context) {
		// Get the session for the user
		session, err := store.Get(c.Request, "session-name")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Clear the session and save changes
		session.Options.MaxAge = -1
		if err := session.Save(c.Request, c.Writer); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Return success response
		c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
	}
}
