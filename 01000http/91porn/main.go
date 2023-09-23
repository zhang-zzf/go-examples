package main

import (
	"github.com/91porn/infra/http/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := setupRouter()
	engine.Run("0.0.0.0:80")
}

func setupRouter() *gin.Engine {
	engine := gin.Default()
	engine.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	apiV1 := engine.Group("/api/v1")
	apiV1.GET("/videos", controller.GetVideos)
	apiV1.GET("/videos/:view_key/stream", controller.GetVideoStream)
	apiV1.GET("/categories", controller.GetCategories)
	return engine
}
