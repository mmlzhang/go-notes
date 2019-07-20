package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-rest-gorm-rbac-sample/api/user"
)


// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		user.ApplyRoutes(api)
	}
}
