package dev

import (
	"database/sql"
	"fmt"
)

func search_user() {
	db, errOpen := sql.Open("sqlite3", ".../db.db")

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
	statement.Exec()
	db.Close()
}
