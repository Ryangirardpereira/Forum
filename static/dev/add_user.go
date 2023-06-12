package dev

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func add_db(pw string, email string, f_name string, l_name string) {
	db, err := sql.Open("sqlite3", "../db.db")
	if err != nil {
		fmt.Println("erreur lors de la connexion", err)
		return
	}

	defer db.Close()

	query := "INSERT INTO user (password, username, email, first_name, last_name) values (?,?,?,?,?)"

	db.Query(query, pw, email, f_name, l_name)

	fmt.Println("un nouvel utilisateur nous a rejoint")
}
