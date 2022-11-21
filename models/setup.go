package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open((mysql.Open("artesiai_amogus:netorare10@tcp(103.28.53.179:3306)/artesiai_ikipiro")))
	if err != nil {
		panic((err))
	}

	// database.AutoMigrate()
	DB = database
}
