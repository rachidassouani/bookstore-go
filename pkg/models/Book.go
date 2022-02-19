package models

import (
	"github.com/jinzhu/gorm"
	"github.com/rachidassouani/go-bookstore/pkg/configs"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	configs.Connect()
	db = configs.GetDb()
	db.AutoMigrate(&Book{})
}

func (book *Book) New() *Book {
	db.NewRecord(book)
	db.Create(&book)
	return book
}

func FindAll() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func FindBookById(id int64) (*Book, *gorm.DB) {
	var book Book
	db := db.Where("id=?", id).Find(&book)
	return &book, db
}

func Delete(id int64) Book {
	var book Book
	db.Where("id=?", id).Delete(book)
	return book
}
