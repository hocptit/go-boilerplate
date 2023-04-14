package auth

import (
	_ "fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}
func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

func NoAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
