package user

import (
	"github.com/gin-rest-gorm-rbac-sample/database/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"github.com/gin-rest-gorm-rbac-sample/utils"
	"github.com/gin-rest-gorm-rbac-sample/lib/common"
)

type User = models.User


func getUser(c *gin.Context) {
	userID := c.Param("id")
	db := c.MustGet("db").(*gorm.DB)

	var user User

	db.First(&user, userID)
	if user.Id == o {
		c.Json(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}
	c.JSON(200, user.Serialize())
}


func create(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	type RequestBody struct {
		Email string `json:"email" binding:"required"`
		Name string `json:"name" binding:"required"`
		Password string `json:"password" binding:"required"`
		Age int `json:"age" binding:"required"`
	}

	var body RequestBody
	if err := c.BindJSON(&body); err != nil {
		fmt.Println(body)
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "miss!"})
		return
	}

	var exists User
	if err := db.Where("email = ?", body.Email).First(&exists).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"status": http.StatusConflict, "message": "email already exists!"})
		return
	}

	hash, hashErr := utils.HashPassword(body, Password)
	if hashErr != nil {
		c.Json(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "email existancy"})
	}

	// create user
	user := User{
		Name: body.Name,
		Email: body.Email,
		Password: hash,
		Age: body.Age,

	}

	db.NewRecord(user)
	db.Create(&user)

	c.JSON(200, common.JSON{
		"user": user.Serialize(),
	})
}