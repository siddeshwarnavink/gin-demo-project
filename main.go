package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"siddapp/controllers"
)

var validate *validator.Validate

func main() {
	validate = validator.New()

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	recipieController := new(controllers.RecipieController)
	createGroup := router.Group("/create")
	{
		createGroup.GET("", recipieController.GetCreate)
		createGroup.POST("", recipieController.PostCreate)
	}
	editGroup := router.Group("/edit/:id")
	{
		editGroup.GET("", recipieController.GetEdit)
		editGroup.POST("", recipieController.PostEdit)
	}
	router.GET("/delete/:id", recipieController.Delete)
	router.GET("/", recipieController.Home)

	router.Run()
}
