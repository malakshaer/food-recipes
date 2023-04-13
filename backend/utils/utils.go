package utils

import (
	"context"
	"errors"
	"golang-food-recipes/database"
	"golang-food-recipes/models"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/sessions"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// "github.com/gin-contrib/sessions"

var store sessions.Store

func init() {
	store = sessions.NewCookieStore([]byte("secret-key"))
}

var userCollection *mongo.Collection = database.OpenCollection("users")

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

// Hash a password string using SHA-256 and base64 encoding
func HashPassword(password string) (string, error) {
	// Generate a bcrypt hash of the password
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	// Convert the hashed password to a string and return it
	hashedPassword := string(hashed)
	return hashedPassword, nil
}

// Generate a JWT token for a user by ID
func GenerateToken(user models.User) (string, error) {
	// Set token expiration time
	expirationTime := time.Now().Add(30000 * time.Minute).Unix()

	// Create claims for JWT token
	claims := jwt.MapClaims{
		"id":  user.ID,
		"exp": expirationTime,
	}

	// Sign the token with the secret key and claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("secret-key"))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// Check if user is logged in
func GetUserIDFromSession(r *http.Request, store *sessions.CookieStore) (string, error) {
	// Get the session from the cookie store.
	session, err := store.Get(r, "session-name")
	if err != nil {
		return "", err
	}

	// Get the user ID from the session.
	userID, ok := session.Values["userID"].(string)
	if !ok || userID == "" {
		return "", errors.New("invalid userID in session")
	}

	return userID, nil
}
