package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type Data struct {
	id_user    int
	password   string
	email      string
	username   string
	first_name string
	last_name  string
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
	db, err := sql.Open("sqlite3", "")
	if err != nil {
		fmt.Println("Erreur lors de la connexion à la base de données:", err)
		return
	}

	defer db.Close()

	query := "INSERT INTO user (id_user, password, email, username, first_name, last_name) VALUES (1,'test','test@gmail.com','test','mon_prenom_est','mon_nom_est')"

	rows, err := db.Query(query)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

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
