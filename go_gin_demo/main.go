package main

import (
	"github.com/gin-gonic/gin"
	"go_learn/go_gin_demo/database"
	"go_learn/go_gin_demo/middleware"

	"go_learn/go_gin_demo/api"
)

func main() {
	app := gin.Default()
	db, _ := database.Initialize()
	app.Use(middleware.JWTMiddleware())
	app.Use(database.Inject(db))

	api.ApplyRoutes(app)
	app.Run()
}
