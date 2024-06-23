package main

import (
	"github.com/oscar503sv/gin-jwt/config"
	"github.com/oscar503sv/gin-jwt/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	config.ConnectDB()
	routes.UserRoutes(router)
	router.Run()
}
