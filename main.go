package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
        c.Header("Content-Type", "text/html")
        c.String(http.StatusOK, "<h1>Hello world</h1>")
	})
	r.Run()
}
