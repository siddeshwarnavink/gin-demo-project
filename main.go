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

	createGroup := router.Group("/create")
	{
		createGroup.GET("", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "create.tmpl", gin.H{})
		})

		createGroup.POST("", func(ctx *gin.Context) {
			var form DishForm
			if err := ctx.ShouldBindWith(&form, binding.Form); err != nil {
				ctx.HTML(http.StatusBadRequest, "create.tmpl", gin.H{"error": err.Error()})
				return
			}

			db.Create(&Dish{Name: form.Name, Description: form.Description, Thumbnail: form.Thumbnail})

			ctx.Redirect(http.StatusFound, "/")
		})
	}

	editGroup := router.Group("/edit/:id")
	{
		editGroup.GET("", func(ctx *gin.Context) {
			id := ctx.Param("id")

			var dish Dish
			db.First(&dish, id)

			dishForm := DishForm{Name: dish.Name, Description: dish.Description, Thumbnail: dish.Thumbnail}

			ctx.HTML(http.StatusOK, "edit.tmpl", gin.H{"form": dishForm})
		})

		editGroup.POST("", func(ctx *gin.Context) {
			id := ctx.Param("id")

			var form DishForm
			if err := ctx.ShouldBindWith(&form, binding.Form); err != nil {
				ctx.HTML(http.StatusBadRequest, "edit.tmpl", gin.H{"error": err.Error(), "form": form})
				return
			}

			var dish Dish
			db.First(&dish, id)

			db.Model(&dish).Updates(Dish{Name: form.Name, Description: form.Description, Thumbnail: form.Thumbnail})

			ctx.Redirect(http.StatusFound, "/")
		})
	}

	router.GET("/delete/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")

		var dish Dish
		db.First(&dish, id)
        db.Delete(&dish)

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
