package handler

import (
<<<<<<< HEAD
=======
	"Forum/static/dev"
>>>>>>> f751755 (last update)
	"fmt"
	"net/http"
	"text/template"
)

func LogupHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("static/html/NewID.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Execute(w, Acte)
}
