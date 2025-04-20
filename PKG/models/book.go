package models

import (
	"github.com/jinzhu/gorm"
	"github.com/awatar101/go_bookstore/pkg/config"
)

var db *gorm.DB

type Book struct {
	// 	"gorm:primary_key" json:"id"
	gorm.model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b) // create a new record
	db.Create(&b)   // save the record to the database
	return b
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books) // find all books in the database
	return books
}

func GetBookById(bookId int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", bookId).Find(&getBook) // find the book by ID
	return &getBook, db
}

func DeleteBook(bookId int64) Book {
	var book Book
	db.Where("ID=?", bookId).Delete(book) // delete the book by ID
	return book
}
