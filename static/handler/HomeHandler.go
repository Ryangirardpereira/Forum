package handler

import (
	"fmt"
	"net/http"
	"text/template"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("Forum/static/html/Home.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Execute(w, nil)
}
