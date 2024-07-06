package main

import (
	"net/http"
	"uacforcast/WebServer/services"

	"github.com/gin-gonic/gin"
)

// Main Webserver
func main() {
  r := gin.Default()
  
  // MARK: Ping
  r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })

  // MARK: Todays Forcast
  r.GET("/today-forcast", services.FetchUACForcast)

  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

