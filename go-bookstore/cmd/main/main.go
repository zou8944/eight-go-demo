package main

import (
	"github.com/gorilla/mux"
	"go-bookstore/pkg/routes"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Println("Server start on port: 8080")
	log.Fatalln(http.ListenAndServe(":8080", r))
}
