package controller

import (
	"github.com/91porn/infra/db/mysql/_91porn"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type categoryResp struct {
	Id        uint
	Category  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func GetCategories(c *gin.Context) {
	c.JSON(http.StatusOK, queryAllCategories())
}

func queryAllCategories() *[]categoryResp {
	var data []categoryResp
	_91porn.Db().Table("tb_category").Find(&data)
	return &data
}
