package handler

import (
	"Forum/Forum/static/dev"
	"fmt"
	"net/http"
	"text/template"
)

func LogupHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("Forum/static/html/NewID.html")
	if err != nil {
		fmt.Println(err)
		return
	}

	name := r.FormValue("Id")
	mdp := dev.Hash(r.FormValue("Mdp"))
	email := r.FormValue("Email")
	date := r.FormValue("Date")
	sport := r.FormValue("Sport")

	fmt.Println(name)

	if name != "" && mdp != "" && email != "" {
		dev.New_user(name, mdp, email, date, sport)
	}
	t.Execute(w, nil)
}
