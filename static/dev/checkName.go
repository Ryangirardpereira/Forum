package dev

import "fmt"

func CheckName(nameTry string) string {

	BDD := recup()

	// fmt.Println(BDD[0].name)

	for _, valeur := range BDD {
		fmt.Println(valeur.name)
		if valeur.name == nameTry {
			return "Attention"
		}
	}

	return ""
}
