package main

import (
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("asserts/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":80", nil)
}
