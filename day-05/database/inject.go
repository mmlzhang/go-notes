package database

import (
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
)

func Inject(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}
