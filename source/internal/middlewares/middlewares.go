package middlewares

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func AuthValidation() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("Authorization")
		if apiKey == "ApiKey "+os.Getenv("API_KEY") {
			c.Next()
		} else {
			c.JSON(http.StatusForbidden, nil)
		}
	}
}
