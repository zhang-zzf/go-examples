package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var p = fmt.Println

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	db, err := sql.Open("mysql", "root:test@(10.0.9.18:3306)/test?parseTime=true&loc=Local")
	checkErr(err)
	defer db.Close()
	checkErr(db.Ping())
	createTable(db)
	defer dropTable(db)
	insertIntoTable(db)
	queryFromTable(db)
	queryAll(db)
	deleteFromTable(db)
}

func deleteFromTable(db *sql.DB) {
	result, err := db.Exec(`delete from users where id = ?`, 1)
	checkErr(err)
	p(result)
}

type User struct {
	id        int
	username  string
	password  string
	createdAt time.Time
}

func queryAll(db *sql.DB) {
	rows, err := db.Query(`select * from users`)
	checkErr(err)
	defer rows.Close()
	var users []User
	for rows.Next() {
		var u User
		checkErr(rows.Scan(&u.id, &u.username, &u.password, &u.createdAt))
		users = append(users, u)
	}
	checkErr(rows.Err())
}

func queryFromTable(db *sql.DB) {
	var (
		id        int
		username  string
		password  string
		createdAt time.Time
	)
	err := db.QueryRow(`select * from users where id=?`, 1).Scan(&id, &username, &password, &createdAt)
	checkErr(err)
	p(id, username, password, createdAt)
}

func insertIntoTable(db *sql.DB) {
	username := "johndoe"
	password := "secret"
	createdAt := time.Now()
	result, err := db.Exec(`insert into users(username, password, created_at) values (?,?,?)`,
		username, password, createdAt)
	checkErr(err)
	lastInsertId, _ := result.LastInsertId()
	rowAffected, _ := result.RowsAffected()
	p("insertIntoTable result ->", result, lastInsertId, rowAffected)
}

func dropTable(db *sql.DB) {
	result, err := db.Exec(` DROP TABLE IF EXISTS users; `)
	checkErr(err)
	p("dropTable result ->", result)
}

func createTable(db *sql.DB) {
	query := `
    CREATE TABLE users (
        id INT AUTO_INCREMENT,
        username TEXT NOT NULL,
        password TEXT NOT NULL,
        created_at DATETIME,
        PRIMARY KEY (id)
    );`
	result, err := db.Exec(query)
	checkErr(err)
	p("createTable result ->", result)
}
