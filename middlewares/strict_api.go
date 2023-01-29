package middlewares

import (
	"net/http"
	"os"

	"github.com/adityarizkyramadhan/freepass-2023/utils/response"
	"github.com/gin-gonic/gin"
)

func SecureAPI() gin.HandlerFunc {
	return extractAuth
}

func extractAuth(c *gin.Context) {
	username := c.Request.Header.Get("X-API-Username")
	password := c.Request.Header.Get("X-API-Password")
	if username == os.Getenv("API_USERNAME") && password == os.Getenv("API_PASSWORD") {
		c.Next()
		return
	} else {
		c.JSON(http.StatusUnauthorized, response.ResponseWhenFail("unauthorized user"))
		c.Abort()
	}
}
