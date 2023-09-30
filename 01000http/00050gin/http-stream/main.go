package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	engine := gin.Default()
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	engine.GET("/videos/:view_key/stream", GetVideoStream)
	engine.GET("/videos/:view_key/pic", GetVideoPic)
	// index.html
	engine.StaticFile("/", "./assets/index.html")
	engine.StaticFile("/favicon.ico", "./assets/favicon.ico")
	// assets
	engine.Static("/assets", "./assets")
	engine.Run("0.0.0.0:80")
}

func GetVideoPic(c *gin.Context) {
	viewKey := c.Param("view_key")
	file, err := searchFile(viewKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.File(file.Name())
}

func GetVideoStream(c *gin.Context) {
	videoFilePath := c.Param("view_key")
	if videoFilePath == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "video not exists"})
		return
	}
	file, err := searchFile(videoFilePath)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	stat, _ := file.Stat()
	http.ServeContent(c.Writer, c.Request, videoFilePath, stat.ModTime(), file)
}

func searchFile(path string) (*os.File, error) {
	parentDir := "/tmp"
	fileDir, err := os.Open(parentDir)
	if err != nil {
		return nil, err
	}
	dirEntries, err := fileDir.ReadDir(0)
	if err != nil {
		return nil, err
	}
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			continue
		}
		info, err := entry.Info()
		if err != nil {
			continue
		}
		file, err := os.Open(filepath.Join(parentDir, info.Name(), path))
		if err != nil {
			continue
		} else {
			// stop search
			return file, nil
		}
	}
	return nil, errors.New("file not exists => " + path)
}
