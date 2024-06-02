package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(API_KEY string) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header["X-Api-Key"]
		if len(header) == 0 || header[0] != API_KEY {
			c.JSON(http.StatusForbidden, gin.H{"error": "x-api-key header missing"})
			c.Abort()
		} else {
			c.Next()
			return
		}

	}
}
