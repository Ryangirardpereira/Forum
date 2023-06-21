package main

import (
	"fmt"
	"log"
	"net/http"

	"Forum/Forum/static/dev"
	"Forum/Forum/static/handler"
)

func main() {
	fmt.Println("http://localhost:8080")
	dev.Co_BDD()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", handler.HomeHandler)
	http.HandleFunc("/signUp", handler.LogupHandler)
	http.HandleFunc("/signIn", handler.LoginHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
