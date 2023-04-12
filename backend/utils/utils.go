package utils

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"golang-food-recipes/database"
	"golang-food-recipes/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Get user by email
func GetUserByEmail(email string) (models.User, error) {
	var user models.User
	collection := database.OpenCollection("users")
	err := collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

// Generate a random ID string
func GenerateId() string {
	id := primitive.NewObjectID()
	return id.Hex()
}

// Hash a password string using SHA-256 and base64 encoding
func HashPassword(password string) (string, error) {
	hash := sha256.Sum256([]byte(password))
	encodedHash := base64.URLEncoding.EncodeToString(hash[:])
	return encodedHash, nil
}

// Generate a JWT token for a user
func GenerateToken(user models.User) (string, error) {
	// Set token expiration time to 30 minutes from now
	expirationTime := time.Now().Add(30 * time.Minute).Unix()

	// Create claims for JWT token
	claims := jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"exp":   expirationTime,
	}

	// Sign the token with the secret key and claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("secretkey"))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
