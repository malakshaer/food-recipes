package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username        string             `json:"username,omitempty" json:"UserName,omitempty"`
	Email           string             `json:"email,omitempty" json:"Email,omitempty"`
	Password        string             `json:"password"`
	ConfirmPassword string             `json:"confirm_password" bson:"confirm_password,omitempty"`
	CurrentPassword string             `json:"current_password" bson:"current_password,omitempty"`
	ProfileImage    string             `json:"profileimage,omitempty" json:"ProfileImage,omitempty"`
	ProfileBio      string             `json:"profilebio,omitempty" json:"ProfileBio,omitempty"`
	Token           string             `json:"token"`
	Recipes         []Recipe           `json:"Recipes,omitempty" json:"recipes,omitempty"`
	SavedRecipes    []SavedRecipes     `json:"SavedRecipes,omitempty" json:"savedrecipes,omitempty"`
	LikedRecipes    []LikedRecipes     `json:"LikedRecipes,omitempty" json:"likedrecipes,omitempty"`
}

type Recipe struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name              string             `bson:"Name,omitempty" json:"name,omitempty"`
	Ingredients       []Ingredients      `bson:"Ingredients,omitempty" json:"ingredients,omitempty"`
	Instructions      string             `bson:"Instructions,omitempty" json:"instructions,omitempty"`
	TotalTime         string             `bson:"TotalTime,omitempty" json:"totaltime,omitempty"`
	RecipeCategory    string             `bson:"RecipeCategory,omitempty" json:"recipecategory,omitempty"`
	RecipeImage       string             `bson:"RecipeImage,omitempty" json:"recipeimage,omitempty"`
	RecipeAuthorID    primitive.ObjectID `bson:"RecipeAuthorID,omitempty" json:"recipeauthorid,omitempty"`
	RecipeAuthorName  string             `bson:"RecipeAuthorName,omitempty" json:"recipeauthorname,omitempty"`
	RecipeAuthorImage string             `bson:"RecipeAuthorImage,omitempty" json:"recipeauthorimage,omitempty"`
	RecipeAuthorBio   string             `bson:"RecipeAuthorBio,omitempty" json:"recipeauthorbio,omitempty"`
	RecipeCreatedAt   string             `bson:"RecipeCreatedAt,omitempty" json:"recipecreatedat,omitempty"`
	RecipeUpdatedAt   string             `bson:"RecipeUpdatedAt,omitempty" json:"recipeupdatedat,omitempty"`
	Likes             int32              `bson:"Likes,omitempty" json:"likes,omitempty"`
}

type Ingredients struct {
	ID   primitive.ObjectID `bson:"ingredientid,omitempty" json:"ingredientid,omitempty"`
	Text string             `bson:"Text,omitempty" json:"text,omitempty"`
}

type SavedRecipes struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	RecipeID          primitive.ObjectID `bson:"RecipeID,omitempty" json:"recipeid,omitempty"`
	RecipeAuthorID    primitive.ObjectID `bson:"RecipeAuthorID,omitempty" json:"recipeauthorid,omitempty"`
	Name              string             `bson:"Name,omitempty" json:"name,omitempty"`
	Ingredients       []Ingredients      `bson:"Ingredients,omitempty" json:"ingredients,omitempty"`
	Instructions      string             `bson:"Instructions,omitempty" json:"instructions,omitempty"`
	TotalTime         string             `bson:"TotalTime,omitempty" json:"totaltime,omitempty"`
	Category          string             `bson:"Category,omitempty" json:"category,omitempty"`
	RecipeImage       string             `bson:"RecipeImage,omitempty" json:"recipeimage,omitempty"`
	RecipeAuthorName  string             `bson:"RecipeAuthorName,omitempty" json:"recipeauthorname,omitempty"`
	RecipeAuthorImage string             `bson:"RecipeAuthorImage,omitempty" json:"recipeauthorimage,omitempty"`
	RecipeAuthorBio   string             `bson:"RecipeAuthorBio,omitempty" json:"recipeauthorbio,omitempty"`
}

type LikedRecipes struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	RecipeID primitive.ObjectID `bson:"RecipeID,omitempty" json:"recipeid,omitempty"`
	UserID   primitive.ObjectID `bson:"UserID,omitempty" json:"userid,omitempty"`
}
