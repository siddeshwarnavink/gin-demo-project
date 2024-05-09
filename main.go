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

	router.GET("/create", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "create.tmpl", gin.H{})
	})

	router.POST("/create", func(ctx *gin.Context) {
		name := ctx.PostForm("name")
		description := ctx.PostForm("description")
		thumbnail := ctx.PostForm("thumbnail")

		db.Create(&Dish{Name: name, Description: description, Thumbnail: thumbnail})

        ctx.Redirect(http.StatusFound, "/")
	})

	router.GET("/", func(ctx *gin.Context) {
		var dishes []Dish
		db.Find(&dishes)

		ctx.HTML(http.StatusOK, "home.tmpl", gin.H{
			"Dishes": dishes,
		})
	})

	router.Run()
}
