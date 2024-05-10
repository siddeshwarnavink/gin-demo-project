package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	"siddapp/config"
	"siddapp/forms"
	"siddapp/models"
)

var validate *validator.Validate

func main() {
	db := config.GetDB()

	validate = validator.New()

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	createGroup := router.Group("/create")
	{
		createGroup.GET("", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "create.tmpl", gin.H{})
		})

		createGroup.POST("", func(ctx *gin.Context) {
			var form forms.DishForm
			if err := ctx.ShouldBindWith(&form, binding.Form); err != nil {
				ctx.HTML(http.StatusBadRequest, "create.tmpl", gin.H{"error": err.Error()})
				return
			}

			db.Create(&models.Dish{Name: form.Name, Description: form.Description, Thumbnail: form.Thumbnail})

			ctx.Redirect(http.StatusFound, "/")
		})
	}

	editGroup := router.Group("/edit/:id")
	{
		editGroup.GET("", func(ctx *gin.Context) {
			id := ctx.Param("id")

			var dish models.Dish
			db.First(&dish, id)

			dishForm := forms.DishForm{Name: dish.Name, Description: dish.Description, Thumbnail: dish.Thumbnail}

			ctx.HTML(http.StatusOK, "edit.tmpl", gin.H{"form": dishForm})
		})

		editGroup.POST("", func(ctx *gin.Context) {
			id := ctx.Param("id")

			var form forms.DishForm
			if err := ctx.ShouldBindWith(&form, binding.Form); err != nil {
				ctx.HTML(http.StatusBadRequest, "edit.tmpl", gin.H{"error": err.Error(), "form": form})
				return
			}

			var dish models.Dish
			db.First(&dish, id)

			db.Model(&dish).Updates(models.Dish{Name: form.Name, Description: form.Description, Thumbnail: form.Thumbnail})

			ctx.Redirect(http.StatusFound, "/")
		})
	}

	router.GET("/delete/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")

		var dish models.Dish
		db.First(&dish, id)
		db.Delete(&dish)

		ctx.Redirect(http.StatusFound, "/")
	})

	router.GET("/", func(ctx *gin.Context) {
		var dishes []models.Dish
		db.Find(&dishes)

		ctx.HTML(http.StatusOK, "home.tmpl", gin.H{
			"Dishes": dishes,
		})
	})

	router.Run()
}
