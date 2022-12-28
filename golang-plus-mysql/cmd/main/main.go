package main

import (
	"log"
	"net/http"

	"github.com/alekseiapa/mini-go-projects/golang-plus-mysql/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}