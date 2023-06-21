package dev

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func New_user(nameUser string, mdpUser string, emailUser string, birthday string, sport string) {
	db, errOpen := sql.Open("sqlite3", "db.db")

	if errOpen != nil {
		fmt.Println("erreur :", errOpen)
		return
	}

	statement, err := db.Prepare("INSERT INTO User (pw,username,email,birthday,sport) VALUES(?,?,?,?,?)")
	if err != nil {
		fmt.Println(err)
		fmt.Println("error Prepare new user")
		return
	}
	statement.Exec(mdpUser, nameUser, emailUser, birthday, sport)
	db.Close()

	fmt.Println("OK")
}
