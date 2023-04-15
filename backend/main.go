package main

import (
	"golang-food-recipes/middleware"
	"golang-food-recipes/routes"

	"github.com/gin-gonic/gin"
)

// var host = os.Getenv("DB_HOST")
// var port = os.Getenv("DB_PORT")

func main() {

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	routes.AuthRoutes(router)

	router.Use(middleware.Authentication())
	routes.UserRoutes(router)
	routes.RecipeRoutes(router)
	routes.ActionRoutes(router)
	// router.Run(host + ":" + port)
	router.Run("localhost:8080")
}
