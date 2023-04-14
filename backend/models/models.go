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
	Recipes         []Recipe           `json:"Recipes,omitempty" json:"recipes,omitempty"`
}

type Recipe struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name            string             `bson:"Name,omitempty" json:"name,omitempty"`
	Ingredients     string             `bson:"Ingredients,omitempty" json:"ingredients,omitempty"`
	Instructions    string             `bson:"Instructions,omitempty" json:"instructions,omitempty"`
	TotalTime       string             `bson:"TotalTime,omitempty" json:"totaltime,omitempty"`
	RecipeCategory  string             `bson:"RecipeCategory,omitempty" json:"recipecategory,omitempty"`
	RecipeImage     string             `bson:"RecipeImage,omitempty" json:"recipeimage,omitempty"`
	RecipeAuthorID  primitive.ObjectID `bson:"RecipeAuthorID,omitempty" json:"recipeauthorid,omitempty"`
	RecipeCreatedAt string             `bson:"RecipeCreatedAt,omitempty" json:"recipecreatedat,omitempty"`
	RecipeUpdatedAt string             `bson:"RecipeUpdatedAt,omitempty" json:"recipeupdatedat,omitempty"`
}
