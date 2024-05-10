package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"siddapp/config"
	"siddapp/forms"
	"siddapp/models"
)

type RecipieController struct{}

var db = config.GetDB()

func (ctl RecipieController) Home(ctx *gin.Context) {
	var dishes []models.Dish
	db.Find(&dishes)

	ctx.HTML(http.StatusOK, "home.tmpl", gin.H{
		"Dishes": dishes,
	})
}

func (ctl RecipieController) GetCreate(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "create.tmpl", gin.H{})
}

func (ctl RecipieController) PostCreate(ctx *gin.Context) {
	var form forms.DishForm
	if err := ctx.ShouldBindWith(&form, binding.Form); err != nil {
		ctx.HTML(http.StatusBadRequest, "create.tmpl", gin.H{"error": err.Error()})
		return
	}

	db.Create(&models.Dish{Name: form.Name, Description: form.Description, Thumbnail: form.Thumbnail})

	ctx.Redirect(http.StatusFound, "/")
}

func (ctl RecipieController) GetEdit(ctx *gin.Context) {
	id := ctx.Param("id")

	var dish models.Dish
	db.First(&dish, id)

	dishForm := forms.DishForm{Name: dish.Name, Description: dish.Description, Thumbnail: dish.Thumbnail}

	ctx.HTML(http.StatusOK, "edit.tmpl", gin.H{"form": dishForm})
}

func (ctl RecipieController) PostEdit(ctx *gin.Context) {
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
}

func (ctl RecipieController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	var dish models.Dish
	db.First(&dish, id)
	db.Delete(&dish)

	ctx.Redirect(http.StatusFound, "/")
}
