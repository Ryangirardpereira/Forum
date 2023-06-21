package handler

import (
	"fmt"
	"net/http"
	"text/template"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("static/html/Home.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("ici")
	t.Execute(w, nil)
}
