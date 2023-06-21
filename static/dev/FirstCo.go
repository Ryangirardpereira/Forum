package dev

import (
	"database/sql"
	"fmt"
)

func Co_BDD() {
	db, err := sql.Open("sqlite3", "db.db")
	if err != nil {
		fmt.Println("erreur lors de la connexion", err)
		return
	}

	err = db.Ping()

	if err != nil {
		fmt.Println("erreur lors de la verification a la connexion", err)
		return
	}

	defer db.Close()

	fmt.Println("la connexion à la base de donnée est établie")
}
