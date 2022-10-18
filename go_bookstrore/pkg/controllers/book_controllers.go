package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/dimeprog/BookStore-Manager-App/pkg/models"
	"github.com/dimeprog/BookStore-Manager-App/pkg/utils"

)

// book variable
var NewBook models.Book

//  getBooks endpoint
func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBook()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//  getBook by ID endpoint
func GetBookById(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	bookid := param["ID"]
	ID, err := strconv.ParseInt(bookid, 0, 0)
	if err != nil {
		fmt.Println("Error occured when parsing")
	}
	bookDetails, _ := models.GetBookByID(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//  delete endpoint
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	bookid := param["ID"]
	ID, err := strconv.ParseInt(bookid, 0, 0)
	if err != nil {
		fmt.Println("Error occured when parsing")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//  create new book endpoint
func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

//  update endpoint
func UpdateBook(w http.ResponseWriter, r *http.Request) {

}
