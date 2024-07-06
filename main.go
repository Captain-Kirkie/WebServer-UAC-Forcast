package main

import (
	"net/http"
	"uacforcast/WebServer/services"

	"github.com/gin-gonic/gin"
)

// Main Webserver
func main() {
  r := gin.Default()
  r.GET("/ping", func(c *gin.Context) {
    services.FetchUACForcast()
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })
  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

