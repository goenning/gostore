package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/goenning/gostore/env"
	"github.com/gorilla/mux"
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
	port := env.GetEnvOrDefault("PORT", "8080")
	r := mux.NewRouter()
	r.HandleFunc("/", index)
	fmt.Printf("Server is up and listening on port %s.\n", port)
	http.ListenAndServe(":"+port, r)
}
