package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello %v\n", r.URL.Path)
	})
	http.HandleFunc("/headers", func(w http.ResponseWriter, r *http.Request) {
		for key, val := range r.Header {
			fmt.Fprintf(w, "%s: %v\n", key, val)
		}
	})
	http.HandleFunc("/context", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("server: hello handler started")
		defer fmt.Println("server: hello handler ended")
		ctx := r.Context()
		select {
		case <-time.After(10 * time.Second):
			fmt.Fprintf(w, "hello")
		case <-ctx.Done():
			err := ctx.Err()
			fmt.Println("server err:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	http.ListenAndServe(":80", nil)
}
