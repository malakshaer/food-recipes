package main

import (
	"golang-food-recipes/middleware"
	"golang-food-recipes/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Use(middleware.Authentication())
	routes.AuthRoutes(router)
	router.Run("localhost:8080")
}
