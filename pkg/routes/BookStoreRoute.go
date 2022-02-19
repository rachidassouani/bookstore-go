package routes

import (
	"github.com/gorilla/mux"
	"github.com/rachidassouani/go-bookstore/pkg/controllers"
)

var BookStoreRoute = func(router *mux.Router) {
	router.HandleFunc("/books/", controllers.New).Methods("POST")
	router.HandleFunc("books", controllers.Books).Methods("GET")
	router.HandleFunc("books/{bookId}", controllers.FindBookById).Methods("GET")
	router.HandleFunc("books/{bookId}", controllers.Delete).Methods("DELETE")
}
