package handler

import (
	"Forum/Forum/static/dev"
	"net/http"
)

func recupForm(r *http.Request) {
	name := r.FormValue("Identifiant")
	mdp := dev.Hash(r.FormValue("Mdp"))
	email := r.FormValue("Email")
	date := r.FormValue("Date")
	sport := r.FormValue("Sport")
	if name != "" && mdp != "" && email != "" {
		dev.New_user(name, mdp, email, date, sport)
	}
}
