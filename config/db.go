package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"siddapp/models"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Dish{})

    return db
}

func GetDB() *gorm.DB {
    db := InitDB()
	if db == nil {
		panic("database not initilized")
	}

	return db
}
