package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	engine.Run("0.0.0.0:80")
}
