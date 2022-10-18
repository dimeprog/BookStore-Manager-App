package models

import (
	"github.com/jinzhu/gorm"

	"github.com/dimeprog/BookStore-Manager-App/pkg/config"

)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"name" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

// connect to database  and get db from config file
func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

// method create book inside the db
func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

// method to all the books from the db
func GetAllBook() []Book {
	var Book []Book
	db.Find(&Book)
	return Book
}

// method to get a single book
func GetBookByID(id int64) (*Book, *gorm.DB) {
	var getBook Book
	db.Where("ID=?", id).Find(&getBook)
	return &getBook, db
}

// method to delete  a single book from db
func DeleteBook(id int64) Book {
	var book Book
	db.Where("ID=?", id).Delete(book)
	return book
}
