package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Dish struct {
	gorm.Model
	ID          uint
	Name        string
	Description string
	Thumbnail   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Dish{})

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		var dishes []Dish
		db.Find(&dishes)

		c.HTML(http.StatusOK, "index.tmpl", gin.H{
            "Dishes": dishes,
		})
	})

	router.Run()
}
