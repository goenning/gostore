package main

import (
	"fmt"
	"net/http"

	"os"

	"github.com/goenning/gostore/handlers"
	"github.com/gorilla/mux"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port != "" {
		return port
	}
	return "8080"
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.Index)
	fmt.Println("Server is up and listening on port 8080.")
	http.ListenAndServe(":"+getPort(), r)
}
