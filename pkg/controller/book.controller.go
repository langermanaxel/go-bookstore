package controller

import (
	"encoding/json"
	"fmt"
	"go-bookstore/pkg/model"
	"go-bookstore/pkg/util"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var NewBook model.Book

func GetBooks(response http.ResponseWriter, request *http.Request) {
	newBooks := model.GetBooks()
	res, _ := json.Marshal(newBooks)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(res)
}

func GetBookById(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing...")
	}
	bookDetails, _ := model.GetBookById(Id)
	res, _ := json.Marshal(bookDetails)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(res)
}

func CreateBook(response http.ResponseWriter, request *http.Request) {
	newBook := &model.Book{}
	util.ParseBody(request, newBook)
	b := newBook.CreateBook()
	res, _ := json.Marshal(b)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(res)
}

func UpdateBook(response http.ResponseWriter, request *http.Request) {
	bookModel := &model.Book{}
	util.ParseBody(request, bookModel)
	vars := mux.Vars(request)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing...")
	}
	bookDetails, db := model.GetBookById(Id)
	if bookModel.Name != "" {
		bookDetails.Name = bookModel.Name
	}
	if bookModel.Author != "" {
		bookDetails.Author = bookModel.Author
	}
	if bookModel.Publication != "" {
		bookDetails.Publication = bookModel.Publication
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(res)
}

func DeleteBook(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing...")
	}
	book := model.DeleteBook(Id)
	res, _ := json.Marshal(book)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(res)
}
