package controller

import (
	"errors"
	"github.com/91porn/infra/db/mysql/_91porn"
	"github.com/91porn/infra/http/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type getVideosReq struct {
	Tag    string `form:"tag"`
	SortBy string `form:"sort_by"`
	Title  string `form:"title"`
	Limit  int    `form:"limit"`
	Offset int    `form:"offset"`
}

type videoResp struct {
	Id            uint
	Title         string
	ViewKey       string
	Duration      string
	AddedAt       string
	VideoFilePath string
	PicFilePath   string
	IsHd          int
	Popularity    int
	Favorites     int
	Like          int
	Unlike        int
	Tags          string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func GetVideos(c *gin.Context) {
	var req getVideosReq
	if err := c.Bind(&req); err != nil {
		panic(err)
	}
	initPageIfNeeded(&req)
	log.Printf("GetVideos req -> %#v\n", req)
	var resp *model.Response[videoResp]
	if req.Tag != "" {
		resp = queryPageByTag(req)
	} else if req.Title != "" {
		resp = queryPageByTitle(req.Title, req.SortBy, req.Offset, req.Limit)
	}
	c.JSON(http.StatusOK, resp)
}

func GetVideoInfo(c *gin.Context) {
	viewKey := c.Param("view_key")
	video := queryVideoByViewKey(viewKey)
	c.JSON(http.StatusOK, video)
}

func GetVideoStream(c *gin.Context) {
	viewKey := c.Param("view_key")
	video := queryVideoByViewKey(viewKey)
	videoFilePath := video.VideoFilePath
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

func GetPicture(c *gin.Context) {
	viewKey := c.Param("view_key")
	video := queryVideoByViewKey(viewKey)
	picFilePath := video.PicFilePath
	if picFilePath == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "video not exists"})
		return
	}
	file, err := searchFile(picFilePath)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	stat, _ := file.Stat()
	http.ServeContent(c.Writer, c.Request, picFilePath, stat.ModTime(), file)
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

func queryVideoByViewKey(viewKey string) *videoResp {
	var data videoResp
	_91porn.Db().Raw("select * from tb_video where view_key = ?", viewKey).Scan(&data)
	return &data
}

func initPageIfNeeded(req *getVideosReq) {
	if req.Limit == 0 {
		req.Limit = model.PageDefaultLimit
	}
}

func queryPageByTitle(title string, sortBy string, offset int, limit int) *model.Response[videoResp] {
	var data []videoResp
	var total int64
	// count
	sql := "select count(id) from tb_video where title like ?"
	_91porn.Db().Raw(sql, "%"+title+"%").Count(&total)
	switch sortBy {
	case "-popularity":
		sql := "select * from tb_video where title like ? order by ? desc limit ?, ?"
		_91porn.Db().Raw(sql, "%"+title+"%", sortBy, offset, limit).Scan(&data)
	default:
		log.Printf("queryByPageByTitle unknown sortBy %v", sortBy)
		sql := "select * from tb_video where title like ? order by added_at desc limit ?, ?"
		_91porn.Db().Raw(sql, "%"+title+"%", offset, limit).Scan(&data)
	}
	return model.NewResponse(&data, &model.Page{Total: total, Offset: offset, Limit: limit})
}

func queryPageByTag(req getVideosReq) *model.Response[videoResp] {
	var total int64
	var data []videoResp
	req.Tag = "%" + req.Tag + "%"
	// count
	sql := "select count(id) from tb_video where tags like ?"
	_91porn.Db().Raw(sql, req.Tag).Count(&total)
	// data
	switch req.SortBy {
	case "-popularity":
		//sql := "select * from tb_video where tags like ? order by ? desc limit ?, ?"
		//log.Printf("queryPageByTag sortBy -> %v", sortBy)
		//_91porn.Db().Raw(sql, "%"+tag+"%", sortBy[1:], offset, limit).Scan(&data)
		// order by 后面不能跟 'popularity'
		// select * from tb_video where tags like '%每月最热%' order by 'popularity' desc;
		req.SortBy = req.SortBy[1:]
		sql := "select * from tb_video where tags like @Tag order by popularity desc limit @Offset, @Limit"
		log.Printf("queryPageByTag sortBy -> %v", req.SortBy)
		_91porn.Db().Raw(sql, req).Scan(&data)
	default:
		log.Printf("queryPageByTag sortBy -> %v", req.SortBy)
		sql := "select * from tb_video where tags like @Tag order by added_at desc limit @Offset, @Limit"
		_91porn.Db().Raw(sql, req).Scan(&data)
	}
	return model.NewResponse(&data, &model.Page{Total: total, Offset: req.Offset, Limit: req.Limit})
}
