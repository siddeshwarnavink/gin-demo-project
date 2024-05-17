package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

    "siddapp/routes"
)

var validate *validator.Validate

func main() {
	validate = validator.New()

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

    routes.SetupRoutes(router) 

	router.Run()
}
