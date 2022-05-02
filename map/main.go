package main

import (
	"fmt"
	"log"
	"net/http"
)

type Type string

const (
	MINE Type = "MINE"
)

type Service struct {
	row, col    int32
	serviceType Type
}

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		log.Println("hello")
		_, err := fmt.Fprintf(w, "Hello!")
		if err != nil {
			return
		}
	})

	http.HandleFunc("/index.html", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
