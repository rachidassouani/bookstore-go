package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rachidassouani/go-bookstore/pkg/routes"
)

func main() {
	result := mux.NewRouter()
	routes.BookStoreRoute(result)
	http.Handle("/", result)
	log.Fatal(http.ListenAndServe("localhost:8080", result))

}
