package handler

import (
	"Forum/static/dev"
	"fmt"
	"net/http"
	"text/template"
)

type Action struct {
	Action  string
	Message string
}

var Acte Action

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("Forum/static/html/Id.html")
	if err != nil {
		fmt.Println(err)
		return
	}

	nameUser := r.FormValue("Id")
	mdp := r.FormValue("Mdp")

	if nameUser != "" && mdp != "" {
		if dev.Check_log(nameUser, mdp) != 0 {
			Acte.Action = "/"
		} else {
			Acte.Action = "/signIn"
			Acte.Message = "L'identifiant ou le mot de passe n'est pas valide !"
		}
	}
	t.Execute(w, Acte)

}
