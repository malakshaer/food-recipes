package models

import "go.mongodb.org/mongo-driver/bson/primitive"

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

// type Recipe struct {
// 	ID             int64  `json:"id"`
// 	Name           string `json:"name"`
// 	Ingredients    string `json:"ingredients"`
// 	Instructions   string `json:"instructions"`
// 	PrepTime       string `json:"prep_time"`
// 	CookTime       string `json:"cook_time"`
// 	TotalTime      string `json:"total_time"`
// 	RecipeCategory string `json:"recipe_category"`
// 	RecipeImage    string `json:"recipe_image"`
// 	RecipeAuthor   string `json:"recipe_author_id"`
// 	RecipeDate     string `json:"recipe_date"`
// }

// type Like struct {
// 	ID       int64  `json:"id"`
// 	RecipeID string `json:"recipe_id"`
// 	UserID   string `json:"user_id"`
// }

// type SavedRecipe struct {
// 	ID       int64  `json:"id"`
// 	RecipeID string `json:"recipe_id"`
// 	UserID   string `json:"user_id"`
// }
