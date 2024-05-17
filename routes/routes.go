package routes

import (
	"siddapp/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
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
}
