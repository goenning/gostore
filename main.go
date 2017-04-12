package main

import (
	"fmt"
	"html/template"
	"net/http"

	"math/big"

	"github.com/goenning/gostore/env"
	"github.com/goenning/gostore/expanders"
	"github.com/goenning/gostore/models"
	"github.com/gorilla/mux"
)

func createTemplate(fileName string) *template.Template {
	path := fmt.Sprintf("./views/%s.html", fileName)
	return template.Must(template.New("base.html").Funcs(expanders.All()).ParseFiles("./views/base.html", path))
}

var templates = map[string]*template.Template{
	"index": createTemplate("index"),
}

func index(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Go Store :)",
		"Products": []*models.Product{
			models.NewProduct("Metallica T-Shirt", "The oldest Metallica T-Shirt you'll ever find.", big.NewFloat(86.23)),
			models.NewProduct("Scott Road Bicycle", "Road Bicycle in mint conditions. Highly Recommended!", big.NewFloat(158.80)),
			models.NewProduct("iPhone 7", "A brand new Apple iPhone 7.", big.NewFloat(1699.99)),
		},
	}

	w.WriteHeader(http.StatusOK)
	templates["index"].Execute(w, data)
}

func main() {
	port := env.GetEnvOrDefault("PORT", "8080")
	r := mux.NewRouter()
	r.PathPrefix("/public").Handler(http.FileServer(http.Dir("./")))
	r.HandleFunc("/", index)
	fmt.Printf("Server is up and listening on port %s.\n", port)
	http.ListenAndServe(":"+port, r)
}
