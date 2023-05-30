package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Data struct {
	name string
}

var UserData Data

func homeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./static/html/Home.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Execute(w, UserData)
}

func Co_BDD() {
	db, err := sql.Open("mysql", "root:Paf151004!@tcp(localhost:3306)/BDD_Forum")
	if err != nil {
		fmt.Println("Erreur lors de la connexion à la base de données:", err)
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Erreur lors de la vérification de la connexion à la base de données:", err)
		return
	}
}

func main() {
	fmt.Println("http://localhost:8080")
	Co_BDD()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", homeHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
