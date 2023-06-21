package dev

import (
	"fmt"
)

func Check_log(nameTry string, pw string) int {

	BDD := recup()

	// fmt.Println(BDD[0].name)

	for _, valeur := range BDD {
		if valeur.name == nameTry && check_hash(pw, valeur.password) {
			fmt.Println("ce sont vos identifiant")
			return valeur.ID
		}
	}
	fmt.Println("ce ne sont pas vos identifiant")

	return 0
}
