package _91porn

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var _91porn *gorm.DB

func Db() *gorm.DB {
	if _91porn != nil {
		return _91porn
	} else {
		// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
		dsn := "root:test@tcp(10.10.10.8:3306)/91porn?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
		if err != nil {
			panic(err)
		}
		_91porn = db
		return _91porn
	}
}
