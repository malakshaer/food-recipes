package middleware

import (
	"context"
	"fmt"
	"golang-food-recipes/database"
	"golang-food-recipes/models"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection("users")
var secretKey string = os.Getenv("SECRET_KEY")

func GetUserByID(id string) (models.User, error) {
	var user models.User

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, err
	}

	filter := bson.M{"_id": objID}
	err = userCollection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

type Claims struct {
	UserID string `json:"userID"`
	jwt.StandardClaims
}

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Authorization header from the request
		tokenString := c.GetHeader("Authorization")

		// Check if the token is present in the header
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			return
		}

		// Extract the JWT token from the Authorization header
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		// Parse and validate the JWT token using the secret key
		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			// Make sure that the signing method is HMAC
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			// Use the secret key to verify the token
			return []byte(secretKey), nil
		})

		// Check for errors during JWT verification
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		// Check if the token is valid
		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		// Extract the user ID from the JWT token claims
		claims, ok := token.Claims.(*Claims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			return
		}

		// Retrieve the user object from the database
		user, err := GetUserByID(claims.UserID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		// Add the user object to the context
		c.Set("user", user)

		// Call the next handler
		c.Next()
	}
}
