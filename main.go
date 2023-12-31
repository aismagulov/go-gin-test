package main

import (
	"aismagulov/go-gin-test/controller"
	"aismagulov/go-gin-test/middleware"
	"embed"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

//go:embed templates/*
var f embed.FS

func main() {
	router := gin.Default()
	templ := template.Must(template.New("").ParseFS(f, "templates/*.html"))
	router.SetHTMLTemplate(templ)

	router.StaticFS("/public", http.FS(f))

	// router.LoadHTMLGlob("templates/*.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	var secret string = os.Getenv("SESSION_SECRET")
	var redispass string = os.Getenv("REDIS_PASSWORD")

	log.Println("secret " + secret)
	log.Println("redispass " + redispass)

	// store := cookie.NewStore([]byte(secret))
	log.Println("Connecting to redis")
	store, error := redis.NewStore(10, "tcp", "localhost:6379", redispass, []byte(secret))
	//store, error := redis.NewStore(10, "tcp", "go-gin-test_redis-session_1:6379", "", []byte(secret))
	if error != nil {
		log.Println("Cannot connect to redis")
		panic(error)
	}

	store.Options(sessions.Options{MaxAge: 30})

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
