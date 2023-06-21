package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Post struct {
	Content string `json:"content"`
}

var posts []Post

func post() {
	http.HandleFunc("/posts", handlePosts)
	http.ListenAndServe(":8080", nil)
}

func handlePosts(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Erreur de lecture du corps de la requête", http.StatusInternalServerError)
			return
		}

		var post Post
		err = json.Unmarshal(body, &post)
		if err != nil {
			http.Error(w, "Erreur de décodage JSON", http.StatusBadRequest)
			return
		}

		posts = append(posts, post)
	}

	jsonResponse, err := json.Marshal(posts)
	if err != nil {
		http.Error(w, "Erreur de codage JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
