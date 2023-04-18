package main

import (
	"golang-food-recipes/middleware"
	"golang-food-recipes/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowHeaders = []string{"Authorization", "Content-Type"}
	router.Use(cors.New(config))

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	routes.AuthRoutes(router)

	router.Use(middleware.Authentication())
	routes.UserRoutes(router)
	routes.RecipeRoutes(router)
	routes.ActionRoutes(router)

	router.Run("localhost:8080")
}
