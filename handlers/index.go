package handlers

import (
	"net/http"

	"github.com/goenning/gostore/models"
)

//Index is our main page handler
func Index(w http.ResponseWriter, r *http.Request) {
	products := []models.Product{
		{
			Title: "iMac 27\" 5K",
			Price: 1600,
		},
		{
			Title: "Giant Bike in mint conditions",
			Price: 1299.99,
		},
	}
	data := &models.IndexViewModel{
		Title:    "Go Store :)",
		Products: products,
	}

	w.WriteHeader(http.StatusOK)
	render(w, "index.html", data)
}
