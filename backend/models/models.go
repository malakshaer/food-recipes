package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username        string             `json:"username,omitempty" json:"UserName,omitempty"`
	Email           string             `json:"email,omitempty" json:"Email,omitempty"`
	Password        string             `json:"password"`
	ConfirmPassword string             `json:"confirm_password"`
	ProfileImage    string             `json:"profileimage,omitempty" json:"ProfileImage,omitempty"`
	ProfileBio      string             `json:"profilebio,omitempty" json:"ProfileBio,omitempty"`
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
