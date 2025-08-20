package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/shivanshcharak/book-server/pkg/models"
	"github.com/shivanshcharak/book-server/pkg/utils"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, err := json.Marshal(newBooks)
	if err != nil {
		http.Error(w, "failed to marshal books", http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["bookId"]
	ID, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		http.Error(w, "Failed to parse the id", http.StatusInternalServerError)
		return
	}

	bookDetails, db := models.GetBooksById(ID)
	if db.Error != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	res, err := json.Marshal(bookDetails)
	if err != nil {
		http.Error(w, "Something went wrong when marshaling book details", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	if err := utils.ParseBody(r, book); err != nil {
		http.Error(w, "Invalid requesr body", http.StatusBadRequest)
	}
	createdBook, err := book.CreateBook()
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
	res, err := json.Marshal(createdBook)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	updateBook := &models.Book{}
	if err := utils.ParseBody(r, updateBook); err != nil {
		http.Error(w, "Something went wrong while parsing", http.StatusInternalServerError)
		return
	}
	vars := mux.Vars(r)
	bookid := vars["bookId"]
	ID, err := strconv.ParseInt(bookid, 10, 64)
	if err != nil {
		http.Error(w,"Something went wrong while parsing",http.StatusInternalServerError)
		return
	}
	bookDetails, db:=models.GetBooksById(uint(ID))
	if db.Error!=nil{
		http.Error(w, "No book found with that id", http.StatusInternalServerError)
		return
	}
	bookDetails.Author=updateBook.Author,
	bookDetails.Name=updateBook.Name,
	bookDetails.Publication=updateBook.Publication
	db.Save(&bookDetails)
	res, err:=json.Marshal(bookDetails)
	if err!=nil{
		http.Error(w, "Something went wrong while masrahlling", http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
