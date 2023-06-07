package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

type Album struct {
	ID     int64   `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float32 `json:"price"`
}

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
	albums, err := albumsByArtist("John Coltrane")
	if err != nil {
		log.Fatal(err)
	}
	albumsJson, _ := json.Marshal(albums)
	fmt.Printf("Albums found: %v\n", string(albumsJson))
	var id int64 = 1
	album, err := albumById(id)
	if err != nil {
		log.Fatal(err)
	}
	albumJson, _ := json.Marshal(album)
	fmt.Printf("Album By Id %d: %s", id, string(albumJson))
	_, err = addAlbum(Album{
		Title:  "zhang.zzf",
		Artist: "zhang.zzf",
		Price:  3.45,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func albumsByArtist(name string) ([]Album, error) {
	var albums []Album
	rows, err := db.Query("select `id`, `title`, `price`, `artist` from album where artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)
	for rows.Next() {
		var r Album
		// 必须和列名保持一致
		if err = rows.Scan(&r.ID, &r.Title, &r.Price, &r.Artist); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, r)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return albums, nil
}

func albumById(id int64) (Album, error) {
	var r Album
	row := db.QueryRow("select * from album where `id` = ?", id)
	if err := row.Scan(&r.ID, &r.Title, &r.Artist, &r.Price); err != nil {
		if err == sql.ErrNoRows {
			return r, fmt.Errorf("albumById %d: no such album", id)
		}
		return r, fmt.Errorf("albuById %d: %v", id, err)
	}
	return r, nil
}

func addAlbum(alb Album) (int64, error) {
	insertSql := "insert into album(title, artist, price) values (?, ?, ?)"
	resp, err := db.Exec(insertSql, alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, err
	}
	id, err := resp.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}
