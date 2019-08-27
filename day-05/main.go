package main

import (
	"github.com/gin-gonic/gin"
	"go-notes/day-05/database"
	"go-notes/day-05/middleware"

	"go-notes/day-05/api"
)

func main() {
	app := gin.Default()
	db, _ := database.Initialize()
	app.Use(middleware.JWTMiddleware())
	app.Use(database.Inject(db))

	api.ApplyRoutes(app)
	app.Run()
}
