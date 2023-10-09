package config

import (
	"github.com/elysiamori/assignment3-09/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBConn() *gorm.DB {
	data := "root:@tcp(localhost:3306)/weather?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(data), &gorm.Config{})
	if err != nil {
		panic("Failed Connected Database" + err.Error())
	}

	db.AutoMigrate(&models.Weather{})

	return db
}
