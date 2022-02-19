package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rachidassouani/go-bookstore/pkg/models"
	"github.com/rachidassouani/go-bookstore/pkg/util"
)

var books models.Book

func Books(w http.ResponseWriter, r *http.Request) {
	books := models.FindAll()
	result, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)

}

func FindBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing..")
	}
	book, _ := models.FindBookById(id)

	result, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)

}

func New(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	util.ParseBody(r, book)
	craetedBook := book.New()
	result, _ := json.Marshal(craetedBook)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)

}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing..")
	}
	deletedBook := models.Delete(id)
	result, _ := json.Marshal(deletedBook)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
