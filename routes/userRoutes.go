package routes

import (
	"github.com/oscar503sv/gin-jwt/controllers"
	"github.com/oscar503sv/gin-jwt/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.POST("/register", controllers.RegisterUser)
	router.POST("/login", controllers.LoginUser)

	authorized := router.Group("/")
	authorized.Use(middlewares.JwtAuth())
	authorized.GET("/users", controllers.GetAllUsers)
	authorized.GET("/users/:id", controllers.GetUserById)
	authorized.PUT("/users/:id", controllers.UpdateUser)
	authorized.DELETE("/users/:id", controllers.DeleteUser)
	authorized.POST("/logout", controllers.LogoutUser)
}
