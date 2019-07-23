package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authorized(c *gin.Context) {
	_, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "msg": "must login"})
		c.Abort()
	}
}