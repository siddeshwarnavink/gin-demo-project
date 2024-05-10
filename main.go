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

	recipie := new(controllers.RecipieController)
	createGroup := router.Group("/create")
	{
		createGroup.GET("", recipie.GetCreate)
		createGroup.POST("", recipie.PostCreate)
	}
	editGroup := router.Group("/edit/:id")
	{
		editGroup.GET("", recipie.GetEdit)
		editGroup.POST("", recipie.PostEdit)
	}
	router.GET("/delete/:id", recipie.Delete)
	router.GET("/", recipie.Home)

	router.Run()
}
