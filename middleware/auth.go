package middleware

import (
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionID := session.Get("id")

		log.Println("session: ", sessionID.(string))
		if sessionID == nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "unauthorized",
			})
			c.Abort()
		}
	}
}
