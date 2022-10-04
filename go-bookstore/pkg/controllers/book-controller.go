package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go-bookstore/pkg/models"
	"net/http"
	"strconv"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	w.Header().Set("Content-Type", "Application/json")
	_ = json.NewEncoder(w).Encode(books)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book, _ := models.GetBookById(ID)
	w.Header().Set("Content-Type", "Application/json")
	_ = json.NewEncoder(w).Encode(book)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	b := book.CreateBook()
	w.Header().Set("Content-Type", "Application/json")
	_ = json.NewEncoder(w).Encode(b)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	b, err := models.DeleteBook(ID)
	if err != nil {
		fmt.Println("error while query db")
	}
	w.Header().Set("Content-Type", "Application/json")
	_ = json.NewEncoder(w).Encode(b)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	var inputBook models.Book
	_ = json.NewDecoder(r.Body).Decode(&inputBook)
	b, db := models.GetBookById(ID)
	if inputBook.Name != "" {
		b.Name = inputBook.Name
	}
	if inputBook.Author != "" {
		b.Author = inputBook.Author
	}
	if inputBook.Publication != "" {
		b.Publication = inputBook.Publication
	}
	db.Save(b)
	w.Header().Set("Content-Type", "Application/json")
	_ = json.NewEncoder(w).Encode(b)
}
