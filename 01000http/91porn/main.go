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

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

func setupRouter() *gin.Engine {
	engine := gin.Default()
	// just for test
	//engine.Use(Cors())
	engine.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	// 静态文件
	engine.StaticFile("/", "./assets/index.html")
	engine.StaticFile("/favicon.ico", "./assets/favicon.ico")
	// assets
	engine.Static("/assets", "./assets/assets")
	//
	apiV1 := engine.Group("/api/v1")
	apiV1.GET("/videos", controller.GetVideos)
	apiV1.GET("/videos/:view_key", controller.GetVideoInfo)
	apiV1.GET("/videos/:view_key/mp4", controller.GetVideoStream)
	apiV1.GET("/videos/:view_key/pic", controller.GetPicture)
	apiV1.GET("/categories", controller.GetCategories)
	return engine
}
