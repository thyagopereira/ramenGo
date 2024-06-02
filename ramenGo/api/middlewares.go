package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(API_KEY string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header["x-api-key"][0] != API_KEY {
			c.JSON(http.StatusForbidden, gin.H{"error": "x-api-key header missing"})
		}
		c.Next()
	}
}
