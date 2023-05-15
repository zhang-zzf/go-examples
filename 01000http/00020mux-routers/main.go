package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	ReadBook := func(w http.ResponseWriter, r *http.Request) {
		pathVars := mux.Vars(r)
		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", pathVars["title"], pathVars["pageNo"])
	}
	r.HandleFunc("/books/{title}/page/{pageNo}", ReadBook).Methods("GET")
	http.ListenAndServe(":80", r)
}
