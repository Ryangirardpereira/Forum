package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("http://localhost:8080")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
