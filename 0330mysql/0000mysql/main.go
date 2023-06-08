package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

type Album struct {
	ID       int64   `json:"id"`
	Title    string  `json:"title"`
	Artist   string  `json:"artist"`
	Price    float32 `json:"price"`
	Quantity int     `json:"quantity"`
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
	fmt.Printf("Album By Id %d: %s\n", id, string(albumJson))
	_, err = addAlbum(Album{
		Title:  "zhang.zzf",
		Artist: "zhang.zzf",
		Price:  3.45,
	})
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	orderID, err := CreateOrder(ctx, 3, 2, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("create new order: %d\n", orderID)
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
	if err := row.Scan(&r.ID, &r.Title, &r.Artist, &r.Price, &r.Quantity); err != nil {
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

func CreateOrder(ctx context.Context, albumID, quantity, custID int) (int64, error) {
	fail := func(err error) (int64, error) {
		return 0, fmt.Errorf("CreateOrder: %v", err)
	}
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return fail(err)
	}
	//Defer a rollback in case anything fails
	defer tx.Rollback()
	// Confirm that album inventory is enough for the order.
	var enough bool
	err = tx.
		QueryRowContext(ctx, "select (quantity>?) from album where id = ?", quantity, albumID).
		Scan(&enough)
	if err != nil {
		if err == sql.ErrNoRows {
			return fail(fmt.Errorf("no such album: %d", albumID))
		}
		return fail(err)
	}
	if !enough {
		return fail(fmt.Errorf("not enough quantity"))
	}
	// Update the album inventory to remove quantity in the order.
	_, err = tx.ExecContext(ctx, "update album set quantity = quantity-? where id = ?", quantity, albumID)
	if err != nil {
		return fail(err)
	}
	result, err := tx.ExecContext(ctx, "insert into album_order(album_id, custom_id, quantity) values (?,?,?)",
		albumID, custID, quantity)
	if err != nil {
		return fail(err)
	}
	// get the ID of the order item just created
	var orderID int64
	if orderID, err = result.LastInsertId(); err != nil {
		return fail(err)
	}
	// Commit the transaction
	if err = tx.Commit(); err != nil {
		return fail(err)
	}
	return orderID, nil
}
