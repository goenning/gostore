package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	tpl, _ := template.ParseFiles("index.html")
	data := map[string]string{
		"Title": "Go Store :)",
	}
	w.WriteHeader(http.StatusOK)
	tpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", index)
	fmt.Println("Server is up and listening on port 8080.")
	http.ListenAndServe(":8080", nil)
}
