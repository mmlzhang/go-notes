package database

import (
	"fmt"
	"go_learn/day-05/database/models"
)

func Initialize() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:admin@/test?charset=utf8&parseTime=True&loc=Local")
	db.LogMode(true)
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to database")
	models.Migrate(db)
	return db, err

}