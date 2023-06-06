package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func main() {
	cfg := mysql.Config{
		Net:                  "tcp",
		Addr:                 "dj1r1b5hi2x7vwle.zhanfengzhang.top:3306",
		DBName:               "recordings",
		User:                 "root",
		Passwd:               "test",
		AllowNativePasswords: true,
	}
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected.")

}
