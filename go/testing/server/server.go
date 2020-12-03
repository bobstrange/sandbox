package main

import (
	"fmt"
	"log"
	"net/http"
)

func Route() *http.ServeMux {
	m := http.NewServeMux()
	m.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		fmt.Fprintf(w, "Hi, %s", r.FormValue("name"))
	})
	return m
}

func main() {
	m := Route()
	log.Fatal(http.ListenAndServe(":8080", m))
}
