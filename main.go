package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
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

type DishForm struct {
	Name        string `form:"name" binding:"required"`
	Description string `form:"description" binding:"required"`
	Thumbnail   string `form:"thumbnail" binding:"required,url"`
}

var validate *validator.Validate

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Dish{})

	validate = validator.New()

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/create", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "create.tmpl", gin.H{})
	})

	router.POST("/create", func(ctx *gin.Context) {
		var form DishForm
		if err := ctx.ShouldBindWith(&form, binding.Form); err != nil {
			ctx.HTML(http.StatusBadRequest, "create.tmpl", gin.H{"error": err.Error()})
			return
		}

		db.Create(&Dish{Name: form.Name, Description: form.Description, Thumbnail: form.Thumbnail})

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
