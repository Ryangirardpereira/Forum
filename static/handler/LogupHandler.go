package handler

import (
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
	t.Execute(w, Acte)
}
