package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `bson:"_id"`
	Username        string             `json:"username"`
	Email           string             `json:"email"`
	Password        string             `json:"password"`
	ConfirmPassword string             `json:"confirm_password"`
	ProfileImage    string             `json:"profile_image"`
	ProfileBio      string             `json:"profile_bio"`
	Token           string             `json:"token"`
}

type Recipe struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Name           string             `json:"name" binding:"required"`
	Ingredients    string             `json:"ingredients" binding:"required"`
	Instructions   string             `json:"instructions" binding:"required"`
	TotalTime      string             `json:"total_time" binding:"required"`
	RecipeCategory string             `json:"recipe_category" binding:"required"`
	RecipeImage    string             `json:"recipe_image"`
	RecipeAuthor   primitive.ObjectID `json:"recipe_author_id"`
	RecipeDate     string             `json:"recipe_date"`
}
