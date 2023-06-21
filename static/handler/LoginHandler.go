package handler

import (
	"Forum/Forum/static/dev"
	"fmt"
	"net/http"
	"text/template"
)

type Action struct {
	Action  string
	Message string
}

var acte Action

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("Forum/static/html/Id.html")
	if err != nil {
		fmt.Println(err)
		return
	}

	nameUser := r.FormValue("Id")
	mdp := r.FormValue("Mdp")

	fmt.Println(nameUser)

	if nameUser != "" && mdp != "" {
		if dev.Check_log(nameUser, mdp) != 0 {
			acte.Action = "/"
		} else {
			acte.Action = "/signIn"
			acte.Message = "L'identifiant ou le mot de passe n'est pas valide !"
		}
	}
	t.Execute(w, acte)

}
