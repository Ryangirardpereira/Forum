package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"Forum/static/dev"
	"Forum/static/handler"
)

func Home(w http.ResponseWriter, r *http.Request) {
	test := template.Must(template.ParseFiles(("Home.tmpl")))
	if err := test.Execute(w, test); err != nil {
		log.Println("err")
	}
}

func main() {
	fmt.Println("http://localhost:8080")
	static := http.FileServer(http.Dir("../css"))
	http.Handle("/static/", http.StripPrefix("/static/", static))
	http.HandleFunc("/", Home)
	http.HandleFunc("/signUp", handler.LogupHandler)
	http.HandleFunc("/signIn", handler.LoginHandler)
	dev.Co_BDD()
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
