package dev

import (
	"database/sql"
	"fmt"
	"log"
)

type Database struct {
	ID       int
	name     string
	password string
	email    string
	birthday string
	sport    string
}

var UserData Database
var ListUser []Database

func recup() []Database {
	db, errOpen := sql.Open("sqlite3", "db.db")

	if errOpen != nil {
		fmt.Println("erreur :", errOpen)
		return nil
	}

	rows, err := db.Query("SELECT * FROM User")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&UserData.ID, &UserData.name, &UserData.password, &UserData.email, &UserData.birthday, &UserData.sport)
		if err != nil {
			log.Fatal(err)
		}
		ListUser = append(ListUser, UserData)
	}

	return ListUser
}
