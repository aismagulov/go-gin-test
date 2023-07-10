package main

import (
	"aismagulov/go-gin-test/controller"
	"aismagulov/go-gin-test/middleware"
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	var secret string = os.Getenv("SESSION_SECRET")
	store := cookie.NewStore([]byte(secret))
	router.Use(sessions.Sessions("mysession", store))
	router.POST("/login", controller.Login)
	router.GET("/logout", controller.Logout)

	auth := router.Group("/auth")
	auth.Use(middleware.Authentication())
	auth.GET("/test", testEndpoint)

	router.Run("0.0.0.0:8080")
}

func testEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World!",
	})
}
