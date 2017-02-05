package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/goenning/gostore/handlers"
	"github.com/gorilla/mux"
)

func getEnvOrDefault(env string, def string) string {
	v := os.Getenv(env)
	if v != "" {
		return v
	}
	return def
}

func main() {
	port := getEnvOrDefault("PORT", "8080")
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.Index)
	fmt.Printf("Server is up and listening on port %s.\n", port)
	http.ListenAndServe(":"+port, r)
}
