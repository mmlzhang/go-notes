package main

import (
	"github.com/gin-gonic/gin"
	"go_learn/day-05/database"
	"go_learn/day-05/middleware"

	"go_learn/day-05/api"
)

func main() {
	app := gin.Default()
	db, _ := database.Initialize()
	app.Use(middleware.JWTMiddleware())
	app.Use(database.Inject(db))

	api.ApplyRoutes(app)
	app.Run()
}
